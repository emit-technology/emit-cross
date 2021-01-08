// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package serotest

import (
	"testing"

	utils "github.com/emit-technology/emit-cross/shared/sero"
	"github.com/sero-cash/go-sero/common"
)

func DeployMintableSrc20(t *testing.T, client *utils.Client, src20Handler common.Address, cy string) common.Address {
	addr, err := utils.DeployMintableSrc20(client, src20Handler,cy)
	if err != nil {
		t.Fatal(err)
	}
	return addr
}
