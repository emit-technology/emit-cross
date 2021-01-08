// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/emit-technology/emit-cross/bindings/sero/Collector"
	"github.com/emit-technology/emit-cross/common"
	"math/big"

	seroCommon "github.com/sero-cash/go-sero/common"

	"github.com/ChainSafe/chainbridge-utils/keystore"
	bridge "github.com/emit-technology/emit-cross/bindings/sero/Bridge"
	src20Handler "github.com/emit-technology/emit-cross/bindings/sero/SRC20Handler"
)

var (
	RelayerAddresses = []seroCommon.Address{
		common.GenCommonAddress(keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]),
		common.GenCommonAddress(keystore.TestKeyRing.EthereumKeys[keystore.BobKey]),
		common.GenCommonAddress(keystore.TestKeyRing.EthereumKeys[keystore.CharlieKey]),
		common.GenCommonAddress(keystore.TestKeyRing.EthereumKeys[keystore.DaveKey]),
		common.GenCommonAddress(keystore.TestKeyRing.EthereumKeys[keystore.EveKey]),
	}

	ZeroAddress = seroCommon.Address{}
)

type DeployedContracts struct {
	BridgeAddress            seroCommon.Address
	SRC20HandlerAddress      seroCommon.Address
	SignatureColletorContact seroCommon.Address
}

// DeployContracts deploys Bridge, Relayer, ERC20Handler, ERC721Handler and CentrifugeAssetHandler and returns the addresses
func DeployContracts(client *Client, chainID uint8, initialRelayerThreshold *big.Int) (*DeployedContracts, error) {
	bridgeAddr, err := deployBridge(client, chainID, RelayerAddresses, initialRelayerThreshold)
	if err != nil {
		return nil, err
	}

	erc20HandlerAddr, err := deployERC20Handler(client, bridgeAddr)
	if err != nil {
		return nil, err
	}
	singnatureCollectorAddr, err := deploySignatureCollector(client, RelayerAddresses)

	deployedContracts := DeployedContracts{bridgeAddr, erc20HandlerAddr, singnatureCollectorAddr}

	return &deployedContracts, nil

}

func deployBridge(client *Client, chainID uint8, relayerAddrs []seroCommon.Address, initialRelayerThreshold *big.Int) (seroCommon.Address, error) {

	bridgeAddr, tx, _, err := bridge.DeployBridge(client.Opts, client.Client, chainID, big.NewInt(100), seroCommon.Address{})
	if err != nil {
		return ZeroAddress, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return bridgeAddr, nil

}

func deployERC20Handler(client *Client, bridgeAddress seroCommon.Address) (seroCommon.Address, error) {

	erc20HandlerAddr, tx, _, err := src20Handler.DeploySRC20Handler(client.Opts, client.Client, bridgeAddress)
	if err != nil {
		return ZeroAddress, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return erc20HandlerAddr, nil
}

func deploySignatureCollector(client *Client, relayerAddrs []seroCommon.Address) (seroCommon.Address, error) {

	signatureCollectorAddress, tx, _, err := Collector.DeploySignatureCollector(client.Opts, client.Client, relayerAddrs, big.NewInt(1))
	if err != nil {
		return ZeroAddress, err
	}

	err = WaitForTx(client, tx)
	if err != nil {
		return ZeroAddress, err
	}

	return signatureCollectorAddress, nil
}
