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

func (w *writer) signDestProposal(m types.TransferMsg, bridge []byte) (*seroTypes.Transaction, error) {
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

	recipient := m.Payload[1]

	//var receiptAddr := m.Recipient
	if len(recipient) == 96 {
		short := common.GenContractAddress(seroCommon.BytesToAddress(recipient[:]))
		copy(recipientAddr[:], short[:])
	} else if len(recipient) == 21 {
		copy(recipientAddr[:], recipient[1:])

	} else {
		copy(recipientAddr[:], recipient)
	}

	data = append(data, recipientAddr[:]...)
	data = append(data, seroCommon.LeftPadBytes(m.Payload[0], 32)...)
	if m.Type == types.NonFungibleTransfer {
		w.log.Info("matadata", "data", hexutil.Encode(m.Payload[2]))
		metadata := m.Payload[2]
		data = append(data, metadata...)
		data = append(data, seroCommon.LeftPadBytes(m.Payload[3], 32)...)

	}

	dataHash := common.Hash(data)
	w.log.Info("dataHash", "data", hexutil.Encode(data[:]), "hash", hexutil.Encode(dataHash[:]))

	sig, err := ethCrypto.Sign(dataHash[:], w.conn.Keypair().PrivateKey())

	w.log.Info("signature", "sin", hexutil.Encode(sig[:]))

	if err != nil {
		w.log.Error("ethCrypto.Sign failed", "err", err)
		return nil, err
	}

	return w.commitProposalSignature(m, sig)

}

// commitErc20ProposalSignature submits a proposal with signature
// a  proposal signature will try to be submitted up to the TxRetryLimit times
func (w *writer) commitProposalSignature(m types.TransferMsg, sign []byte) (*seroTypes.Transaction, error) {
	recipient := m.Payload[1]
	amountOrTokenId := new(big.Int).SetBytes(m.Payload[0])
	transferType := uint8(1)

	if m.Type == types.FungibleTransfer {
		w.log.Info("commitFungibleProposalSignature",
			"src", m.SourceId,
			"dest", m.DestinationId,
			"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]),
			"reciepient", hexutil.Encode(recipient), "amount", amountOrTokenId.String(),
			"sign", hexutil.Encode(sign))
	} else {
		transferType = uint8(2)
		w.log.Info("commitNonFungibleProposalSignature",
			"src", m.SourceId,
			"dest", m.DestinationId,
			"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]),
			"reciepient", hexutil.Encode(recipient), "tokenId", amountOrTokenId.String(),
			"metaData", string(m.Payload[2]),
			"sign", hexutil.Encode(sign))
	}

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
		transferType,
		sign,
	)
	w.conn.UnlockOpts()

	if err == nil {
		w.log.Info("Submitted proposal sign", "tx", tx.Hash(),
			"src", m.SourceId,
			"dest", m.DestinationId,
			"depositNonce", m.DepositNonce, "resourceId", hexutil.Encode(m.ResourceId[:]))
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
