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
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	"github.com/emit-technology/emit-cross/chains"
	"math/big"

	"github.com/emit-technology/emit-cross/common"

	"github.com/emit-technology/emit-cross/bindings/sero"
	"github.com/sero-cash/go-sero/accounts/abi/bind"
	seroCommon "github.com/sero-cash/go-sero/common"

	"github.com/ChainSafe/chainbridge-utils/blockstore"
	"github.com/ChainSafe/chainbridge-utils/core"

	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	bridge "github.com/emit-technology/emit-cross/bindings/sero/Bridge"
	src20Handler "github.com/emit-technology/emit-cross/bindings/sero/SRC20Handler"

	connection "github.com/emit-technology/emit-cross/connections/sero"
)

var _ core.Chain = &Chain{}

var _ Connection = &connection.Connection{}

type Connection interface {
	Connect() error
	Keypair() *secp256k1.Keypair
	Opts() *bind.TransactOpts
	CallOpts() *bind.CallOpts
	LockAndUpdateOpts() error
	UnlockOpts()
	Client() *sero.WrappedClient
	EnsureHasBytecode(address seroCommon.Address) error
	LatestBlock() (*big.Int, error)
	WaitForBlock(block *big.Int) error
	Close()
}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     Connection        // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
}

// checkBlockstore queries the blockstore for the latest known block. If the latest block is
// greater than cfg.startBlock, then cfg.startBlock is replaced with the latest known block.
func setupBlockstore(cfg *Config, kp *secp256k1.Keypair) (*blockstore.Blockstore, error) {
	bs, err := blockstore.NewBlockstore(cfg.blockstorePath, cfg.id, common.GenCommonAddress(kp).String())
	if err != nil {
		return nil, err
	}

	if !cfg.freshStart {
		latestBlock, err := bs.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}
		if latestBlock != nil && latestBlock.Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = latestBlock
		}
	}

	return bs, nil
}

func InitializeChain(chainCfg *core.ChainConfig, logger log15.Logger, sysErr chan<- error, m *metrics.ChainMetrics) (*Chain, chains.ProposalState, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, nil, err
	}

	kpI, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath, chainCfg.Insecure)

	if err != nil {
		return nil, nil, err
	}

	kp, _ := kpI.(*secp256k1.Keypair)

	bs, err := setupBlockstore(cfg, kp)
	if err != nil {
		return nil, nil, err
	}

	stop := make(chan int)

	conn := connection.NewConnection(cfg.endpoint, cfg.accountEndpoint, cfg.http, kp, logger, cfg.gasLimit, cfg.maxGasPrice)
	err = conn.Connect()
	if err != nil {
		return nil, nil, err
	}
	err = conn.EnsureHasBytecode(cfg.bridgeContract)
	if err != nil {
		return nil, nil, err
	}
	err = conn.EnsureHasBytecode(cfg.src20HandlerContract)
	if err != nil {
		return nil, nil, err
	}

	bridgeContract, err := bridge.NewBridge(cfg.bridgeContract, conn.Client())
	if err != nil {
		return nil, nil, err
	}

	chainId, err := bridgeContract.ChainID(conn.CallOpts())
	if err != nil {
		return nil, nil, err
	}

	if chainId != uint8(chainCfg.Id) {
		return nil, nil, fmt.Errorf("chainId (%d) and configuration chainId (%d) do not match", chainId, chainCfg.Id)
	}

	src20HandlerContract, err := src20Handler.NewSRC20Handler(cfg.src20HandlerContract, conn.Client())
	if err != nil {
		return nil, nil, err
	}

	if chainCfg.LatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, nil, err
		}
		cfg.startBlock = curr
	}

	listener := NewListener(conn, cfg, logger, bs, stop, sysErr, m)
	listener.setContracts(bridgeContract, src20HandlerContract)

	writer := NewWriter(conn, cfg, logger, stop, sysErr, m)
	writer.setContract(bridgeContract, src20HandlerContract)

	listener.setWriter(writer)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, writer, nil
}

func (c *Chain) SetRouter(r *core.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.setRouter(r)
}

func (c *Chain) Start() error {
	err := c.listener.start()
	if err != nil {
		return err
	}

	err = c.writer.start()
	if err != nil {
		return err
	}

	c.writer.log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Name() string {
	return c.cfg.Name
}

func (c *Chain) LatestBlock() metrics.LatestBlock {
	return c.listener.latestBlock
}

// Stop signals to any running routines to exit
func (c *Chain) Stop() {
	close(c.stop)
	if c.conn != nil {
		c.conn.Close()
	}
}
