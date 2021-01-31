// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"github.com/emit-technology/emit-cross/common"
	"github.com/emit-technology/emit-cross/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sero-cash/go-sero/common/hexutil"
)

// Number of blocks to wait for an finalization event
const ExecuteBlockWatchLimit = 100

const BlockRetryLimit = 5

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")

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
	hasVoted, err := w.bridgeContract.HasVotedOnProposal(w.conn.CallOpts(), common.IDAndNonce(srcId, nonce), dataHash, w.conn.Opts().From)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}

	return hasVoted
}

func (w *writer) shouldVote(m types.BatchVotes) bool {
	dataHash := ConstructErc20ProposalDataHash(w.cfg.erc20HandlerContract, ethCommon.BytesToAddress(m.Recipient), m.Amount)
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

// commitVotes submits all votes  for a  proposal
// a commit will try to be submitted up to the TxRetryLimit times
func (w *writer) commitVotes(m types.BatchVotes) (*ethTypes.Transaction, error) {
	err := w.conn.LockAndUpdateOpts()
	if err != nil {
		w.log.Error("Failed to update tx opts", "err", err)
		return nil, err
	}
	tx, err := w.bridgeContract.CommitVotes(
		w.conn.Opts(),
		uint8(m.SourceId),
		uint64(m.DepositNonce),
		m.ResourceId,
		ethCommon.BytesToAddress(m.Recipient),
		m.Amount,
		m.Signatures,
	)
	w.conn.UnlockOpts()

	if err == nil {
		w.log.Info("commit erc20 proposal vote signatures",
			"tx", tx.Hash(),
			"src", m.SourceId,
			"depositNonce", m.DepositNonce,
			"recipient", hexutil.Encode(m.Recipient),
			"amount", m.Amount.String())
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
