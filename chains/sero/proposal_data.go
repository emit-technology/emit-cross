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
	seroCommon "github.com/sero-cash/go-sero/common"
	"math/big"
)

func ConstructSrc20ProposalDataHash(handler seroCommon.Address, recipient seroCommon.Address, amount *big.Int) [32]byte {
	var data []byte
	handlerContractAddr := common.GenContractAddress(handler)
	recipientAddr := common.GenContractAddress(recipient)
	data = append(data, seroCommon.LeftPadBytes(handlerContractAddr[:], 20)...)
	data = append(data, seroCommon.LeftPadBytes(recipientAddr[:], 20)...)
	data = append(data, seroCommon.LeftPadBytes(amount.Bytes(), 32)...)

	return common.Hash(data)

}
