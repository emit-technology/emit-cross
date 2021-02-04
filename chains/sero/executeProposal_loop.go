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

func (l *listener) excueteProposal() error {
	nextId := l.cfg.ExecuteProposalStartSeq.Uint64()
	l.log.Info("start excueteProposal...", "startWith", nextId)

	for {
		select {
		case <-l.stop:
			return errors.New("excueteProposal terminated")
		default:

			m, err := l.chainDB.GetExecuteMsgById(uint8(l.cfg.id), nextId)
			if err != nil {
				if err != queue.ErrEmpty && err != queue.ErrOutOfBounds {
					l.log.Error("Failed to get executeMsg by Id", "id", nextId, "err", err)
				}
				time.Sleep(BlockRetryInterval)
				continue
			}
			dataHash := ConstructSrc20ProposalDataHash(l.cfg.src20HandlerContract, seroCommon.BytesToAddress(m.Recipient), m.Amount)

			if !l.writer.proposalIsFinalized(m.SourceId, m.DepositNonce, dataHash) {
				tx, err := l.writer.executeProposal(*m)
				if err != nil {
					l.log.Error("Failed to executeProposal", "id", nextId, "err", err)
					time.Sleep(BlockRetryInterval)
					continue
				}
				go l.wacthExecuteProposalResult(*m, tx.Hash())
			}
			l.chainDB.UpdateLastExecuteId(uint8(l.cfg.id), common.GenCommonAddress(l.conn.Keypair()).String(), nextId)
			nextId++
		}
	}

}

func (l *listener) wacthExecuteProposalResult(m types.FTTransfer, tx seroCommon.Hash) {
	begin := time.Now().Unix()
	dataHash := ConstructSrc20ProposalDataHash(l.cfg.src20HandlerContract, seroCommon.BytesToAddress(m.Recipient), m.Amount)
	for {
		if (time.Now().Unix() - begin) > int64(WatchDuration) {
			l.log.Info("execute proposal not blocked,retry", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			l.chainDB.AddExecuteMsg(m)
			return
		}
		time.Sleep(BlockRetryInterval)

		receipt, err := l.conn.Client().TransactionReceipt(context.Background(), tx)
		if err != nil {
			l.log.Warn("wacthExecuteProposalResult get Transaction Receipt failed", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce, "err", err)
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
			l.log.Info("execute proposal success", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		} else {
			l.log.Error("execute proposal failed", "tx", tx.String(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		}

	}

}
