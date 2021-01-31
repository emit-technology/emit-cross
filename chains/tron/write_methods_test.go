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

package tron

import (
	"encoding/json"
	"fmt"
	"github.com/emit-technology/emit-cross/bindings/tron/Bridge"
	"github.com/fbsobreira/gotron-sdk/pkg/abi"
	tron "github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/sero-cash/go-sero/common/hexutil"
	"strconv"
	"testing"
	"time"
)

func TestExcuet(t *testing.T) {
	params := []abi.Param{
		{"uint8": strconv.FormatUint(2, 10)},
		{"uint64": strconv.FormatUint(3, 10)},
		{"bytes32": "000000000000000000000040655D1b70d8a73eF350AEbFE6278a212B018Ade03"},
		{"address": "TYdRxvWxSBRmm76tKX73rULihnpxi5aGjG"},
		{"uint256": "10000000"},
	}

	paramsJson, err := json.Marshal(params)

	grpcClient := tron.NewGrpcClientWithTimeout("grpc.shasta.trongrid.io:50051", 100*time.Second)
	err = grpcClient.Start()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(paramsJson))
	result, err := grpcClient.TriggerConstantContract("", "TVUYNFDGFDrkJD2rSYXGiFLBC52WZFk7WP", "executeProposal(uint8,uint64,bytes32,address,uint256)", string(paramsJson))
	if err != nil {
		t.Fatal(err)
	}
	if result.Result.Code > 0 || result.Result.Message != nil {
		t.Fatal(string(result.Result.Message))
	}

}

func TestGetProposalStatus(t *testing.T) {
	params := []abi.Param{
		{"uint8": "2"},
		{"uint64": "17"},
		{"bytes32": "e308e6957c054a7049cf132d0f4ab31c06b5e0af5c51eb2cae273333b5beca8c"},
	}
	paramsJson, err := json.Marshal(params)

	if err != nil {
		t.Error(err)
	}

	grpcClient := tron.NewGrpcClientWithTimeout("grpc.shasta.trongrid.io:50051", 100*time.Second)
	err = grpcClient.Start()
	if err != nil {
		t.Fatal(err)
	}
	result, err := grpcClient.TriggerConstantContract("", "TNstBgWNyCMxnkbZzCe3eAyiYiuY45sPTi", "getProposal(uint8,uint64,bytes32)", string(paramsJson))
	if err != nil {
		t.Fatal(err)
	}
	if result.Result.Code > 0 || result.Result.Message != nil {
		t.Fatal(string(result.Result.Message))
	}
	proposal, err := Bridge.NewBridgeABI().GetProposal(result.GetConstantResult()[0])
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(proposal.Status)

}

func TestHextoString(t *testing.T) {
	s, _ := hexutil.Decode("0x20")
	fmt.Println(string(s))
}
