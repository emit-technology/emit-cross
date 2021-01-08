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
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ChainSafe/chainbridge-utils/blockstore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/emit-technology/emit-cross/bindings/sero/Bridge"
	"github.com/emit-technology/emit-cross/bindings/sero/SRC20Handler"

	utils "github.com/emit-technology/emit-cross/shared/sero"
	serotest "github.com/emit-technology/emit-cross/shared/sero/testing"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	seroCommon "github.com/sero-cash/go-sero/common"
)

type MockRouter struct {
	msgs chan msg.Message
}

func (r *MockRouter) Send(message msg.Message) error {
	r.msgs <- message
	return nil
}

func createTestListener(t *testing.T, config *Config, contracts *utils.DeployedContracts, stop <-chan int, sysErr chan<- error) (*listener, *MockRouter) {
	// Create copy and add deployed contract addresses
	newConfig := *config
	newConfig.bridgeContract = contracts.BridgeAddress
	newConfig.src20HandlerContract = contracts.SRC20HandlerAddress

	conn := newLocalConnection(t, &newConfig)
	latestBlock, err := conn.LatestBlock()
	if err != nil {
		t.Fatal(err)
	}
	newConfig.startBlock = latestBlock

	bridgeContract, err := Bridge.NewBridge(newConfig.bridgeContract, conn.Client())
	if err != nil {
		t.Fatal(err)
	}
	erc20HandlerContract, err := SRC20Handler.NewSRC20Handler(newConfig.src20HandlerContract, conn.Client())
	if err != nil {
		t.Fatal(err)
	}
	//signatureColletorContact, err := Collector.NewSignatureCollector(newConfig.signatureColletorContact, conn.Client())

	router := &MockRouter{msgs: make(chan msg.Message)}
	listener := NewListener(conn, &newConfig, TestLogger, &blockstore.EmptyStore{}, stop, sysErr, nil)
	listener.setContracts(bridgeContract, erc20HandlerContract)
	listener.setRouter(router)
	// Start the listener
	err = listener.start()
	if err != nil {
		t.Fatal(err)
	}

	return listener, router
}

func verifyMessage(t *testing.T, r *MockRouter, expected msg.Message, errs chan error) {
	// Verify message
	select {
	case m := <-r.msgs:
		err := compareMessage(expected, m)
		if err != nil {
			t.Fatal(err)
		}
	case err := <-errs:
		t.Fatalf("Fatal error: %s", err)
	case <-time.After(TestTimeout):
		t.Fatalf("test timed out")
	}
}

func TestListener_start_stop(t *testing.T) {
	client := serotest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id)
	stop := make(chan int)
	l, _ := createTestListener(t, aliceTestConfig, contracts, stop, nil)

	err := l.start()
	if err != nil {
		t.Fatal(err)
	}

	// Initiate shutdown
	close(stop)
}

func TestListener_Erc20DepositedEvent(t *testing.T) {
	client := serotest.NewClient(t, TestEndpoint, AliceKp)
	contracts := deployTestContracts(t, client, aliceTestConfig.id)
	errs := make(chan error)
	l, router := createTestListener(t, aliceTestConfig, contracts, make(chan int), errs)

	amount := big.NewInt(10)
	src := msg.ChainId(0)
	dst := msg.ChainId(1)
	resourceId := msg.ResourceIdFromSlice(append(seroCommon.LeftPadBytes(contracts.SRC20HandlerAddress.Bytes(), 31), uint8(src)))
	recipient := ethcrypto.PubkeyToAddress(BobKp.PrivateKey().PublicKey)

	serotest.RegisterResource(t, client, contracts.BridgeAddress, contracts.SRC20HandlerAddress, resourceId, "TESTA", big.NewInt(0))

	expectedMessage := msg.NewFungibleTransfer(
		src,
		dst,
		1,
		amount,
		resourceId,
		[]byte{},
	)
	// Create an ERC20 Deposit
	createErc20Deposit(
		t,
		l.bridgeContract,
		client,
		resourceId,

		recipient[:],
		dst,
		amount,
	)

	verifyMessage(t, router, expectedMessage, errs)

	// Create second deposit, verify nonce change
	expectedMessage = msg.NewFungibleTransfer(
		src,
		dst,
		2,
		amount,
		resourceId,
		common.HexToAddress(BobKp.Address()).Bytes(),
	)
	createErc20Deposit(
		t,
		l.bridgeContract,
		client,
		resourceId,

		recipient[:],
		dst,
		amount,
	)

	verifyMessage(t, router, expectedMessage, errs)
}

func compareMessage(expected, actual msg.Message) error {
	if !reflect.DeepEqual(expected, actual) {
		if !reflect.DeepEqual(expected.Source, actual.Source) {
			return fmt.Errorf("Source doesn't match. \n\tExpected: %#v\n\tGot: %#v\n", expected.Source, actual.Source)
		} else if !reflect.DeepEqual(expected.Destination, actual.Destination) {
			return fmt.Errorf("Destination doesn't match. \n\tExpected: %#v\n\tGot: %#v\n", expected.Destination, actual.Destination)
		} else if !reflect.DeepEqual(expected.DepositNonce, actual.DepositNonce) {
			return fmt.Errorf("Deposit nonce doesn't match. \n\tExpected: %#v\n\tGot: %#v\n", expected.DepositNonce, actual.DepositNonce)
		} else if !reflect.DeepEqual(expected.Payload, actual.Payload) {
			return fmt.Errorf("Payload doesn't match. \n\tExpected: %#v\n\tGot: %#v\n", expected.Payload, actual.Payload)
		}
	}
	return nil
}
