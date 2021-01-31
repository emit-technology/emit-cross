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
	"fmt"
	"github.com/sero-cash/go-sero/common/hexutil"
	"math/big"
	"testing"
)

func TestConstructTrc20ProposalDataHash(t *testing.T) {
	hander := "TFXrMv5Aei1HK4jJSa1ncNJVhRy5CpENqz"
	recipient, _ := hexutil.Decode("0x41F88D6518D0A8FAEA34899B0F2831107D3931F644")
	r := ConstructTrc20ProposalDataHash(hander, recipient, new(big.Int).SetUint64(1000))
	fmt.Println(hexutil.Encode(r[:]))
	s, _ := hexutil.Decode("0x524556455254206f70636f6465206578656375746564")
	fmt.Println(string(s))
}
