// Copyright 2020 The emit-bridge Authors
// This file is part of the emit-bridge library.
//
// The emit-bridge library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The emit-bridge library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the emit-bridge library. If not, see <http://www.gnu.org/licenses/>.

package sero

import (
	"context"
	"errors"
	"github.com/sero-cash/go-sero"
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/common/hexutil"
	"github.com/sero-cash/go-sero/core/types"

	"github.com/sero-cash/go-sero/zero/txtool"
	"math/big"

	"github.com/sero-cash/go-sero/rpc"
	"github.com/sero-cash/go-sero/seroclient"
)


type WrappedClient struct {
	SeroClint *seroclient.Client
	SeroRPCClient *rpc.Client
	AccountRPCClient *rpc.Client
}

func Dial(geroUrl,acountUrl string) (*WrappedClient, error) {
	return DialContext(context.Background(), geroUrl,acountUrl)
}

func DialContext(ctx context.Context, geroUrl,acountUrl string) (*WrappedClient, error) {
	seroRPCClient, err := rpc.DialContext(ctx, acountUrl)
	if err !=nil {
		return nil,err
	}
	accountRPCClient, err := rpc.DialContext(ctx, acountUrl)
	if err != nil {
		return nil, err
	}
	return NewClient(seroRPCClient,accountRPCClient), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(seroRPCClint *rpc.Client,accountRPCClient *rpc.Client) *WrappedClient {
	return &WrappedClient{seroclient.NewClient(seroRPCClint),seroRPCClint,accountRPCClient}
}


func (ec *WrappedClient) Close() {
	ec.SeroClint.Close()
	ec.AccountRPCClient.Close()
}


func (ec *WrappedClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return ec.SeroClint.CodeAt(ctx,account,blockNumber)

}

func (ec *WrappedClient) CallContract(ctx context.Context, msg sero.CallMsg, blockNumber *big.Int) ([]byte, error) {

	return ec.SeroClint.CallContract(ctx,msg,blockNumber)
}


// PendingCodeAt returns the contract code of the given account in the pending state.
func (ec *WrappedClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return ec.SeroClint.PendingCodeAt(ctx ,account)
}



// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (ec *WrappedClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return ec.SeroClint.SuggestGasPrice(ctx)
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (ec *WrappedClient) EstimateGas(ctx context.Context, msg sero.CallMsg) (uint64, error) {
	return ec.SeroClint.EstimateGas(ctx,msg)

}

func (ec *WrappedClient) GenContractTx(ctx context.Context, msg sero.CallMsg) (*txtool.GTxParam, error) {
	var param txtool.GTxParam
	if msg.FromPKr == nil {
		return nil, errors.New("from is nil")
	}
	if err := ec.AccountRPCClient.CallContext(ctx, &param, "sero_genTx", toContractTxArgs(msg)); err != nil {
		return nil, err
	}
	return &param, nil
}

func (ec *WrappedClient) CyToContractAddress(ctx context.Context, cy string) (common.Address,error){
	var result string
	err := ec.SeroRPCClient.CallContext(ctx, &result, "sero_currencyToContractAddress", cy)
    if err!=nil {
    	return common.Address{},err
	}
	return common.Base58ToAddress(result), err


}

func (ec *WrappedClient) CommitTx(ctx context.Context, arg *txtool.GTx) error {

	return ec.SeroClint.CommitTx(ctx,arg);

}

func (ec *WrappedClient) GetFullAdress(ctx context.Context,contractAddr common.ContractAddress) (common.Address,error){
	var param =[]common.ContractAddress{}
	param = append(param,contractAddr)
	var result = make(map[common.ContractAddress]common.Address)
	err:=ec.SeroRPCClient.CallContext(ctx,&result,"sero_getFullAddress",param)
	if err !=nil {
		return common.Address{},err
	}
	return result[contractAddr],nil

}


func toContractTxArgs(msg sero.CallMsg) interface{} {
	arg := map[string]interface{}{
		"to": msg.To,
	}
	arg["From"] = msg.From
	var fromPkr common.Address
	copy(fromPkr[:], msg.FromPKr[:])
	arg["RefundTo"] = fromPkr

	if msg.Gas != 0 {
		arg["Gas"] = msg.Gas
	}
	if msg.GasPrice != nil {
		arg["GasPrice"] = msg.GasPrice
	}

	contractArgs := map[string]interface{}{}

	if msg.Currency != "" {
		contractArgs["Currency"] = msg.Currency
	} else {
		contractArgs["Currency"] = "SERO"
	}
	contractArgs["Value"] = msg.Value
	if msg.To != nil {
		contractArgs["To"] = hexutil.Bytes(msg.To[:])
	}
	if len(msg.Data) > 0 {
		contractArgs["Data"] = hexutil.Bytes(msg.Data)
	}
	cmdArgs := map[string]interface{}{}
	cmdArgs["Contract"] = contractArgs

	arg["Cmds"] = cmdArgs
	return arg
}


func (ec *WrappedClient) FilterLogs(ctx context.Context, q sero.FilterQuery) ([]types.Log, error) {
	return ec.SeroClint.FilterLogs(ctx,q)
}

func (ec *WrappedClient) SubscribeFilterLogs(ctx context.Context, q sero.FilterQuery, ch chan<- types.Log) (sero.Subscription, error) {
	return ec.SeroClint.SubscribeFilterLogs(ctx,q,ch)
}

func (ec *WrappedClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return ec.SeroClint.TransactionReceipt(ctx,txHash)

}

func (ec *WrappedClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {

	return ec.SeroClint.TransactionByHash(ctx,hash)
}



