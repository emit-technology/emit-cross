// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"github.com/emit-technology/emit-cross/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

func ConstructErc20ProposalDataHash(handler ethCommon.Address, recipient ethCommon.Address, amount *big.Int) [32]byte {
	var data []byte

	data = append(data, handler[:]...)
	data = append(data, recipient[:]...)
	data = append(data, ethCommon.LeftPadBytes(amount.Bytes(), 32)...)
	return common.Hash(data)

}
