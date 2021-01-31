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
	"github.com/emit-technology/emit-cross/common"
	"github.com/emit-technology/emit-cross/queue"
	"github.com/emit-technology/emit-cross/types"
	seroCommon "github.com/sero-cash/go-sero/common"
	"time"
)

func (l *listener) voteProposal() error {
	lastId, err := l.chainDB.GetLastVoteId(uint8(l.cfg.id), common.GenCommonAddress(l.conn.Keypair()).String())
	if err != nil {
		panic(err)
	}
	nextId := lastId + 1
	l.log.Info("start voteProposal...", "startWith", nextId)

	for {
		select {
		case <-l.stop:
			return errors.New("voteProposal terminated")
		default:
			m, err := l.chainDB.GetProposalMsgById(uint8(l.cfg.id), nextId)
			if err != nil {
				if err != queue.ErrEmpty && err != queue.ErrOutOfBounds {
					l.log.Error("Failed to get Proposal by Id", "id", nextId, "err", err)
				}
				time.Sleep(BlockRetryInterval)
				continue
			}

			if l.writer.shouldVote(*m) {

				tx, err := l.writer.voteProposal(*m)
				if err != nil {
					l.log.Error("Failed to voteProposal", "id", nextId, "err", err)
					time.Sleep(BlockRetryInterval)
					continue
				}
				go l.wacthVoteProposalResult(*m, tx.Hash())
			}
			l.chainDB.UpdateLastVoteId(uint8(l.cfg.id), common.GenCommonAddress(l.conn.Keypair()).String(), nextId)
			nextId = nextId + 1
		}
	}

}

func (l *listener) wacthVoteProposalResult(m types.FTTransfer, tx seroCommon.Hash) {
	begin := time.Now().Unix()
	for {
		if (time.Now().Unix() - begin) > int64(WatchDuration) {
			l.log.Info("vote proposal not blocked,retry", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			l.chainDB.AddProposalMsg(m)
			return
		}
		time.Sleep(BlockRetryInterval)

		receipt, err := l.conn.Client().TransactionReceipt(context.Background(), tx)
		if err != nil {
			l.log.Warn("wacthVoteProposalResult get Transaction Receipt failed", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce, "err", err)
			continue
		}
		if receipt == nil {
			continue
		}

		if receipt.Status == 1 {
			l.log.Info("commit votes success blocked", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		} else {
			l.log.Error("commit votes failed blocked", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		}

	}

}
