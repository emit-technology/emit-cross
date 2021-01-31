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
	"errors"
	"fmt"
	"github.com/emit-technology/emit-cross/crypto/secp256k1"
	log "github.com/emit-technology/emit-cross/log"
	"math/big"
	"sync"
	"time"

	tron "github.com/fbsobreira/gotron-sdk/pkg/client"
)

var BlockRetryInterval = time.Second * 5

type Connection struct {
	endpoint    string
	kp          *secp256k1.Keypair
	gasLimit    *big.Int
	maxGasPrice *big.Int
	conn        *tron.GrpcClient
	//nonce    uint64
	optsLock sync.Mutex
	log      log.Logger
	stop     chan int // All routines should exit when this channel is closed
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(endpoint string, kp *secp256k1.Keypair, log log.Logger, gasLimit, gasPrice *big.Int) *Connection {
	return &Connection{
		endpoint:    endpoint,
		kp:          kp,
		gasLimit:    gasLimit,
		maxGasPrice: gasPrice,
		log:         log,
		stop:        make(chan int),
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to tron chain...", "url", c.endpoint)

	grpcClient := tron.NewGrpcClientWithTimeout(c.endpoint, 100*time.Second)
	err := grpcClient.Start()
	if err != nil {
		return err
	}
	c.conn = grpcClient
	return nil
}

func (c *Connection) Keypair() *secp256k1.Keypair {
	return c.kp
}

func (c *Connection) Client() *tron.GrpcClient {
	return c.conn
}

// LatestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*big.Int, error) {
	block, err := c.conn.GetNowBlock()
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetInt64(block.BlockHeader.RawData.Number), nil
}

// EnsureHasBytecode asserts if contract code exists at the specified address
func (c *Connection) EnsureHasBytecode(contractAddress string) error {
	abi, err := c.conn.GetContractABI(contractAddress)
	if err != nil {
		return err
	}

	if abi == nil {
		return fmt.Errorf("no abi found at %s", contractAddress)
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
		c.conn.Stop()
	}
	close(c.stop)
}
