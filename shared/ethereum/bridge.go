// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/emit-technology/emit-cross/bindings/ethereum/BridgeCounter"
	"github.com/emit-technology/emit-cross/types"
	"math/big"

	"github.com/emit-technology/emit-cross/bindings/ethereum/Bridge"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterResource(client *Client, bridge, handler common.Address, rId types.ResourceId, addr common.Address) error {
	instance, err := Bridge.NewBridge(bridge, client.Client)
	if err != nil {
		return err
	}

	err = client.LockNonceAndUpdate()
	if err != nil {
		return err
	}

	tx, err := instance.AdminSetResource(client.Opts, handler, rId, addr, new(big.Int), big.NewInt(100000000000))
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	client.UnlockNonce()

	return nil
}

func SetBurnable(client *Client, bridge, contract common.Address, resourceId [32]byte) error {
	instance, err := Bridge.NewBridge(bridge, client.Client)
	if err != nil {
		return err
	}

	err = client.LockNonceAndUpdate()
	if err != nil {
		return err
	}

	tx, err := instance.AdminSetBurnable(client.Opts, resourceId, contract)
	if err != nil {
		return err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return err
	}

	client.UnlockNonce()

	return nil
}

func GetDepositNonce(client *Client, bridgeCounter common.Address, chain types.ChainId) (uint64, error) {
	instance, err := BridgeCounter.NewBridgeCounter(bridgeCounter, client.Client)
	if err != nil {
		return 0, err
	}

	count, err := instance.DepositCounts(client.CallOpts, uint8(chain))
	if err != nil {
		return 0, err
	}

	return count, nil
}
