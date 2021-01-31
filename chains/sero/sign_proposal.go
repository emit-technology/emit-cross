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
	"github.com/emit-technology/emit-cross/common"
	"github.com/emit-technology/emit-cross/types"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	seroCommon "github.com/sero-cash/go-sero/common"
	seroTypes "github.com/sero-cash/go-sero/core/types"

	"github.com/sero-cash/go-sero/common/hexutil"
	"math/big"
)

func (w *writer) signDestProposal(m types.FTTransfer, bridge []byte) (*seroTypes.Transaction, error) {
	w.log.Info("Ralayer sign dest proposal", "source", m.SourceId, "dest", m.DestinationId, "nonce", m.DepositNonce)
	var data []byte
	var bridgeAddress [20]byte
	if len(bridge) == 96 {
		short := common.GenContractAddress(seroCommon.BytesToAddress(bridge))
		copy(bridgeAddress[:], short[:])
	} else {
		if len(bridge) == 21 {
			copy(bridgeAddress[:], bridge[1:])
		} else {
			copy(bridgeAddress[:], bridge[:])

		}
	}
	data = append(data, bridgeAddress[:]...)
	data = append(data, new(big.Int).SetUint64(uint64(m.SourceId)).Bytes()...)
	data = append(data, new(big.Int).SetUint64(uint64(m.DestinationId)).Bytes()...)
	data = append(data, seroCommon.LeftPadBytes(new(big.Int).SetUint64(uint64(m.DepositNonce)).Bytes(), 8)...)
	data = append(data, m.ResourceId[:]...)
	var recipientAddr [20]byte
	//var receiptAddr := m.Recipient
	if len(m.Recipient) == 96 {
		short := common.GenContractAddress(seroCommon.BytesToAddress(m.Recipient[:]))
		copy(recipientAddr[:], short[:])
	} else if len(m.Recipient) == 21 {
		copy(recipientAddr[:], m.Recipient[1:])

	} else {
		copy(recipientAddr[:], m.Recipient)
	}

	data = append(data, recipientAddr[:]...)
	data = append(data, seroCommon.LeftPadBytes(m.Amount.Bytes(), 32)...)

	dataHash := common.Hash(data)
	w.log.Info("dataHash", "hash", hexutil.Encode(dataHash[:]))

	sig, err := ethCrypto.Sign(dataHash[:], w.conn.Keypair().PrivateKey())

	if err != nil {
		w.log.Error("ethCrypto.Sign failed", "err", err)
		return nil, err
	}

	return w.commitProposalSignature(m, sig)

}

// commitErc20ProposalSignature submits a proposal with signature
// a  proposal signature will try to be submitted up to the TxRetryLimit times
func (w *writer) commitProposalSignature(m types.FTTransfer, sign []byte) (*seroTypes.Transaction, error) {
	w.log.Info("commitErc20ProposalSignature",
		"src", m.SourceId,
		"dest", m.DestinationId,
		"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]),
		"reciepient", hexutil.Encode(m.Recipient), "amount", m.Amount.String(),
		"sign", hexutil.Encode(sign))

	err := w.conn.LockAndUpdateOpts()
	if err != nil {
		w.log.Error("Failed to update tx opts", "err", err)
		return nil, err
	}

	tx, err := w.signatureCollectorContract.SignProposal(
		w.conn.Opts(),
		uint8(m.SourceId),
		uint8(m.DestinationId),
		uint64(m.DepositNonce),
		sign,
	)
	w.conn.UnlockOpts()

	if err == nil {
		w.log.Info("Submitted proposal sign", "tx", tx.Hash(),
			"src", m.SourceId,
			"dest", m.DestinationId,
			"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]),
			"reciepient", hexutil.Encode(m.Recipient), "amount", m.Amount.String())
		return tx, nil
	} else {
		w.log.Warn("commit proposal sign failed", "dest", m.DestinationId,
			"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]), "err", err)
		return nil, err
	}

}

// hasSigned checks if this relayer has already voted
func (w *writer) hasSigned(srcId, dest types.ChainId, nonce types.Nonce) bool {

	startNonce, err := w.signatureCollectorContract.GetDestinationStartNoncee(w.conn.CallOpts(), uint8(srcId), uint8(dest))

	if err != nil {
		w.log.Error("Failed to get DestinationStartNonce", "err", err)
		return false
	}
	if uint64(nonce) < startNonce {
		return true
	}

	flag, err := w.signatureCollectorContract.HasSignedProposal(w.conn.CallOpts(), uint8(srcId), uint8(dest),
		uint64(nonce),
		common.GenCommonAddress(w.conn.Keypair()))

	if err != nil {
		return false
	}

	return flag

}
