// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package serotest

import (
	"context"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	utils "github.com/emit-technology/emit-cross/shared/sero"
	"math/big"
	"testing"
)


func NewClient(t *testing.T, endpoint string, kp *secp256k1.Keypair) *utils.Client {
	client, err := utils.NewClient(endpoint,endpoint, kp)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func GetLatestBlock(t *testing.T, client *utils.Client) *big.Int {
	block, err := client.Client.SeroClint.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	return block.Number()
}

