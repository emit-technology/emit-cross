// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"testing"

	"github.com/ChainSafe/chainbridge-utils/msg"
	utils "github.com/emit-technology/emit-cross/shared/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterResource(t *testing.T, client *utils.Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address) {
	err := utils.RegisterResource(client, bridge, handler, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func SetBurnable(t *testing.T, client *utils.Client, bridge, contract common.Address, resourceId [32]byte) {
	err := utils.SetBurnable(client, bridge, contract, resourceId)
	if err != nil {
		t.Fatal(err)
	}
}

func GetDepositNonce(t *testing.T, client *utils.Client, bridge common.Address, chain msg.ChainId) uint64 {
	count, err := utils.GetDepositNonce(client, bridge, chain)
	if err != nil {
		t.Fatal(err)
	}
	return count
}
