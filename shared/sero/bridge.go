// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/emit-technology/emit-cross/bindings/sero/BridgeCounter"
	"github.com/emit-technology/emit-cross/types"
	"math/big"

	"github.com/emit-technology/emit-cross/bindings/sero/Bridge"
	seroCommon "github.com/sero-cash/go-sero/common"
)

var Max, _ = new(big.Int).SetString("100000000000000000000000000000", 10)

func RegisterResource(client *Client, bridge, handler seroCommon.Address, rId types.ResourceId, cy string, min *big.Int) error {
	instance, err := Bridge.NewBridge(bridge, client.Client)
	if err != nil {
		return err
	}

	tx, err := instance.AdminSetResourceCurrency(client.Opts, rId, cy, min, Max)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	return nil
}

func SetResoruceTokenAddress(client *Client, bridge, handler seroCommon.Address, rId types.ResourceId, addr seroCommon.Address) error {
	instance, err := Bridge.NewBridge(bridge, client.Client)
	if err != nil {
		return err
	}

	tx, err := instance.AdminSetResourceTokenAddress(client.Opts, rId, addr)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	return nil
}

func GetDepositNonce(client *Client, bridge seroCommon.Address, chain types.ChainId) (uint64, error) {
	instance, err := BridgeCounter.NewBridgeCounter(bridge, client.Client)
	if err != nil {
		return 0, err
	}

	count, err := instance.DepositCounts(client.CallOpts, uint8(chain))
	if err != nil {
		return 0, err
	}

	return count, nil
}
