// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"github.com/emit-technology/emit-cross/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sero-cash/go-sero/common/hexutil"
	"math/big"
)

// Number of blocks to wait for an finalization event
const ExecuteBlockWatchLimit = 100

const BlockRetryLimit = 5

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")

// proposalIsComplete returns true if the proposal state is either Passed, Transferred or Cancelled
func (w *writer) proposalIsComplete(srcId types.ChainId, nonce types.Nonce, typ types.TransferType, dataHash [32]byte) bool {
	if typ == types.FungibleTransfer {
		prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
		if err != nil {
			w.log.Error("Failed to check proposal existence", "err", err)
			return false
		}
		return prop.Status == PassedStatus || prop.Status == TransferredStatus || prop.Status == CancelledStatus
	} else {
		prop, err := w.nftBridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
		if err != nil {
			w.log.Error("Failed to check proposal existence", "err", err)
			return false
		}
		return prop.Status == PassedStatus || prop.Status == TransferredStatus || prop.Status == CancelledStatus
	}
}

// proposalIsComplete returns true if the proposal state is Transferred or Cancelled
func (w *writer) proposalIsFinalized(srcId types.ChainId, nonce types.Nonce, typ types.TransferType, dataHash [32]byte) bool {
	if typ == types.FungibleTransfer {
		prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
		if err != nil {
			w.log.Error("Failed to check proposal existence", "err", err)
			return false
		}
		return prop.Status == TransferredStatus || prop.Status == CancelledStatus // Transferred (3)
	} else {
		prop, err := w.nftBridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
		if err != nil {
			w.log.Error("Failed to check proposal existence", "err", err)
			return false
		}
		return prop.Status == TransferredStatus || prop.Status == CancelledStatus // Transferred (3)
	}

}

func (w *writer) shouldCommitSignatures(m types.ProposalSignatures) bool {
	recipient := ethCommon.BytesToAddress(m.Payload[1])
	amountOrTokenId := new(big.Int).SetBytes(m.Payload[0])
	var dataHash [32]byte
	if m.Type == types.FungibleTransfer {
		dataHash = ConstructErc20ProposalDataHash(w.cfg.erc20HandlerContract, recipient, amountOrTokenId)
	} else {
		feeAmount := new(big.Int).SetBytes(m.Payload[3])

		dataHash = ConstructErc721ProposalDataHash(w.cfg.erc721HandlerContract, recipient, amountOrTokenId, m.Payload[2], feeAmount)

	}

	if w.proposalIsComplete(m.SourceId, m.DepositNonce, m.Type, dataHash) {
		w.log.Info("Proposal has completed", "type", m.Type, "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
		return false
	}

	return true
}

// commitVotes submits all votes  for a  proposal
// a commit will try to be submitted up to the TxRetryLimit times
func (w *writer) commitVotes(m types.ProposalSignatures) (*ethTypes.Transaction, error) {
	err := w.conn.LockAndUpdateOpts()
	if err != nil {
		w.log.Error("Failed to update tx opts", "err", err)
		return nil, err
	}
	recipient := ethCommon.BytesToAddress(m.Payload[1])
	amountOrTokenId := new(big.Int).SetBytes(m.Payload[0])
	feeAmount := new(big.Int)
	var tx *ethTypes.Transaction
	if m.Type == types.FungibleTransfer {
		tx, err = w.bridgeContract.CommitVotes(
			w.conn.Opts(),
			uint8(m.SourceId),
			uint64(m.DepositNonce),
			m.ResourceId,
			recipient,
			amountOrTokenId,
			m.Signatures,
		)
	} else {
		feeAmount = new(big.Int).SetBytes(m.Payload[3])
		var data []byte
		data = append(data, ethCommon.LeftPadBytes(recipient[:], 32)...)
		data = append(data, ethCommon.LeftPadBytes(amountOrTokenId.Bytes(), 32)...)
		data = append(data, ethCommon.LeftPadBytes(big.NewInt(128).Bytes(), 32)...)
		data = append(data, ethCommon.LeftPadBytes(m.Payload[3], 32)...)
		metadata := m.Payload[2]
		len := len(metadata)
		data = append(data, ethCommon.LeftPadBytes(new(big.Int).SetInt64(int64(len)).Bytes(), 32)...)
		if len > 32 {
			n := len / 32
			data = append(data, metadata[:n*32]...)
			d := len % 32
			if d > 0 {
				data = append(data, ethCommon.RightPadBytes(metadata[n*32:], 32)...)

			}

		} else {
			data = append(data, ethCommon.RightPadBytes(metadata, 32)...)

		}

		w.log.Info("votes param", "data", hexutil.Encode(data[:]))

		tx, err = w.nftBridgeContract.CommitVotes(
			w.conn.Opts(),
			uint8(m.SourceId),
			uint64(m.DepositNonce),
			m.ResourceId,
			data,
			m.Signatures,
		)
	}

	w.conn.UnlockOpts()

	if err == nil {
		if m.Type == types.FungibleTransfer {
			w.log.Info("commit erc20 proposal vote signatures",
				"tx", tx.Hash(),
				"src", m.SourceId,
				"depositNonce", m.DepositNonce,
				"recipient", recipient,
				"amount", amountOrTokenId.String())
		} else {
			w.log.Info("commit erc721 proposal vote signatures",
				"tx", tx.Hash(),
				"src", m.SourceId,
				"depositNonce", m.DepositNonce,
				"recipient", recipient,
				"tokenId", amountOrTokenId.String(),
				"feeAmount", feeAmount.String())
		}

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
