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

package types

import (
	"fmt"
	"math/big"
)

type ChainId uint8
type TransferType string
type ResourceId [32]byte
type Recipient []byte

func (r ResourceId) Hex() string {
	return fmt.Sprintf("0x%x", r)
}

type Nonce uint64

func (n Nonce) Big() *big.Int {
	return big.NewInt(int64(n))
}

var FungibleTransfer TransferType = "FungibleTransfer"
var NonFungibleTransfer TransferType = "NonFungibleTransfer"

type TransferMsg struct {
	BlockNum      uint64
	SourceId      ChainId
	DestinationId ChainId
	Type          TransferType
	DepositNonce  Nonce
	ResourceId    ResourceId
	Payload       [][]byte
}

type ProposalSignatures struct {
	BlockNum      uint64
	Type          TransferType
	SourceId      ChainId
	DestinationId ChainId
	DepositNonce  Nonce
	ResourceId    ResourceId
	Payload       [][]byte
	Signatures    [][]byte
}

func NewFungibleTransferMsg(blockNum uint64, source, dest ChainId, nonce Nonce, amount *big.Int, resourceId ResourceId, recipient []byte) TransferMsg {
	return TransferMsg{
		BlockNum:      blockNum,
		SourceId:      source,
		DestinationId: dest,
		Type:          FungibleTransfer,
		DepositNonce:  nonce,
		ResourceId:    resourceId,
		Payload: [][]byte{
			amount.Bytes(),
			recipient,
		},
	}
}

func NewNonFungibleTransferMsg(
	blockNum uint64,
	source, dest ChainId,
	nonce Nonce,
	tokenId *big.Int,
	resourceId ResourceId,
	recipient,
	metadata []byte,
	src20Amount *big.Int) TransferMsg {
	return TransferMsg{
		BlockNum:      blockNum,
		SourceId:      source,
		DestinationId: dest,
		Type:          NonFungibleTransfer,
		DepositNonce:  nonce,
		ResourceId:    resourceId,
		Payload: [][]byte{
			tokenId.Bytes(),
			recipient,
			metadata,
			src20Amount.Bytes(),
		},
	}
}

func NewFungibleProposalSignatures(blockNum uint64, source, dest ChainId, nonce Nonce, amount *big.Int, resourceId ResourceId, recipient []byte, signatures [][]byte) ProposalSignatures {
	return ProposalSignatures{
		BlockNum:      blockNum,
		SourceId:      source,
		DestinationId: dest,
		Type:          FungibleTransfer,
		DepositNonce:  nonce,
		ResourceId:    resourceId,
		Payload: [][]byte{
			amount.Bytes(),
			recipient,
		},
		Signatures: signatures,
	}
}

func NewNonFungibleProposalSignatures(blockNum uint64,
	source, dest ChainId,
	nonce Nonce,
	tokenId *big.Int,
	resourceId ResourceId,
	recipient, metadata []byte,
	src20Amount *big.Int,
	signatures [][]byte) ProposalSignatures {
	return ProposalSignatures{
		BlockNum:      blockNum,
		SourceId:      source,
		DestinationId: dest,
		Type:          NonFungibleTransfer,
		DepositNonce:  nonce,
		ResourceId:    resourceId,
		Payload: [][]byte{
			tokenId.Bytes(),
			recipient,
			metadata,
			src20Amount.Bytes(),
		},
		Signatures: signatures,
	}
}
