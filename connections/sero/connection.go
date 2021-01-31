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
	"errors"
	"fmt"
	"github.com/emit-technology/emit-cross/crypto/secp256k1"
	log "github.com/emit-technology/emit-cross/log"
	"github.com/sero-cash/go-sero/crypto"
	"math/big"
	"sync"
	"time"

	"github.com/emit-technology/emit-cross/bindings/sero"
	"github.com/emit-technology/emit-cross/common"
	"github.com/sero-cash/go-czero-import/c_type"
	"github.com/sero-cash/go-czero-import/superzk"
	"github.com/sero-cash/go-sero/accounts/abi/bind"
	seroCommon "github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/rpc"
	"github.com/sero-cash/go-sero/zero/txtool"
	"github.com/sero-cash/go-sero/zero/txtool/flight"
)

var BlockRetryInterval = time.Second * 5

type Connection struct {
	endpoint        string
	accountEndpoint string
	http            bool
	kp              *secp256k1.Keypair
	gasLimit        *big.Int
	maxGasPrice     *big.Int
	conn            *sero.WrappedClient
	// signer    ethtypes.Signer
	opts     *bind.TransactOpts
	callOpts *bind.CallOpts
	//nonce    uint64
	optsLock sync.Mutex
	log      log.Logger
	stop     chan int // All routines should exit when this channel is closed
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(endpoint string, accountEndpoint string, http bool, kp *secp256k1.Keypair, log log.Logger, gasLimit, gasPrice *big.Int) *Connection {
	return &Connection{
		endpoint:        endpoint,
		accountEndpoint: accountEndpoint,
		http:            http,
		kp:              kp,
		gasLimit:        gasLimit,
		maxGasPrice:     gasPrice,
		log:             log,
		stop:            make(chan int),
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to sero chain...", "url", c.endpoint)
	var seroRpcClient *rpc.Client
	var err error
	// Start http or ws client
	if c.http {
		seroRpcClient, err = rpc.DialHTTP(c.endpoint)
	} else {
		seroRpcClient, err = rpc.DialWebsocket(context.Background(), c.endpoint, "/ws")
	}
	if err != nil {
		return err
	}
	accountRpcClient, err := rpc.DialHTTP(c.accountEndpoint)
	if err != nil {
		return nil
	}
	c.conn = sero.NewClient(seroRpcClient, accountRpcClient)

	c.opts = c.newTransactOpts(big.NewInt(0), c.gasLimit, c.maxGasPrice)

	c.callOpts = &bind.CallOpts{FromPKr: common.GenMainPkr(c.kp)}
	return nil
}

//
//// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Connection) newTransactOpts(value, gasLimit, gasPrice *big.Int) *bind.TransactOpts {
	privKey := crypto.FromECDSA(c.kp.PrivateKey())

	tk := common.GenSeroTK(privKey)
	refundTo := common.GenMainPkr(c.kp)
	return &bind.TransactOpts{
		From:     tk.ToPk(),
		FromPKr:  *refundTo,
		Value:    value,
		GasPrice: gasPrice,
		GasLimit: uint64(gasLimit.Int64()),
		Encrypter: func(txParam *txtool.GTxParam) (*txtool.GTx, error) {
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

func (c *Connection) Keypair() *secp256k1.Keypair {
	return c.kp
}

func (c *Connection) Client() *sero.WrappedClient {
	return c.conn
}

func (c *Connection) Opts() *bind.TransactOpts {
	return c.opts
}

func (c *Connection) CallOpts() *bind.CallOpts {
	return c.callOpts
}

func (c *Connection) SafeEstimateGas(ctx context.Context) (*big.Int, error) {
	gasPrice, err := c.conn.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, err
	}

	// Check we aren't exceeding our limit
	if gasPrice.Cmp(c.maxGasPrice) == 1 {
		return c.maxGasPrice, nil
	} else {
		return gasPrice, nil
	}
}

func (c *Connection) LockAndUpdateOpts() error {
	c.optsLock.Lock()

	gasPrice, err := c.SafeEstimateGas(context.TODO())
	if err != nil {
		c.optsLock.Unlock()
		return err
	}
	c.opts.GasPrice = gasPrice

	return nil
}

func (c *Connection) UnlockOpts() {
	c.optsLock.Unlock()
}

// LatestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.conn.SeroClint.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// EnsureHasBytecode asserts if contract code exists at the specified address
func (c *Connection) EnsureHasBytecode(addr seroCommon.Address) error {
	code, err := c.conn.SeroClint.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.String())
	}
	return nil
}

// WaitForBlock will poll for the block number until the current block is equal or greater than
func (c *Connection) WaitForBlock(block *big.Int) error {
	for {
		select {
		case <-c.stop:
			return errors.New("connection terminated")
		default:
			currBlock, err := c.LatestBlock()
			if err != nil {
				return err
			}

			// Equal or greater than target
			if currBlock.Cmp(block) >= 0 {
				return nil
			}
			c.log.Debug("Block not ready, waiting", "target", block, "current", currBlock)
			time.Sleep(BlockRetryInterval)
			continue
		}
	}
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
	close(c.stop)
}
