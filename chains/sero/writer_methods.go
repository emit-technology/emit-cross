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
	"errors"
	"github.com/emit-technology/emit-cross/common"
	"github.com/emit-technology/emit-cross/types"
	seroCommon "github.com/sero-cash/go-sero/common"
	seroTypes "github.com/sero-cash/go-sero/core/types"

	"github.com/sero-cash/go-sero/common/hexutil"
)

// Number of blocks to wait for an finalization event
const ExecuteBlockWatchLimit = 100

const BlockRetryLimit = 5

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")
var ErrFatalTx = errors.New("submission of transaction failed")
var ErrFatalQuery = errors.New("query of chain state failed")

// proposalIsComplete returns true if the proposal state is either Passed, Transferred or Cancelled
func (w *writer) proposalIsComplete(srcId types.ChainId, nonce types.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == PassedStatus || prop.Status == TransferredStatus || prop.Status == CancelledStatus
}

// proposalIsComplete returns true if the proposal state is Transferred or Cancelled
func (w *writer) proposalIsFinalized(srcId types.ChainId, nonce types.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == TransferredStatus || prop.Status == CancelledStatus // Transferred (3)
}

// hasVoted checks if this relayer has already voted
func (w *writer) hasVoted(srcId types.ChainId, nonce types.Nonce, dataHash [32]byte) bool {
	hasVoted, err := w.bridgeContract.HasVotedOnProposal(w.conn.CallOpts(), common.IDAndNonce(srcId, nonce), dataHash, common.GenCommonAddress(w.conn.Keypair()))

	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}

	return hasVoted
}

func (w *writer) shouldVote(m types.FTTransfer) bool {
	dataHash := ConstructSrc20ProposalDataHash(w.cfg.src20HandlerContract, seroCommon.BytesToAddress(m.Recipient), m.Amount)

	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.SourceId, m.DepositNonce, dataHash) {
		w.log.Info("Proposal complete, not voting", "src", m.SourceId, "nonce", m.DepositNonce)
		return false
	}

	// Check if relayer has previously voted
	if w.hasVoted(m.SourceId, m.DepositNonce, dataHash) {
		w.log.Info("Relayer has already voted, not voting", "src", m.SourceId, "nonce", m.DepositNonce)
		return false
	}

	return true
}

// voteProposal submits a vote proposal
// a vote proposal will try to be submitted up to the TxRetryLimit times
func (w *writer) voteProposal(m types.FTTransfer) (*seroTypes.Transaction, error) {

	dataHash := ConstructSrc20ProposalDataHash(w.cfg.src20HandlerContract, seroCommon.BytesToAddress(m.Recipient), m.Amount)

	err := w.conn.LockAndUpdateOpts()
	if err != nil {
		w.log.Error("Failed to update tx opts", "err", err)
		return nil, err
	}

	tx, err := w.bridgeContract.VoteProposal(
		w.conn.Opts(),
		uint8(m.SourceId),
		uint64(m.DepositNonce),
		m.ResourceId,
		dataHash,
	)
	w.conn.UnlockOpts()

	if err == nil {
		w.log.Info("Submitted proposal vote", "tx", tx.Hash(), "src", m.SourceId, "depositNonce", m.DepositNonce)
		if w.metrics != nil {
			w.metrics.VotesSubmitted.Inc()
		}
		return tx, nil
	} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
		w.log.Debug("Nonce too low, will retry")
	} else {
		w.log.Warn("Voting failed", "source", m.SourceId, "dest", m.DestinationId, "depositNonce", m.DepositNonce, "err", err)
	}
	return nil, err

}

// executeProposal executes the proposal
func (w *writer) executeProposal(m types.FTTransfer) (*seroTypes.Transaction, error) {

	err := w.conn.LockAndUpdateOpts()
	if err != nil {
		w.log.Error("Failed to update nonce", "err", err)
		return nil, err
	}
	recipeint := seroCommon.BytesToAddress(m.Recipient)
	tx, err := w.bridgeContract.ExecuteProposal(
		w.conn.Opts(),
		uint8(m.SourceId),
		uint64(m.DepositNonce),
		m.ResourceId,
		recipeint,
		m.Amount,
	)
	w.conn.UnlockOpts()

	if err == nil {
		w.log.Info("Submitted proposal execution", "tx", tx.Hash(), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce,
			"recipient", hexutil.Encode(m.Recipient),
			"amount", m.Amount.String())
		return tx, nil
	} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
		w.log.Error("Nonce too low, will retry")
	} else {
		w.log.Warn("Execution failed, proposal may already be complete", "err", err)
	}
	return nil, err
}
