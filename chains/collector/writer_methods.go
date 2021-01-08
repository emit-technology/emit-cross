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

package collector

import (
	"errors"
	"github.com/emit-technology/emit-cross/chains"
	"time"

	"github.com/emit-technology/emit-cross/common"

	"github.com/ChainSafe/chainbridge-utils/msg"
)

// Number of blocks to wait for an finalization event
const ExecuteBlockWatchLimit = 100

// Time between retrying a failed tx
const TxRetryInterval = time.Second * 2

// Maximum number of tx retries before exiting
const TxRetryLimit = 10

//Update eth gas price interval
const UpdateEthGasPriceInterval = time.Minute * 5

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")
var ErrFatalTx = errors.New("submission of transaction failed")
var ErrFatalQuery = errors.New("query of chain state failed")

// proposalIsComplete returns true if the proposal state is either Passed, Transferred or Cancelled
func (w *writer) proposalIsComplete(srcId msg.ChainId, dest msg.ChainId, nonce msg.Nonce) bool {
	_, r, a, e := w.state[uint8(srcId)].GetDepositRecord(uint64(nonce), uint8(dest))
	if e != nil {
		w.log.Error("Failed to check deposit existence", "err", e)
		return false
	}
	dataHash := w.state[uint8(dest)].PropsalDataHash(r, a.Bytes())
	prop, err := w.state[uint8(dest)].GetProposalStatus(uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop == PassedStatus || prop == TransferredStatus || prop == CancelledStatus
}

// proposalIsComplete returns true if the proposal state is Transferred or Cancelled
func (w *writer) proposalIsFinalized(srcId msg.ChainId, dest msg.ChainId, nonce msg.Nonce) bool {
	_, r, a, e := w.state[uint8(srcId)].GetDepositRecord(uint64(nonce), uint8(dest))
	if e == nil {
		w.log.Error("Failed to check deposit existence", "err", e)
		return false
	}
	dataHash := w.state[uint8(dest)].PropsalDataHash(r, a.Bytes())
	prop, err := w.state[uint8(dest)].GetProposalStatus(uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop == TransferredStatus || prop == CancelledStatus // Transferred (3)
}

// hasSigned checks if this relayer has already voted
func (w *writer) hasSigned(dest msg.ChainId, nonce msg.Nonce) bool {

	flag, err := w.signatureCollectorContract.HasSignedProposal(w.conn.CallOpts(), uint8(dest),
		uint64(nonce),
		common.GenCommonAddress(w.conn.Keypair()))

	if err != nil {

		w.log.Error("Failed to check proposal HasVote", "err", err)

		return false
	}

	return flag

}

func (w *writer) shouldSign(m chains.Proposal) bool {
	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Source, m.Destination, m.DepositNonce) {
		w.log.Info("Proposal complete, not signing", "src", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
		return false
	}

	// Check if relayer has previously voted
	if w.hasSigned(m.Destination, m.DepositNonce) {
		w.log.Info("Relayer has already sign, not signing", "src", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
		return false
	}

	return true
}
