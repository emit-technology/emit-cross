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

package sero

import (
	"context"
	"errors"
	"fmt"
	"github.com/emit-technology/emit-cross/common"
	"github.com/emit-technology/emit-cross/core"
	log "github.com/emit-technology/emit-cross/log"
	metrics "github.com/emit-technology/emit-cross/metrics/types"
	"github.com/emit-technology/emit-cross/types"
	"math/big"
	"time"

	"github.com/sero-cash/go-sero"
	seroCommon "github.com/sero-cash/go-sero/common"

	"github.com/emit-technology/emit-cross/chains"
	utils "github.com/emit-technology/emit-cross/shared/sero"
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

	go func() {
		err := l.signDestProposal()
		if err != nil {
			l.log.Error("start signDestProposal failed", "err", err)
		}
	}()

	go func() {
		err := l.voteProposal()
		if err != nil {
			l.log.Error("start voteProposal failed", "err", err)
		}
	}()

	if l.cfg.commitNode {
		go func() {
			err := l.excueteProposal()
			if err != nil {
				l.log.Error("start	 excueteProposal failed", "err", err)
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
				l.log.Error("Failed to get events for block", "block", currentBlock, "err", err)
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
			err = l.chainDB.UpdateNextPollBlockNum(uint8(l.cfg.id), common.GenCommonAddress(l.conn.Keypair()).String(), currentBlock.Uint64())
			if err != nil {
				l.log.Error("Failed to write latest block to blockstore", "block", currentBlock, "err", err)
			}
		}
	}
}

// getConcernedEventsForBlock looks for the deposit event and signature envent in the latest block
func (l *listener) getConcernedEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying block for deposit events", "block", latestBlock)

	query := buildQuery([]seroCommon.Address{l.cfg.bridgeContract, l.cfg.nftBridgeContract, l.cfg.signatureColletorContact},
		latestBlock, latestBlock, utils.Deposit, utils.ProposalEvent, utils.SignProposalEvent)

	// querying for logs
	logs, err := l.conn.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}
	signProposalReqMsgs := []types.TransferMsg{}

	voteProposalReqMsgs := []types.TransferMsg{}

	proposalExcuteMsgs := []types.TransferMsg{}

	commitSingaturesMsgs := []types.ProposalSignatures{}

	// read through the log events and handle their deposit event if Handler is recognized
	for _, log := range logs {

		if log.Address == l.cfg.bridgeContract || log.Address == l.cfg.nftBridgeContract {
			if log.Topics[0] == utils.Deposit.GetTopic() {
				destId := types.ChainId(log.Topics[1].Big().Uint64())
				nonce := types.Nonce(log.Topics[3].Big().Uint64())
				var m types.TransferMsg
				var err error
				if log.Address == l.cfg.bridgeContract {
					m, err = l.handleSrc20DepositedEvent(latestBlock.Uint64(), destId, nonce)
				} else {
					m, err = l.handleSrc721DepositedEvent(latestBlock.Uint64(), destId, nonce)
				}

				if err != nil {
					return err
				}

				if l.state[m.DestinationId].IsWithCollector() {
					signProposalReqMsgs = append(signProposalReqMsgs, m)
				} else {
					voteProposalReqMsgs = append(voteProposalReqMsgs, m)
				}

			}

			if log.Topics[0] == utils.ProposalEvent.GetTopic() {
				sourceId := log.Topics[1].Big().Uint64()
				depositNonce := log.Topics[2].Big().Uint64()
				status := log.Topics[3].Big().Uint64()
				if utils.IsPassed(uint8(status)) {
					var m types.TransferMsg
					var err error
					if log.Address == l.cfg.bridgeContract {
						m, err = l.handleSrc20ProposalEvent(latestBlock.Uint64(), types.ChainId(sourceId), types.Nonce(depositNonce))
					} else {
						m, err = l.handleSrc721ProposalEvent(latestBlock.Uint64(), types.ChainId(sourceId), types.Nonce(depositNonce))

					}
					if err != nil {
						return err
					}

					proposalExcuteMsgs = append(proposalExcuteMsgs, m)

				}

			}

		}

		if log.Address == l.cfg.signatureColletorContact {
			sourceId := log.Topics[1].Big().Uint64()
			destId := log.Topics[2].Big().Uint64()
			depositNonce := log.Topics[3].Big().Uint64()
			transferType := new(big.Int).SetBytes(log.Data[:32]).Uint64()
			status := new(big.Int).SetBytes(log.Data[32:]).Uint64()
			if utils.IsPassed(uint8(status)) {
				var m types.ProposalSignatures
				var err error
				if transferType == 1 {
					_, m, err = l.handleDestFungibleProposalSignEvent(latestBlock.Uint64(), types.ChainId(sourceId), types.ChainId(destId), types.Nonce(depositNonce))

				} else {
					_, m, err = l.handleDestNonFungibleProposalSignEvent(latestBlock.Uint64(), types.ChainId(sourceId), types.ChainId(destId), types.Nonce(depositNonce))

				}
				if err != nil {
					return err
				}

				commitSingaturesMsgs = append(commitSingaturesMsgs, m)
			}
		}

		if len(signProposalReqMsgs) > 0 {
			for _, m := range signProposalReqMsgs {
				err = l.chainDB.AddSignReq(m)
				if err != nil {
					l.log.Error("chainDB: failed to add sign req", "err", err)
					return err
				}
			}
		}

		if len(voteProposalReqMsgs) > 0 {
			for _, m := range voteProposalReqMsgs {
				err = l.chainDB.AddProposalMsg(m)
				if err != nil {
					l.log.Error("chainDB: failed to add vote req", "err", err)
					return err
				}
			}
		}

		if len(proposalExcuteMsgs) > 0 {
			for _, m := range proposalExcuteMsgs {
				err = l.chainDB.AddExecuteMsg(m)
				if err != nil {
					l.log.Error("chainDB: failed to add  ExecuteMsg", "err", err)
					return err
				}
			}
		}
		if len(commitSingaturesMsgs) > 0 {
			for _, m := range commitSingaturesMsgs {
				err = l.chainDB.AddBatchVotesMsg(m)
				if err != nil {
					l.log.Error("chainDB: failed to add  batchVotesMsg", "err", err)
					return err
				}
			}
		}

	}

	return nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contracts []seroCommon.Address, startBlock *big.Int, endBlock *big.Int, sigs ...utils.EventSig) sero.FilterQuery {
	topics := []seroCommon.Hash{}
	for _, sig := range sigs {
		topics = append(topics, sig.GetTopic())
	}
	query := sero.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: contracts,
		Topics: [][]seroCommon.Hash{
			topics,
		},
	}
	return query
}

// proposalIsComplete returns true if the proposal state is either Passed, Transferred or Cancelled
func (l *listener) DestProposalIsComplete(srcId types.ChainId, dest types.ChainId, nonce types.Nonce, typ types.TransferType) bool {
	var dataHash [32]byte
	var prop uint8
	var err error
	if typ == types.FungibleTransfer {
		_, r, a, e := l.state[srcId].GetDepositFTRecord(uint64(nonce), uint8(dest))
		if e != nil {
			l.log.Error("Failed to check deposit existence", "err", e)
			return false
		}
		dataHash = l.state[dest].FTPropsalDataHash(r, a)
		prop, err = l.state[dest].GetFTProposalStatus(uint8(srcId), uint64(nonce), dataHash)

	} else {
		_, r, a, metadata, feeAmount, e := l.state[srcId].GetDepositNFTRecord(uint64(nonce), uint8(dest))
		if e != nil {
			l.log.Error("Failed to check deposit existence", "err", e)
			return false
		}
		dataHash = l.state[dest].NFTPropsalDataHash(r, a, metadata, feeAmount)
		prop, err = l.state[dest].GetNFTProposalStatus(uint8(srcId), uint64(nonce), dataHash)

	}

	if err != nil {
		l.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop == PassedStatus || prop == TransferredStatus || prop == CancelledStatus
}

func (l *listener) shouldSign(m types.TransferMsg) bool {
	if l.DestProposalIsComplete(m.SourceId, m.DestinationId, m.DepositNonce, m.Type) {
		l.log.Info("Proposal On dest chain complete", "typ", m.Type, "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
		return false
	}

	// Check if relayer has previously voted
	if l.writer.hasSigned(m.SourceId, m.DestinationId, m.DepositNonce) {
		l.log.Info("Relayer has already sign", "typ", m.Type, "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
		return false
	}

	return true
}
