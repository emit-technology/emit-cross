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

package ethereum

import (
	"context"
	"github.com/ChainSafe/chainbridge-utils/msg"
	utils "github.com/emit-technology/emit-cross/shared/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sero-cash/go-sero/common/hexutil"
	"math/big"
	"time"
)

// createErc20Proposal creates an Erc20 proposal.
// Returns true if the proposal is successfully created or is complete
func (w *writer) commitErc20ProposalSign(m msg.Message) bool {
	dataHash := ConstructErc20ProposalDataHash(w.cfg.erc20HandlerContract, m.Payload[0].([]byte), m.Payload[1].([]byte))

	if w.proposalIsFinalized(m.Source, m.DepositNonce, dataHash) {
		w.log.Info("not need commitVotes, proposal is finalized", "src", m.Source,
			"dst", m.Destination,
			"nonce", m.DepositNonce,
			"recipient", hexutil.Encode(m.Payload[1].([]byte)),
			"amount", new(big.Int).SetBytes(m.Payload[0].([]byte)).String())
		return false
	}

	w.log.Info("add erc20 proposal signature", "src", m.Source,
		"nonce", m.DepositNonce,
		"recipient", hexutil.Encode(m.Payload[1].([]byte)),
		"amount", new(big.Int).SetBytes(m.Payload[0].([]byte)).String())

	new := w.pendingVote.Add(m)
	if new {
		latestBlock, err := w.conn.LatestBlock()
		if err != nil {
			w.log.Error("Unable to fetch latest block", "err", err)
			return false
		}
		go w.waitThenCommit(m, dataHash, latestBlock)
	}
	return true

}

func (s *writer) waitThenCommit(m msg.Message, dataHash [32]byte, latestBlock *big.Int) {

	relayerThreshold, _ := s.bridgeContract.RelayerThreshold(&bind.CallOpts{})

	votes := s.pendingVote.GetPending(m)

	for {
		if s.proposalIsComplete(m.Source, m.DepositNonce, dataHash) {
			return
		}

		if uint64(len(votes)) >= relayerThreshold.Uint64() {
			break
		} else {
			time.Sleep(time.Second)
			votes = s.pendingVote.GetPending(m)
			continue
		}
	}
	latestBlock, err := s.conn.LatestBlock()
	if err != nil {
		s.log.Error("Unable to fetch latest block", "err", err)
		return
	}
	signs := [][]byte{}
	for _, v := range votes {
		signs = append(signs, v)
	}
	go s.watchExcuetResult(m, dataHash, latestBlock)
	s.commitVotes(m, dataHash, signs)
	return

}

// commitVotes submits all votes  for a  proposal
// a commit will try to be submitted up to the TxRetryLimit times
func (w *writer) commitVotes(m msg.Message, dataHash [32]byte, signs [][]byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.conn.LockAndUpdateOpts()
			if err != nil {
				w.log.Error("Failed to update tx opts", "err", err)
				continue
			}

			tx, err := w.bridgeContract.CommitVotes(
				w.conn.Opts(),
				uint8(m.Source),
				uint64(m.DepositNonce),
				m.ResourceId,
				common.BytesToAddress(m.Payload[1].([]byte)),
				big.NewInt(0).SetBytes(m.Payload[0].([]byte)),
				signs,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("commit erc20 proposal vote signatures",
					"tx", tx.Hash(),
					"src", m.Source,
					"depositNonce", m.DepositNonce,
					"recipient", hexutil.Encode(m.Payload[1].([]byte)),
					"amount", new(big.Int).SetBytes(m.Payload[0].([]byte)).String())
				if w.metrics != nil {
					w.metrics.VotesSubmitted.Inc()
				}
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Voting failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for voting, otherwise no need to retry
			if w.proposalIsComplete(m.Source, m.DepositNonce, dataHash) {
				w.log.Info("Proposal voting complete on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.pendingVote.Delete(m)
	w.log.Error("Submission of Vote transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}

// watchExcuetResult watches for the latest block
func (s *writer) watchExcuetResult(m msg.Message, dataHash [32]byte, latestBlock *big.Int) {

	// watching for the latest block, querying and matching the finalized event will be retried up to ExecuteBlockWatchLimit times
	for i := 0; i < ExecuteBlockWatchLimit; i++ {
		select {
		case <-s.stop:
			return
		default:
			// watch for the lastest block, retry up to BlockRetryLimit times
			for waitRetrys := 0; waitRetrys < BlockRetryLimit; waitRetrys++ {
				err := s.conn.WaitForBlock(latestBlock)
				if err != nil {
					s.log.Error("Waiting for block failed", "err", err)
					// Exit if retries exceeded
					if waitRetrys+1 == BlockRetryLimit {
						s.log.Error("Waiting for block retries exceeded, shutting down")
						return
					}
				} else {
					break
				}
			}

			// query for logs
			query := buildQuery(s.cfg.bridgeContract, utils.ProposalEvent, latestBlock, latestBlock)
			evts, err := s.conn.Client().FilterLogs(context.Background(), query)
			if err != nil {
				s.log.Error("Failed to fetch logs", "err", err)
				return
			}

			// execute the proposal once we find the matching finalized event
			for _, evt := range evts {
				sourceId := evt.Topics[1].Big().Uint64()
				depositNonce := evt.Topics[2].Big().Uint64()
				status := evt.Topics[3].Big().Uint64()

				if m.Source == msg.ChainId(sourceId) &&
					m.DepositNonce.Big().Uint64() == depositNonce &&
					utils.IsExecuted(uint8(status)) {
					s.log.Info("Proposal finalized on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
					return
				} else {
					s.log.Trace("Ignoring event", "src", sourceId, "nonce", depositNonce)
				}
			}
			s.log.Trace("No finalization event found in current block", "block", latestBlock, "src", m.Source, "nonce", m.DepositNonce)
			latestBlock = latestBlock.Add(latestBlock, big.NewInt(1))
		}
	}
	s.log.Warn("Block watch limit exceeded,not found finalized even", "source", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
}
