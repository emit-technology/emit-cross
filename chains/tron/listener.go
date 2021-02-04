// Copyright 2020 EMIT Foundation.
// This file is part of E.M.I.T. .

// E.M.I.T. is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// E.M.I.T. is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with E.M.I.T. . If not, see <http://www.gnu.org/licenses/>.

package tron

import (
	"errors"
	"github.com/emit-technology/emit-cross/core"
	tronCore "github.com/fbsobreira/gotron-sdk/pkg/proto/core"

	log "github.com/emit-technology/emit-cross/log"
	metrics "github.com/emit-technology/emit-cross/metrics/types"
	"github.com/emit-technology/emit-cross/types"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"math/big"
	"time"

	"github.com/emit-technology/emit-cross/chains"
	utils "github.com/emit-technology/emit-cross/shared/tron"
	tron "github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

var BlockRetryInterval = time.Second * 10
var WatchDuration = 60 * 30 // 30 Minute

type listener struct {
	cfg                Config
	chainDB            *chains.ChainDB
	conn               Connection
	writer             *writer
	log                log.Logger
	stop               <-chan int
	sysErr             chan<- error // Reports fatal error to core
	latestBlock        metrics.LatestBlock
	metrics            *metrics.ChainMetrics
	blockConfirmations *big.Int
	state              map[types.ChainId]core.ProposalState
}

// NewListener creates and returns a listener
func NewListener(conn Connection, chainDB *chains.ChainDB, cfg *Config, log log.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *listener {
	return &listener{
		cfg:                *cfg,
		chainDB:            chainDB,
		conn:               conn,
		log:                log,
		stop:               stop,
		sysErr:             sysErr,
		latestBlock:        metrics.LatestBlock{LastUpdated: time.Now()},
		metrics:            m,
		blockConfirmations: cfg.blockConfirmations,
	}
}

func (l *listener) setWriter(w *writer) {
	l.writer = w
}

// start registers all subscriptions provided by the config
func (l *listener) start() error {
	l.log.Debug("Starting listener...")

	go func() {
		err := l.pollBlocks()
		if err != nil {
			l.log.Error("Polling blocks failed", "err", err)
		}
	}()

	if l.cfg.commitNode {
		go func() {
			err := l.commitVotes_loop()
			if err != nil {
				l.log.Error("start  commitVotes failed", "err", err)
			}
		}()
	}

	return nil
}

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.cfg.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *listener) pollBlocks() error {
	var currentBlock = l.cfg.startBlock
	l.log.Info("Polling Blocks...", "startBlock", currentBlock.String())
	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:

			latestBlock, err := l.conn.LatestBlock()
			if err != nil {
				l.log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				time.Sleep(BlockRetryInterval)
				continue
			}

			if l.metrics != nil {
				l.metrics.LatestKnownBlock.Set(float64(latestBlock.Int64()))
			}

			// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(l.blockConfirmations) == -1 {
				l.log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
				time.Sleep(BlockRetryInterval)
				continue
			}

			if currentBlock.Uint64()%100 == 10 {
				l.log.Info("Polling Blocks...", "currentBlock", currentBlock.String())
			}

			err = l.getConcernedEventsForBlock(currentBlock)
			if err != nil {
				continue
			}

			if l.metrics != nil {
				l.metrics.BlocksProcessed.Inc()
				l.metrics.LatestProcessedBlock.Set(float64(latestBlock.Int64()))
			}

			l.latestBlock.Height = big.NewInt(0).Set(latestBlock)
			l.latestBlock.LastUpdated = time.Now()

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			err = l.chainDB.UpdateNextPollBlockNum(uint8(l.cfg.id), address.PubkeyToAddress(l.conn.Keypair().GetPublicKey()).String(), currentBlock.Uint64())
			if err != nil {
				l.log.Error("Failed to write latest block to blockstore", "block", currentBlock, "err", err)
			}
		}
	}
}

func IsTriggerSmartContract(tx api.TransactionExtention) bool {
	if tx.Transaction == nil || tx.Transaction.RawData == nil {
		return false
	}
	if len(tx.Transaction.RawData.Contract) == 0 {
		return false
	}
	for _, contract := range tx.Transaction.RawData.Contract {
		if contract.Parameter.TypeUrl == "type.googleapis.com/protocol.TriggerSmartContract" {
			return true
		}
	}
	return false
}

func (l *listener) IsTriggerBridgeContract(tx api.TransactionExtention) bool {
	if !IsTriggerSmartContract(tx) {
		return false
	}
	for _, contract := range tx.Transaction.RawData.Contract {
		if contract.Parameter.TypeUrl == "type.googleapis.com/protocol.TriggerSmartContract" {
			var reply tronCore.TriggerSmartContract
			err := reply.XXX_Unmarshal(tx.Transaction.RawData.Contract[0].Parameter.Value)
			if err != nil {
				return true
			} else {
				if common.EncodeCheck(reply.GetContractAddress()) == l.cfg.bridgeContract {
					return true
				}
			}
		}
	}
	return false
}

// getConcernedEventsForBlock looks for the deposit event and signature envent in the latest block
func (l *listener) getConcernedEventsForBlock(latestBlock *big.Int) error {

	blockInfo, err := l.conn.Client().GetBlockInfoByNum(latestBlock.Int64())

	if err != nil {
		l.log.Error("Failed get blockInfo", "block", latestBlock, "err", err)
		return err
	}
	//txs := block.GetTransactions()
	txinfos := []*tron.TransactionInfo{}
	for _, tx := range blockInfo.GetTransactionInfo() {
		if common.EncodeCheck(tx.GetContractAddress()) == l.cfg.bridgeContract {
			txinfos = append(txinfos, tx)
		}

	}
	voteReqMsgs := []types.FTTransfer{}
	for _, txinfo := range txinfos {
		logs := txinfo.GetLog()
		for _, log := range logs {
			if common.ToHex(log.Topics[0]) == utils.Deposit.GetTopic() {
				destId := types.ChainId(new(big.Int).SetBytes(log.Data[:32]).Uint64())
				nonce := types.Nonce(new(big.Int).SetBytes(log.Data[64:]).Uint64())
				m, err := l.handleTrc20DepositedEvent(latestBlock.Uint64(), destId, nonce)

				if err != nil {
					return err
				}
				voteReqMsgs = append(voteReqMsgs, m)

			}

		}

	}
	if len(voteReqMsgs) > 0 {
		for _, m := range voteReqMsgs {
			if l.state[m.DestinationId].IsWithCollector() {
				err = l.chainDB.AddSignReq(m)
				if err != nil {
					l.log.Error("chainDB: failed to add sign req", "err", err)
					return err
				}
			} else {
				err = l.chainDB.AddProposalMsg(m)
				if err != nil {
					l.log.Error("chainDB: failed to add sign req", "err", err)
					return err
				}
			}

		}
	}

	return nil
}
