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
	"github.com/emit-technology/emit-cross/chains"
	"github.com/emit-technology/emit-cross/common"
	utils "github.com/emit-technology/emit-cross/shared/sero"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	seroCommon "github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/common/hexutil"
	"math/big"
	"time"
)

//signErc20Proposal
// Returns true if the proposal is successfully created or is complete
func (w *writer) signProposal(m chains.Proposal) bool {
	w.log.Info("Ralayer sign proposal", "source", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
	var data []byte
	var bridgeAddress [20]byte
	bridge := w.state[uint8(m.Destination)].GetBridgeAddress()
	if len(bridge) == 96 {
		short := common.GenContractAddress(seroCommon.BytesToAddress(bridge))
		copy(bridgeAddress[:], short[:])
	} else {
		copy(bridgeAddress[:], bridge[:])
	}
	data = append(data, bridgeAddress[:]...)
	data = append(data, new(big.Int).SetUint64(uint64(m.Source)).Bytes()...)
	data = append(data, new(big.Int).SetUint64(uint64(m.Destination)).Bytes()...)
	data = append(data, seroCommon.LeftPadBytes(new(big.Int).SetUint64(uint64(m.DepositNonce)).Bytes(), 8)...)
	data = append(data, m.ResourceId[:]...)
	receiptAddr := m.Recipient
	if len(m.Recipient) == 96 {
		short := common.GenContractAddress(seroCommon.BytesToAddress(m.Recipient[:]))
		copy(receiptAddr[:], short[:])
	}
	data = append(data, receiptAddr...)
	data = append(data, seroCommon.LeftPadBytes(m.Amount.Bytes(), 32)...)

	dataHash := utils.Hash(data)

	if !w.shouldSign(m) {
		return false
	}

	sig, err := ethCrypto.Sign(dataHash[:], w.conn.Keypair().PrivateKey())

	if err != nil {
		w.log.Error("ethCrypto.Sign failed", "err", err)
		return false
	}

	go w.commitProposalSignature(m, sig)

	return true
}

// commitErc20ProposalSignature submits a proposal with signature
// a  proposal signature will try to be submitted up to the TxRetryLimit times
func (w *writer) commitProposalSignature(m chains.Proposal, sign []byte) {
	w.log.Info("commitErc20ProposalSignature",
		"src", m.Source,
		"dest", m.Destination,
		"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]),
		"reciepient", hexutil.Encode(m.Recipient), "amount", m.Amount.String())

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

			tx, err := w.signatureCollectorContract.CommitSign(
				w.conn.Opts(),
				uint8(m.Source),
				uint8(m.Destination),
				uint64(m.DepositNonce),
				sign,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Submitted proposal sign", "tx", tx.Hash(),
					"src", m.Source,
					"dest", m.Destination,
					"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]),
					"reciepient", hexutil.Encode(m.Recipient), "amount", m.Amount.String())
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("commit proposal sign failed", "dest", m.Destination,
					"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]))
				time.Sleep(TxRetryInterval)
			}
		}
	}
	w.log.Error("commit proposal sign failed", "dest", m.Destination,
		"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]))
	w.sysErr <- ErrFatalTx
}
