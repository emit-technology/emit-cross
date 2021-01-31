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
	"github.com/emit-technology/emit-cross/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	seroCommon "github.com/sero-cash/go-sero/common"
	"math/big"
)

func ConstructTrc20ProposalDataHash(handler string, recipient []byte, amount *big.Int) [32]byte {
	var data []byte
	h, _ := address.Base58ToAddress(handler)
	data = append(data, seroCommon.LeftPadBytes(h[1:], 20)...)
	data = append(data, seroCommon.LeftPadBytes(recipient[1:], 20)...)
	data = append(data, seroCommon.LeftPadBytes(amount.Bytes(), 32)...)

	return common.Hash(data)

}
