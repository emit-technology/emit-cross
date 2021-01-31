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

type FTTransfer struct {
	BlockNum      uint64
	SourceId      ChainId
	DestinationId ChainId
	DepositNonce  Nonce
	ResourceId    ResourceId
	Recipient     Recipient
	Amount        *big.Int
}

type BatchVotes struct {
	BlockNum      uint64
	SourceId      ChainId
	DestinationId ChainId
	DepositNonce  Nonce
	ResourceId    ResourceId
	Recipient     Recipient
	Amount        *big.Int
	Signatures    [][]byte
}
