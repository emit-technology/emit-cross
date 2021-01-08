// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sero-cash/go-sero/common"
)

type EventSig string

func (es EventSig) GetTopic() (ret common.Hash) {

	h := crypto.Keccak256Hash([]byte(es))
	copy(ret[:], h[:])
	return
}

const (
	Deposit       EventSig = "Deposit(uint8,bytes32,uint64)"
	ProposalEvent EventSig = "ProposalEvent(uint8,uint64,uint8,bytes32,bytes32)"
	SignProposal  EventSig = "SignProposal(uint8,uint8,uint64,bytes)"
)

type ProposalStatus int

const (
	Inactive ProposalStatus = iota
	Active
	Passed
	Executed
	Cancelled
)

func IsActive(status uint8) bool {
	return ProposalStatus(status) == Active
}

func IsFinalized(status uint8) bool {
	return ProposalStatus(status) == Passed
}

func IsExecuted(status uint8) bool {
	return ProposalStatus(status) == Executed
}
