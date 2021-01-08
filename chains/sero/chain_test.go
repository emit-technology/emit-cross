// Copyright 2020 EMIT Foundation.
// This file is part of E.M.I.T. .

// E.M.I.T. is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// E.M.I.T. is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with E.M.I.T. . If not, see <http://www.gnu.org/licenses/>.

package sero

import (
	"github.com/sero-cash/go-sero/common"
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	serotest "github.com/emit-technology/emit-cross/shared/sero/testing"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func TestChain_ListenerShutdownOnFailure(t *testing.T) {
	client := serotest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, msg.ChainId(1))
	cfg := &core.ChainConfig{
		Id:             msg.ChainId(1),
		Name:           "alice",
		Endpoint:       TestEndpoint,
		From:           keystore.AliceKey,
		Insecure:       true,
		KeystorePath:   keystore.AliceKey,
		BlockstorePath: "",
		FreshStart:     true,
		Opts: map[string]string{
			"bridge":       contracts.BridgeAddress.String(),
			"src20Handler": contracts.SRC20HandlerAddress.String(),
			"gasLimit":     big.NewInt(DefaultGasLimit).String(),
			"maxGasPrice":  big.NewInt(DefaultGasPrice).String(),
		},
	}
	sysErr := make(chan error)
	chain, _, err := InitializeChain(cfg, TestLogger, sysErr, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = chain.Start()
	if err != nil {
		t.Fatal(err)
	}

	// Cause critical failure to polling mechanism
	chain.conn.Client().Close()

	// Pull expected error
	select {
	case err := <-sysErr:
		if err.Error() != ErrFatalPolling.Error() {
			t.Fatalf("Unexpected error: %s", err)
		}
	case <-time.After(time.Second * 30):
		t.Fatal("Test timed out")
	}

	// Tell everyone to shutdown
	chain.Stop()
}

func TestChain_WriterShutdownOnFailure(t *testing.T) {
	// Setup contracts and params for erc20 transfer
	client := serotest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, msg.ChainId(1))
	cy := "TESTA"
	src20Contract := serotest.DeployMintableSrc20(t, client, contracts.SRC20HandlerAddress, cy)
	src := msg.ChainId(5) // Not yet used, nonce should be 0
	dst := msg.ChainId(1)
	amount := big.NewInt(10)
	resourceId := msg.ResourceIdFromSlice(append(common.LeftPadBytes(src20Contract.Bytes(), 31), uint8(src)))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)
	serotest.RegisterResource(t, client, contracts.BridgeAddress, contracts.SRC20HandlerAddress, resourceId, cy, big.NewInt(0))
	serotest.SetResoruceTokenAddress(t, client, contracts.BridgeAddress, contracts.SRC20HandlerAddress, resourceId, src20Contract)

	// Start a chain
	cfg := &core.ChainConfig{
		Id:             dst,
		Name:           "alice",
		Endpoint:       TestEndpoint,
		From:           keystore.AliceKey,
		Insecure:       true,
		KeystorePath:   keystore.AliceKey,
		BlockstorePath: "",
		FreshStart:     true,
		Opts: map[string]string{
			"bridge":       contracts.BridgeAddress.String(),
			"src20Handler": contracts.SRC20HandlerAddress.String(),
			"gasLimit":     big.NewInt(DefaultGasLimit).String(),
			"maxGasPrice":  big.NewInt(DefaultGasPrice).String(),
		},
	}
	sysErr := make(chan error)
	chain, _, err := InitializeChain(cfg, TestLogger, sysErr, nil)
	if err != nil {
		t.Fatal(err)
	}

	r := core.NewRouter(TestLogger)
	chain.SetRouter(r)

	err = chain.Start()
	if err != nil {
		t.Fatal(err)
	}

	// Submit some messages
	message := msg.NewFungibleTransfer(src, dst, 1, amount, resourceId, recipient.Bytes())

	for i := 0; i < 5; i++ {
		err = chain.listener.router.Send(message)
		if err != nil {
			t.Fatal(err)
		}

		message.DepositNonce++
	}

	time.Sleep(time.Second)
	// Cause critical failure for submitting txs
	chain.conn.Client().Close()

	// Pull expected error
	select {
	case err := <-sysErr:
		if err.Error() != ErrFatalPolling.Error() &&
			err.Error() != ErrFatalTx.Error() &&
			err.Error() != ErrFatalQuery.Error() {
			t.Fatalf("Unexpected error: %s", err)
		}
	case <-time.After(time.Second * 30):
		t.Fatal("Test timed out")
	}

	// Tell everyone to shutdown
	chain.Stop()
}
