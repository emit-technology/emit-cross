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
	"math/big"
	"reflect"
	"testing"

	"github.com/ChainSafe/chainbridge-utils/core"
)

//TestParseChainConfig tests parseChainConfig with all handlerContracts provided
func TestParseChainConfig(t *testing.T) {

	input := core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"bridge":             "0x1234",
			"erc20Handler":       "0x1234",
			"erc721Handler":      "0x1234",
			"genericHandler":     "0x1234",
			"gasLimit":           "10",
			"maxGasPrice":        "20",
			"http":               "true",
			"startBlock":         "10",
			"blockConfirmations": "50",
		},
	}

	out, err := parseChainConfig(&input)

	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		name:         "chain",
		id:           1,
		endpoint:     "endpoint",
		from:         "0x0",
		keystorePath: "./keys",
		//bridgeContract:         common.HexToAddress("0x1234"),
		//src20HandlerContract:   common.HexToAddress("0x1234"),
		//src721HandlerContract:  common.HexToAddress("0x1234"),
		//genericHandlerContract: common.HexToAddress("0x1234"),
		gasLimit:           big.NewInt(10),
		maxGasPrice:        big.NewInt(20),
		http:               true,
		startBlock:         big.NewInt(10),
		blockConfirmations: big.NewInt(50),
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

//TestChainConfigOneContract Tests chain config providing only one contract
func TestChainConfigOneContract(t *testing.T) {

	input := core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"bridge":       "0x1234",
			"erc20Handler": "0x1234",
			"gasLimit":     "10",
			"maxGasPrice":  "20",
			"http":         "true",
			"startBlock":   "10",
		},
	}

	out, err := parseChainConfig(&input)

	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		name:         "chain",
		id:           1,
		endpoint:     "endpoint",
		from:         "0x0",
		keystorePath: "./keys",
		//bridgeContract:       common.HexToAddress("0x1234"),
		//erc20HandlerContract: common.HexToAddress("0x1234"),
		gasLimit:           big.NewInt(10),
		maxGasPrice:        big.NewInt(20),
		http:               true,
		startBlock:         big.NewInt(10),
		blockConfirmations: big.NewInt(0),
	}

	if !reflect.DeepEqual(&expected, out) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, out)
	}
}

func TestRequiredOpts(t *testing.T) {
	// No opts provided
	input := core.ChainConfig{
		Id:           0,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts:         map[string]string{},
	}

	_, err := parseChainConfig(&input)

	if err == nil {
		t.Error("config missing chainId field but no error reported")
	}

	// Empty bridgeContract provided
	input = core.ChainConfig{
		Id:           0,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts:         map[string]string{"bridge": ""},
	}

	_, err2 := parseChainConfig(&input)

	if err2 == nil {
		t.Error("config missing chainId field but no error reported")
	}

}

func TestExtraOpts(t *testing.T) {
	input := core.ChainConfig{
		Name:         "chain",
		Id:           1,
		Endpoint:     "endpoint",
		From:         "0x0",
		KeystorePath: "./keys",
		Insecure:     false,
		Opts: map[string]string{
			"bridge":        "0x1234",
			"gasLimit":      "10",
			"maxGasPrice":   "20",
			"http":          "true",
			"incorrect_opt": "error",
		},
	}

	_, err := parseChainConfig(&input)

	if err == nil {
		t.Error("Config should not accept incorrect opts.")
	}
}
