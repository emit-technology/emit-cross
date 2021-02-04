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
	"github.com/emit-technology/emit-cross/bindings/sero/Collector"
	"github.com/emit-technology/emit-cross/chains"
	"github.com/emit-technology/emit-cross/common"
	"github.com/emit-technology/emit-cross/core"
	"github.com/emit-technology/emit-cross/crypto/secp256k1"
	"github.com/emit-technology/emit-cross/keystore"
	log "github.com/emit-technology/emit-cross/log"
	metrics "github.com/emit-technology/emit-cross/metrics/types"
	"github.com/emit-technology/emit-cross/types"
	"math/big"

	"github.com/emit-technology/emit-cross/bindings/sero"
	"github.com/sero-cash/go-sero/accounts/abi/bind"
	seroCommon "github.com/sero-cash/go-sero/common"

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
	stop     chan<- int
}

func InitializeChain(chainCfg *core.ChainConfig, chainDB *chains.ChainDB, logger log.Logger, sysErr chan<- error, m *metrics.ChainMetrics) (*Chain, core.ProposalState, error) {
	cfg, err := parseChainConfig(chainCfg)

	if err != nil {
		return nil, nil, err
	}

	kpI, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath, chainCfg.Insecure)

	if err != nil {
		return nil, nil, err
	}

	kp, _ := kpI.(*secp256k1.Keypair)

	err = setupStart(cfg, chainDB, kp)
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

	signatureCollectorContract, err := Collector.NewSignatureCollector(cfg.signatureColletorContact, conn.Client())
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

	listener := NewListener(conn, chainDB, cfg, logger, stop, sysErr, m)

	writer := NewWriter(conn, cfg, logger, stop, sysErr, m)
	writer.setContract(bridgeContract, src20HandlerContract, signatureCollectorContract)

	listener.setWriter(writer)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		listener: listener,
		stop:     stop,
	}, writer, nil
}

func setupStart(cfg *Config, chainDB *chains.ChainDB, kp *secp256k1.Keypair) error {

	if !cfg.freshStart {
		next, err := chainDB.GetNextPollBlockNum(uint8(cfg.id), common.GenCommonAddress(kp).String())
		if err != nil {
			return err
		}

		if new(big.Int).SetUint64(next).Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = new(big.Int).SetUint64(next)
		}

		lastSignSeq, err := chainDB.GetLastSignId(uint8(cfg.id), common.GenCommonAddress(kp).String())
		if err != nil {
			return err
		}

		if new(big.Int).SetUint64(lastSignSeq+1).Cmp(cfg.SignMsgStartSeq) == 1 {
			cfg.SignMsgStartSeq = new(big.Int).SetUint64(lastSignSeq + 1)
		}

		lastVoteSeq, err := chainDB.GetLastVoteId(uint8(cfg.id), common.GenCommonAddress(kp).String())
		if err != nil {
			return err
		}

		if new(big.Int).SetUint64(lastVoteSeq+1).Cmp(cfg.VoteProposalStartSeq) == 1 {
			cfg.VoteProposalStartSeq = new(big.Int).SetUint64(lastVoteSeq + 1)
		}

		lastExecuteSeq, err := chainDB.GetLastExecuteId(uint8(cfg.id), common.GenCommonAddress(kp).String())
		if err != nil {
			return err
		}

		if new(big.Int).SetUint64(lastExecuteSeq+1).Cmp(cfg.ExecuteProposalStartSeq) == 1 {
			cfg.ExecuteProposalStartSeq = new(big.Int).SetUint64(lastExecuteSeq + 1)
		}

	}
	return nil
}

func (c *Chain) Start() error {
	err := c.listener.start()
	if err != nil {
		return err
	}

	c.listener.log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) Id() types.ChainId {
	return c.cfg.Id
}

func (c *Chain) Name() string {
	return c.cfg.Name
}

func (c *Chain) LatestBlock() metrics.LatestBlock {
	return c.listener.latestBlock
}

func (c *Chain) SetState(state map[types.ChainId]core.ProposalState) {
	c.listener.state = state
}

// Stop signals to any running routines to exit
func (c *Chain) Stop() {
	close(c.stop)
	if c.conn != nil {
		c.conn.Close()
	}
}
