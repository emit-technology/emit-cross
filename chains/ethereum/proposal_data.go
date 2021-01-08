// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	utils "github.com/emit-technology/emit-cross/shared/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func ConstructErc20ProposalDataHash(handler common.Address, amount []byte, recipient []byte) [32]byte {
	var data []byte

	data = append(data, handler[:]...)
	data = append(data, recipient[:]...)
	data = append(data, common.LeftPadBytes(amount, 32)...)
	return utils.Hash(data)

}
