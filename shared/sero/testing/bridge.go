// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package serotest

import (
	"math/big"
	"testing"

	utils "github.com/emit-technology/emit-cross/shared/sero"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/sero-cash/go-sero/common"
)

func RegisterResource(t *testing.T, client *utils.Client, bridge, handler common.Address, rId msg.ResourceId, cy string,min *big.Int) {
	err := utils.RegisterResource(client, bridge, handler, rId, cy,min)
	if err != nil {
		t.Fatal(err)
	}
}

func SetResoruceTokenAddress(t *testing.T, client *utils.Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address){
	err := utils.SetResoruceTokenAddress(client, bridge, handler, rId,addr)
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
