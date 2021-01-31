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
	"errors"
	"github.com/emit-technology/emit-cross/queue"
	"github.com/emit-technology/emit-cross/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"time"
)

func (l *listener) commitVotes() error {
	lastId, err := l.chainDB.GetLastBatchVotesId(uint8(l.cfg.id), l.conn.Keypair().Address())

	if err != nil {
		l.log.Warn("GetLastBatchVotesId failed", "err", err)
		l.sysErr <- err
	}

	nextId := lastId + 1

	l.log.Info("start commitVotes...", "startWith", nextId)

	for {
		select {
		case <-l.stop:
			return errors.New("commitVotes terminated")
		default:
			m, err := l.chainDB.GetBatchVotesById(uint8(l.cfg.id), nextId)
			if err != nil {
				if err != queue.ErrEmpty && err != queue.ErrOutOfBounds {
					l.log.Error("Failed to get bachVotes by Id", "id", nextId, "err", err)
				}
				time.Sleep(BlockRetryInterval)
				continue
			}
			if l.writer.shouldVote(*m) {

				tx, err := l.writer.commitVotes(*m)
				if err != nil {
					l.log.Error("Failed to commitVotes", "id", nextId, "err", err)
					time.Sleep(BlockRetryInterval)
					continue
				}
				go l.wacthCommitVotesResult(*m, tx.Hash())
			}
			l.chainDB.UpdateLastBatchVotesId(uint8(l.cfg.id), l.conn.Keypair().Address(), nextId)
			nextId++
		}
	}

}

func (l *listener) wacthCommitVotesResult(m types.BatchVotes, tx ethcommon.Hash) {
	begin := time.Now().Unix()
	dataHash := ConstructErc20ProposalDataHash(l.cfg.erc20HandlerContract, ethcommon.BytesToAddress(m.Recipient), m.Amount)
	for {
		if (time.Now().Unix() - begin) > int64(WatchDuration) {
			l.log.Info("commit votes not blocked,retry", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			l.chainDB.AddBatchVotesMsg(m)
			return
		}
		time.Sleep(30 * time.Second)
		receipt, err := l.conn.Client().TransactionReceipt(context.Background(), tx)
		if err != nil {
			l.log.Warn("wacthCommitVoteSResult get Transaction Receipt failed", "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce, "err", err)
			continue
		}
		if receipt == nil {

			continue
		}

		if l.writer.proposalIsFinalized(m.SourceId, m.DepositNonce, dataHash) {
			l.log.Info("Proposal finalized on chain", "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		}

		if receipt.Status == 1 {
			l.log.Info("commit votes success", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		} else {
			l.log.Error("commit votes failed", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		}

	}

}
