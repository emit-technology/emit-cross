// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/emit-technology/emit-cross/bindings/sero/BridgeCounter"
	"math/big"

	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/emit-technology/emit-cross/bindings/sero/Bridge"
	seroCommon "github.com/sero-cash/go-sero/common"
)

var Max, _ = new(big.Int).SetString("100000000000000000000000000000", 10)

func RegisterResource(client *Client, bridge, handler seroCommon.Address, rId msg.ResourceId, cy string, min *big.Int) error {
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

func SetResoruceTokenAddress(client *Client, bridge, handler seroCommon.Address, rId msg.ResourceId, addr seroCommon.Address) error {
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

func GetDepositNonce(client *Client, bridge seroCommon.Address, chain msg.ChainId) (uint64, error) {
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

func IDAndNonce(srcId msg.ChainId, nonce msg.Nonce) *big.Int {
	var data []byte
	data = append(data, nonce.Big().Bytes()...)
	data = append(data, uint8(srcId))
	return big.NewInt(0).SetBytes(data)
}
