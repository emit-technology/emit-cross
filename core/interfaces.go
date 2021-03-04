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

package core

import (
	"math/big"
)

type ProposalState interface {
	GetDepositFTRecord(nonce uint64, dest uint8) (resourceID [32]byte, destinationRecipientAddress []byte, amount *big.Int, err error)
	GetDepositNFTRecord(nonce uint64, dest uint8) (src721ResourceID [32]byte,
		destinationRecipientAddress []byte,
		tokenId *big.Int,
		metadata []byte,
		src20Amount *big.Int,
		err error)
	FTPropsalDataHash(recipient []byte, amount *big.Int) [32]byte
	NFTPropsalDataHash(recipient []byte, tokenId *big.Int, metadata []byte, feeAmount *big.Int) [32]byte
	GetFTProposalStatus(source uint8, nonce uint64, dataHash [32]byte) (uint8, error)
	GetNFTProposalStatus(source uint8, nonce uint64, dataHash [32]byte) (uint8, error)

	GetBridgeAddress() []byte
	GetNFTBridgeAddress() []byte
	IsWithCollector() bool
}
