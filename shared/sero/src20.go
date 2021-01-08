// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"context"
	SRC20 "github.com/emit-technology/emit-cross/bindings/sero/SRC20Mintable"
	"github.com/sero-cash/go-sero/common"
)

// DeployMintAndApprove deploys a new erc20 contract, mints to the deployer, and approves the erc20 handler to transfer those token.
func DeployMintableSrc20(client *Client, src20Handler common.Address,cy string) (common.Address, error) {

	// Deploy
	_, tx, erc20Instance, err :=SRC20.DeploySRC20Mintable(client.Opts, client.Client, cy, cy,18)
	if err != nil {
		return common.Address{}, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return common.Address{}, err
	}
	addr,err:=client.Client.CyToContractAddress(context.Background(),cy);
	if err !=nil {
		return common.Address{},err
	}

	_,err=erc20Instance.AddMinter(client.Opts,src20Handler)
	if err !=nil {
		return common.Address{},err
	}

	return addr,nil



}

