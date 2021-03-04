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
	"crypto/sha256"
	"encoding/json"
	"errors"
	"github.com/emit-technology/emit-cross/common"
	"github.com/emit-technology/emit-cross/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	tronCommon "github.com/fbsobreira/gotron-sdk/pkg/common"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/sero-cash/go-sero/common/hexutil"
	"google.golang.org/protobuf/proto"
	"math/big"
	"strconv"
	"time"
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

func (w *writer) prosalStatus(srcId types.ChainId, nonce types.Nonce, dataHash [32]byte) (uint8, error) {
	params := []abi.Param{
		{"uint8": strconv.FormatUint(uint64(srcId), 10)},
		{"uint64": strconv.FormatUint(uint64(nonce), 10)},
		{"bytes32": hexutil.Encode(dataHash[:])[2:]},
	}
	paramsJson, err := json.Marshal(params)

	if err != nil {
		w.log.Error("Failed to Marshal getProposal params", "err", err)
		return 0, err
	}

	result, err := w.conn.Client().TriggerConstantContract("", w.cfg.bridgeContract, "getProposal(uint8,uint64,bytes32)", string(paramsJson))
	if err != nil {
		w.log.Error("Failed to TriggerConstantContract getProposal", "err", err)
		return 0, err
	}
	if result.Result.Code > 0 || result.Result.Message != nil {
		w.log.Error("Failed to TriggerConstantContract getProposal", "err", string(result.Result.Message))
		return 0, errors.New(string(result.Result.Message))
	}
	proposal, err := w.bridgeAbi.GetProposal(result.GetConstantResult()[0])
	if err != nil {
		return 0, err
	}
	return proposal.Status, nil
}

// proposalIsComplete returns true if the proposal state is either Passed, Transferred or Cancelled
func (w *writer) proposalIsComplete(srcId types.ChainId, nonce types.Nonce, dataHash [32]byte) bool {
	status, err := w.prosalStatus(srcId, nonce, dataHash)
	if err != nil {
		return false
	}

	return status == PassedStatus || status == TransferredStatus || status == CancelledStatus
}

// proposalIsComplete returns true if the proposal state is Transferred or Cancelled
func (w *writer) proposalIsFinalized(srcId types.ChainId, nonce types.Nonce, dataHash [32]byte) bool {
	status, err := w.prosalStatus(srcId, nonce, dataHash)
	if err != nil {
		return false
	}
	return status == TransferredStatus || status == CancelledStatus // Transferred (3)
}

// hasVoted checks if this relayer has already voted
func (w *writer) hasVoted(srcId types.ChainId, nonce types.Nonce, dataHash [32]byte) bool {

	params := []abi.Param{
		{"uint72": strconv.FormatUint(common.IDAndNonce(srcId, nonce).Uint64(), 10)},
		{"bytes32": hexutil.Encode(dataHash[:])[2:]},
		{"address": address.PubkeyToAddress(w.conn.Keypair().GetPublicKey()).String()},
	}
	paramsJson, err := json.Marshal(params)
	if err != nil {
		w.log.Error("Failed to Marshal hasVotedOnProposal params", "err", err)
		return false
	}

	result, err := w.conn.Client().TriggerConstantContract("", w.cfg.bridgeContract, "_hasVotedOnProposal(uint72,bytes32,address)", string(paramsJson))
	if err != nil {
		w.log.Error("Failed to TriggerConstantContract hasVotedOnProposal", "err", err)
		return false
	}
	if result.Result.Code > 0 || result.Result.Message != nil {
		w.log.Error("Failed to TriggerConstantContract hasVotedOnProposal", "err", string(result.Result.Message))
		return false
	}
	status := uint8(new(big.Int).SetBytes(result.GetConstantResult()[0]).Uint64())
	if status == 1 {
		return true
	}
	return false
}

func (w *writer) shouldVote(m types.TransferMsg) bool {
	recipient := m.Payload[1]
	amount := new(big.Int).SetBytes(m.Payload[0])
	dataHash := ConstructTrc20ProposalDataHash(w.cfg.trc20HandlerContract, recipient, amount)

	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.SourceId, m.DepositNonce, dataHash) {
		w.log.Info("Proposal complete, not voting", "type", m.Type, "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
		return false
	}

	// Check if relayer has previously voted
	if w.hasVoted(m.SourceId, m.DepositNonce, dataHash) {
		w.log.Info("Relayer has already voted, not voting", "type", m.Type, "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce)
		return false
	}

	return true
}

func (w *writer) signTx(tx *core.Transaction) (*core.Transaction, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {

		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	signature, err := crypto.Sign(hash, w.conn.Keypair().PrivateKey())
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)
	return tx, nil
}

func (w *writer) commitVotes(m types.ProposalSignatures) (*api.TransactionExtention, error) {

	signatureStrs := []string{}

	for _, s := range m.Signatures {
		signatureStrs = append(signatureStrs, hexutil.Encode(s)[2:])
	}
	recipient := m.Payload[1]
	amount := new(big.Int).SetBytes(m.Payload[0])

	params := []abi.Param{
		{"uint8": strconv.FormatUint(uint64(m.SourceId), 10)},
		{"uint64": strconv.FormatUint(uint64(m.DepositNonce), 10)},
		{"bytes32": hexutil.Encode(m.ResourceId[:])[2:]},
		{"address": tronCommon.EncodeCheck(recipient)},
		{"uint256": amount.String()},
		{"bytes[]": signatureStrs},
	}

	paramsJson, err := json.Marshal(params)

	if err != nil {
		w.log.Error("commitVotes Failed to Marshal params", "err", err)
		return nil, err
	}

	w.log.Info("commitVotes", "paramJson", string(paramsJson))

	tx, err := w.bridgeAbi.CommitVotes(w.conn.Client(),
		address.PubkeyToAddress(w.conn.Keypair().GetPublicKey()).String(),
		w.cfg.bridgeContract,
		uint8(m.SourceId),
		uint64(m.DepositNonce),
		m.ResourceId,
		ethCommon.BytesToAddress(recipient[1:]),
		amount,
		m.Signatures)

	if err != nil {
		w.log.Warn("TriggerContract commitVotes ", "src", m.SourceId, "dest", m.DestinationId, "depositNonce", m.DepositNonce, "err", err)
		return nil, err
	}

	signedTx, err := w.signTx(tx.Transaction)

	if err != nil {
		w.log.Warn("commitVotes sign tx err", "tx", hexutil.Encode(tx.GetTxid()), "src", m.SourceId, "dest", m.DestinationId, "depositNonce", m.DepositNonce, "err", err)
		return nil, err
	}

	ret, err := w.conn.Client().Broadcast(signedTx)
	if err == nil {

		if ret.Message == nil {
			w.log.Info("commitVotes", "tx", hexutil.Encode(tx.GetTxid()), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce,
				"recipient", tronCommon.EncodeCheck(recipient),
				"amount", amount.String())
			return tx, nil
		} else {
			w.log.Warn("Broadcast  commitVotes tx err", "tx", hexutil.Encode(tx.GetTxid()), "src", m.SourceId, "depositNonce", m.DepositNonce, "err", string(ret.Message))
			return nil, errors.New(string(ret.Message))
		}

	} else {
		w.log.Error("commitVotes", "tx", hexutil.Encode(tx.GetTxid()), "src", m.SourceId, "dst", m.DestinationId, "nonce", m.DepositNonce,
			"recipient", tronCommon.EncodeCheck(recipient),
			"amount", amount.String(), "err", err)
		return nil, err
	}

}
