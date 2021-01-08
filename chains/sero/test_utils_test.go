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
	"github.com/emit-technology/emit-cross/common"
	"github.com/sero-cash/go-sero/common/hexutil"
	"math/big"
	"testing"
	"time"

	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/emit-technology/emit-cross/bindings/sero/Bridge"
	connection "github.com/emit-technology/emit-cross/connections/sero"
	utils "github.com/emit-technology/emit-cross/shared/sero"
	seroCommon "github.com/sero-cash/go-sero/common"
)

const TestEndpoint = "ws://localhost:8545"

var TestLogger = newTestLogger("test")
var TestTimeout = time.Second * 30

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]

var TestRelayerThreshold = big.NewInt(2)
var TestChainId = msg.ChainId(0)

var aliceTestConfig = createConfig("alice", nil, nil)

func createConfig(name string, startBlock *big.Int, contracts *utils.DeployedContracts) *Config {
	cfg := &Config{
		name:                 name,
		id:                   0,
		endpoint:             TestEndpoint,
		from:                 name,
		keystorePath:         "",
		blockstorePath:       "",
		freshStart:           true,
		bridgeContract:       seroCommon.Address{},
		src20HandlerContract: seroCommon.Address{},
		gasLimit:             big.NewInt(DefaultGasLimit),
		maxGasPrice:          big.NewInt(DefaultGasPrice),
		http:                 false,
		startBlock:           startBlock,
		blockConfirmations:   big.NewInt(0),
	}

	if contracts != nil {
		cfg.bridgeContract = contracts.BridgeAddress
		cfg.src20HandlerContract = contracts.SRC20HandlerAddress
	}

	return cfg
}

func newTestLogger(name string) log15.Logger {
	tLog := log15.New("chain", name)
	tLog.SetHandler(log15.LvlFilterHandler(log15.LvlError, tLog.GetHandler()))
	return tLog
}

func newLocalConnection(t *testing.T, cfg *Config) *connection.Connection {
	kp := keystore.TestKeyRing.EthereumKeys[cfg.from]
	conn := connection.NewConnection(TestEndpoint, TestEndpoint, false, kp, TestLogger, big.NewInt(DefaultGasLimit), big.NewInt(DefaultGasPrice))
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func deployTestContracts(t *testing.T, client *utils.Client, id msg.ChainId) *utils.DeployedContracts {
	contracts, err := utils.DeployContracts(
		client,
		uint8(id),
		TestRelayerThreshold,
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("=======================================================")
	fmt.Printf("Bridge: %s\n", contracts.BridgeAddress.String())
	fmt.Printf("Src20Handler: %s\n", contracts.SRC20HandlerAddress.String())
	fmt.Println("========================================================")

	return contracts
}

func createErc20Deposit(
	t *testing.T,
	contract *Bridge.Bridge,
	client *utils.Client,
	rId msg.ResourceId,
	destRecipient []byte,
	destId msg.ChainId,
	amount *big.Int,
) {

	// Incrememnt Nonce by one
	if _, err := contract.DepositFT(
		client.Opts,
		uint8(destId),
		rId,
		destRecipient[:],
	); err != nil {
		t.Fatal(err)
	}
}

func TestConstructSrc20ProposalDataHash(t *testing.T) {
	var data []byte
	handlerContractAddr := common.GenContractAddress(seroCommon.Base58ToAddress("237YryP5zc86Dkzyj822T4EYP28vkKXj51qTr2n8u7iJrz9cHHp7qcFaxUCee5d4vrkRbcH48bnumYs5XADQQ2Qq"))
	recipientAddr := common.GenContractAddress(seroCommon.Base58ToAddress("2CQF13CUPTMsxyj4dnAftgQMWzsXzLdaHPYxz1USiAX2ApQGc35rjyQHQ4CYXw82q3yMqRUo9hD58L6ovHegQAFetf2wtXUuhQiAEgDeZHh2jWiB2LjtuayfypcHP1VNVKMV"))
	data = append(data, seroCommon.LeftPadBytes(handlerContractAddr[:], 20)...)
	data = append(data, seroCommon.LeftPadBytes(recipientAddr[:], 20)...)
	amount, _ := new(big.Int).SetString("100000000000000000000", 10)
	data = append(data, seroCommon.LeftPadBytes(amount.Bytes(), 32)...)
	fmt.Println(hexutil.Encode(data))
	hash := utils.Hash(data)
	fmt.Println(hexutil.Encode(hash[:]))

}
