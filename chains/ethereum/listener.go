// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"errors"
	"fmt"
	"github.com/emit-technology/emit-cross/chains"
	"github.com/emit-technology/emit-cross/core"
	"github.com/emit-technology/emit-cross/log"
	metrics "github.com/emit-technology/emit-cross/metrics/types"
	utils "github.com/emit-technology/emit-cross/shared/ethereum"
	"github.com/emit-technology/emit-cross/types"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

var BlockRetryInterval = time.Second * 5

var WatchDuration = 90 * time.Minute

type listener struct {
	cfg                Config
	chainDB            *chains.ChainDB
	conn               Connection
	writer             *writer
	log                log15.Logger
	stop               <-chan int
	sysErr             chan<- error // Reports fatal error to core
	latestBlock        metrics.LatestBlock
	metrics            *metrics.ChainMetrics
	blockConfirmations *big.Int
	state              map[types.ChainId]core.ProposalState
}

// NewListener creates and returns a listener
func NewListener(conn Connection, chainDB *chains.ChainDB, cfg *Config, log log15.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *listener {
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
			panic(err)
		}
	}()
	if l.cfg.commitNode {
		go func() {
			err := l.commitVotes()
			if err != nil {
				l.log.Error("loop commitVotes failed", "err", err)
				panic(err)
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

			//l.log.Info("pollBlock","chainId",l.cfg.name,"currentBlock",currentBlock.String())
			// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(l.blockConfirmations) == -1 {
				l.log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
				time.Sleep(BlockRetryInterval)
				continue
			}

			if currentBlock.Uint64()%100 == 10 {
				l.log.Info("Polling Blocks...", "currentBlock", currentBlock.String())
			}

			// Parse out events
			err = l.getDepositEventsForBlock(currentBlock)

			if err != nil {
				l.log.Error("Failed to get events for block", "block", currentBlock, "err", err)
				continue
			}

			// Write to block store. Not a critical operation, no need to retry

			if l.metrics != nil {
				l.metrics.BlocksProcessed.Inc()
				l.metrics.LatestProcessedBlock.Set(float64(latestBlock.Int64()))
			}

			l.latestBlock.Height = big.NewInt(0).Set(latestBlock)
			l.latestBlock.LastUpdated = time.Now()

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))

			err = l.chainDB.UpdateNextPollBlockNum(uint8(l.cfg.id), l.conn.Keypair().Address(), currentBlock.Uint64())

			if err != nil {
				l.log.Error("Failed to write next start block", "block", currentBlock, "err", err)
			}
		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying block for deposit events", "block", latestBlock)
	query := buildQuery(l.cfg.bridgeContract, utils.Deposit, latestBlock, latestBlock)

	// querying for logs
	logs, err := l.conn.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	ms := []types.FTTransfer{}
	// read through the log events and handle their deposit event if Handler is recognized
	for _, log := range logs {
		destId := types.ChainId(log.Topics[1].Big().Uint64())
		//rId := msg.ResourceIdFromSlice(log.Topics[2].Bytes())
		nonce := types.Nonce(log.Topics[3].Big().Uint64())

		m, err := l.handleErc20DepositedEvent(latestBlock.Uint64(), destId, nonce)

		if err != nil {
			return err
		}

		ms = append(ms, m)
	}

	if len(ms) > 0 {

		for _, m := range ms {
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

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract ethcommon.Address, sig utils.EventSig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethcommon.Address{contract},
		Topics: [][]ethcommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}
