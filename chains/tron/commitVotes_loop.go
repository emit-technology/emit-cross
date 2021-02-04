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
	"github.com/emit-technology/emit-cross/queue"
	"github.com/emit-technology/emit-cross/types"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	tronCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"time"
)

func (l *listener) repairCommitVotes(seqs []uint64) {

	for _, seq := range seqs {
		m, err := l.chainDB.GetBatchVotesById(uint8(l.cfg.id), seq)
		if err != nil {
			if err != queue.ErrEmpty && err != queue.ErrOutOfBounds {
				l.log.Error("Failed to get executeMsg by Id", "id", seq, "err", err)
				continue
			}

		} else {
			panic(err)
		}
		dataHash := ConstructTrc20ProposalDataHash(l.cfg.trc20HandlerContract, m.Recipient, m.Amount)
		if !l.writer.proposalIsFinalized(m.SourceId, m.DepositNonce, dataHash) {
			l.log.Info("commitVotes", "seq", seq, "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce,
				"recipient", tronCommon.EncodeCheck(m.Recipient),
				"amount", m.Amount.String())
			tx, err := l.writer.commitVotes(*m)
			if err != nil {
				l.log.Error("Failed to executeProposal", "id", seq, "err", err)
				continue
			}
			go l.wacthCommitVotesResult(*m, tx.Txid)
		}

	}

}

func (l *listener) commitVotes_loop() error {
	//lastId, err := l.chainDB.GetLastBatchVotesId(uint8(l.cfg.id), address.PubkeyToAddress(l.conn.Keypair().GetPublicKey()).String())
	nextId := l.cfg.commitVotesStartSeq.Uint64()
	l.log.Info("start commitVotes...", "startWith", l.cfg.commitVotesStartSeq)
	//if err != nil {
	//	panic(err)
	//}
	for {
		select {
		case <-l.stop:
			return errors.New("commitVotes terminated")
		default:
			m, err := l.chainDB.GetBatchVotesById(uint8(l.cfg.id), nextId)
			if err != nil {
				if err != queue.ErrEmpty && err != queue.ErrOutOfBounds {
					l.log.Error("Failed to get executeMsg by Id", "id", nextId, "err", err)
				}
				time.Sleep(BlockRetryInterval)
				continue
			}
			dataHash := ConstructTrc20ProposalDataHash(l.cfg.trc20HandlerContract, m.Recipient, m.Amount)
			if !l.writer.proposalIsFinalized(m.SourceId, m.DepositNonce, dataHash) {
				l.log.Info("commitVotes", "seq", nextId, "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce,
					"recipient", tronCommon.EncodeCheck(m.Recipient),
					"amount", m.Amount.String())
				tx, err := l.writer.commitVotes(*m)
				if err != nil {
					l.log.Error("Failed to executeProposal", "id", nextId, "err", err)
					time.Sleep(BlockRetryInterval)
					continue
				}
				go l.wacthCommitVotesResult(*m, tx.Txid)
			}
			l.chainDB.UpdateLastBatchVotesId(uint8(l.cfg.id), address.PubkeyToAddress(l.conn.Keypair().GetPublicKey()).String(), nextId)
			nextId++
		}
	}

}

func (l *listener) wacthCommitVotesResult(m types.BatchVotes, txId []byte) {
	begin := time.Now().Unix()
	dataHash := ConstructTrc20ProposalDataHash(l.cfg.trc20HandlerContract, m.Recipient, m.Amount)
	for {
		if (time.Now().Unix() - begin) > int64(WatchDuration) {
			l.log.Info("commit votes not blocked,retry", "tx", tronCommon.ToHex(txId), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			l.chainDB.AddBatchVotesMsg(m)
			return
		}
		time.Sleep(BlockRetryInterval)
		receipt, err := l.conn.Client().GetTransactionInfoByID(tronCommon.ToHex(txId))
		if err != nil {
			l.log.Warn("wacthCommitVotesResult get Transaction Receipt failed", "tx", tronCommon.ToHex(txId), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce, "err", err)
			continue
		}
		if receipt == nil {

			continue
		}
		if l.writer.proposalIsFinalized(m.SourceId, m.DepositNonce, dataHash) {
			l.log.Info("Proposal finalized on chain", "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		}
		if receipt.Receipt.Result == 1 {
			l.log.Info("commit votes success", "tx", tronCommon.ToHex(txId), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		} else {
			l.log.Error("commit votes failed", "tx", tronCommon.ToHex(txId), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
			return
		}

	}

}
