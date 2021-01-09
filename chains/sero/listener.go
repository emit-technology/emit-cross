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
	"math/big"
	"time"

	"github.com/emit-technology/emit-cross/bindings/sero/SRC20Handler"
	"github.com/sero-cash/go-sero"
	seroCommon "github.com/sero-cash/go-sero/common"

	"github.com/ChainSafe/chainbridge-utils/blockstore"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/emit-technology/emit-cross/bindings/sero/Bridge"
	"github.com/emit-technology/emit-cross/chains"
	utils "github.com/emit-technology/emit-cross/shared/sero"
)

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("listener block polling failed")

type listener struct {
	cfg                  Config
	conn                 Connection
	router               chains.Router
	writer               *writer
	bridgeContract       *Bridge.Bridge // instance of bound bridge contract
	src20HandlerContract *SRC20Handler.SRC20Handler
	log                  log15.Logger
	blockstore           blockstore.Blockstorer
	stop                 <-chan int
	sysErr               chan<- error // Reports fatal error to core
	latestBlock          metrics.LatestBlock
	metrics              *metrics.ChainMetrics
	blockConfirmations   *big.Int
}

// NewListener creates and returns a listener
func NewListener(conn Connection, cfg *Config, log log15.Logger, bs blockstore.Blockstorer, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *listener {
	return &listener{
		cfg:                *cfg,
		conn:               conn,
		log:                log,
		blockstore:         bs,
		stop:               stop,
		sysErr:             sysErr,
		latestBlock:        metrics.LatestBlock{LastUpdated: time.Now()},
		metrics:            m,
		blockConfirmations: cfg.blockConfirmations,
	}
}

// setContracts sets the listener with the appropriate contracts
func (l *listener) setContracts(bridge *Bridge.Bridge, src20Handler *SRC20Handler.SRC20Handler) {
	l.bridgeContract = bridge
	l.src20HandlerContract = src20Handler
}

// sets the router
func (l *listener) setRouter(r chains.Router) {
	l.router = r
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

	return nil
}

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.cfg.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *listener) pollBlocks() error {
	l.log.Info("Polling Blocks...")
	var currentBlock = l.cfg.startBlock
	var retry = BlockRetryLimit
	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				l.log.Error("Polling failed, retries exceeded")
				l.sysErr <- ErrFatalPolling
				return nil
			}

			latestBlock, err := l.conn.LatestBlock()
			if err != nil {
				l.log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				retry--
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

			err = l.getConcernedEventsForBlock(currentBlock)
			if err != nil {
				l.log.Error("Failed to get events for block", "block", currentBlock, "err", err)
				retry--
				continue
			}

			// Write to block store. Not a critical operation, no need to retry
			err = l.blockstore.StoreBlock(currentBlock)
			if err != nil {
				l.log.Error("Failed to write latest block to blockstore", "block", currentBlock, "err", err)
			}

			if l.metrics != nil {
				l.metrics.BlocksProcessed.Inc()
				l.metrics.LatestProcessedBlock.Set(float64(latestBlock.Int64()))
			}

			l.latestBlock.Height = big.NewInt(0).Set(latestBlock)
			l.latestBlock.LastUpdated = time.Now()

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = BlockRetryLimit
		}
	}
}

// getConcernedEventsForBlock looks for the deposit event and signature envent in the latest block
func (l *listener) getConcernedEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying block for deposit events", "block", latestBlock)
	query := buildQuery(l.cfg.bridgeContract, utils.Deposit,
		latestBlock, latestBlock)

	// querying for logs
	logs, err := l.conn.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	// read through the log events and handle their deposit event if handler is recognized
	for _, log := range logs {

		if log.Address == l.cfg.bridgeContract {
			destId := msg.ChainId(log.Topics[1].Big().Uint64())
			nonce := msg.Nonce(log.Topics[3].Big().Uint64())
			m, err := l.handleSrc20DepositedEvent(destId, nonce)

			if err != nil {
				return err
			}
			err = l.router.Send(m)

			if err != nil {
				l.log.Error("subscription error: failed to route message", "err", err)
				return err
			}

		}
	}

	return nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract seroCommon.Address, sig utils.EventSig, startBlock *big.Int, endBlock *big.Int) sero.FilterQuery {
	query := sero.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []seroCommon.Address{contract},
		Topics: [][]seroCommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}
