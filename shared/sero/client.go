// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"context"
	"fmt"
	"github.com/emit-technology/emit-cross/crypto/secp256k1"
	"github.com/sero-cash/go-sero/crypto"
	"math/big"
	"sync"
	"time"

	"github.com/sero-cash/go-czero-import/c_type"

	"github.com/emit-technology/emit-cross/bindings/sero"
	"github.com/emit-technology/emit-cross/common"
	"github.com/sero-cash/go-czero-import/superzk"

	"github.com/sero-cash/go-sero/zero/txtool"
	"github.com/sero-cash/go-sero/zero/txtool/flight"

	"github.com/sero-cash/go-sero/accounts/abi/bind"
	serotypes "github.com/sero-cash/go-sero/core/types"
	"github.com/sero-cash/go-sero/rpc"
)

const DefaultGasLimit = 6721975
const DefaultMaxGasPrice = 20000000000

var ExpectedBlockTime = time.Second

type Client struct {
	Client    *sero.WrappedClient
	Opts      *bind.TransactOpts
	CallOpts  *bind.CallOpts
	nonceLock sync.Mutex
}

func GenTransactor(kp *secp256k1.Keypair) *bind.TransactOpts {
	privKey := crypto.FromECDSA(kp.PrivateKey())
	tk := common.GenSeroTK(privKey)
	refundTo := common.GenMainPkr(kp)
	return &bind.TransactOpts{
		From:    tk.ToPk(),
		FromPKr: *refundTo,
		Value:   nil,
		Encrypter: func(txParam *txtool.GTxParam) (*txtool.GTx, error) {
			//priKey := crypto.FromECDSA(kp.PrivateKey())
			var seed c_type.Uint256
			copy(seed[:], privKey)
			sk := superzk.Seed2Sk(&seed, 1)
			gtx, err := flight.SignTx(&sk, txParam)
			if err != nil {
				return nil, err
			}
			return &gtx, nil
		},
	}
}

func NewClient(geroEndpoint string, accountEndpoint string, kp *secp256k1.Keypair) (*Client, error) {
	ctx := context.Background()
	seroRPCClient, err := rpc.DialWebsocket(ctx, geroEndpoint, "/ws")
	if err != nil {
		return nil, err
	}
	accountRPCClient, err := rpc.DialContext(ctx, accountEndpoint)
	if err != nil {
		return nil, err
	}
	client := sero.NewClient(seroRPCClient, accountRPCClient)

	opts := GenTransactor(kp)
	//opts.Nonce = big.NewInt(0)
	opts.Value = big.NewInt(0)              // in wei
	opts.GasLimit = uint64(DefaultGasLimit) // in units
	opts.GasPrice = big.NewInt(DefaultMaxGasPrice)
	opts.Context = ctx

	return &Client{
		Client: client,
		Opts:   opts,
		CallOpts: &bind.CallOpts{
			FromPKr: common.GenMainPkr(kp),
		},
	}, nil
}

/*
func (c *Client) LockNonceAndUpdate() error {
	c.nonceLock.Lock()
	nonce, err := c.Client.PendingNonceAt(context.Background(), c.Opts.From)
	if err != nil {
		c.nonceLock.Unlock()
		return err
	}
	c.Opts.Nonce.SetUint64(nonce)
	return nil
}

func (c *Client) UnlockNonce() {
	c.nonceLock.Unlock()
}*/

// WaitForTx will query the chain at ExpectedBlockTime intervals, until a receipt is returned.
// Returns an error if the tx failed.
func WaitForTx(client *Client, tx *serotypes.Transaction) error {
	retry := 10
	for retry > 0 {
		receipt, err := client.Client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			retry--
			time.Sleep(ExpectedBlockTime)
			continue
		}

		if receipt.Status != 1 {
			return fmt.Errorf("transaction failed on chain")
		}
		return nil
	}
	return nil
}
