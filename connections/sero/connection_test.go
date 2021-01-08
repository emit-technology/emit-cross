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
	"context"
	"math/big"
	"testing"

	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/log15"
	seroUtils "github.com/emit-technology/emit-cross/shared/sero"
)

var TestEndpoint = "ws://localhost:8546"
var TestAccountEndpoint = "http://localhost:8545"

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var GasLimit = big.NewInt(seroUtils.DefaultGasLimit)
var MaxGasPrice = big.NewInt(seroUtils.DefaultMaxGasPrice)

func TestConnect(t *testing.T) {
	conn := NewConnection(TestEndpoint, TestAccountEndpoint, false, AliceKp, log15.Root(), GasLimit, MaxGasPrice)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
}

func TestConnection_SafeEstimateGas(t *testing.T) {
	// MaxGasPrice is the constant price on the dev network, so we increase it here by 1 to ensure it adjusts
	conn := NewConnection(TestEndpoint, TestAccountEndpoint, false, AliceKp, log15.Root(), GasLimit, MaxGasPrice.Add(MaxGasPrice, big.NewInt(1)))
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	price, err := conn.SafeEstimateGas(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if price.Cmp(MaxGasPrice) == 0 {
		t.Fatalf("Gas price should be less than max. Suggested: %s Max: %s", price.String(), MaxGasPrice.String())
	}
}

func TestConnection_SafeEstimateGasMax(t *testing.T) {
	maxPrice := big.NewInt(1)
	conn := NewConnection(TestEndpoint, TestAccountEndpoint, false, AliceKp, log15.Root(), GasLimit, maxPrice)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	price, err := conn.SafeEstimateGas(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if price.Cmp(maxPrice) != 0 {
		t.Fatalf("Gas price should equal max. Suggested: %s Max: %s", price.String(), maxPrice.String())
	}
}
