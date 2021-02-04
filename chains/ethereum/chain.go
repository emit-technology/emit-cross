// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
The ethereum package contains the logic for interacting with ethereum chains.

There are 3 major components: the connection, the listener, and the writer.
The currently supported transfer types are Fungible (ERC20), Non-Fungible (ERC721), and generic.

Connection

The connection contains the ethereum RPC client and can be accessed by both the writer and listener.

Listener

The listener polls for each new block and looks for deposit events in the bridge contract. If a deposit occurs, the listener will fetch additional information from the Handler before constructing a message and forwarding it to the router.

Writer

The writer recieves the message and creates a proposals on-chain. Once a proposal is made, the writer then watches for a finalization event and will attempt to execute the proposal if a matching event occurs. The writer skips over any proposals it has already seen.
*/
package ethereum

import (
	"fmt"
	"github.com/emit-technology/emit-cross/chains"
	"github.com/emit-technology/emit-cross/core"
	"github.com/emit-technology/emit-cross/crypto/secp256k1"
	log15 "github.com/emit-technology/emit-cross/log"
	metrics "github.com/emit-technology/emit-cross/metrics/types"
	"github.com/emit-technology/emit-cross/types"
	"math/big"

	bridge "github.com/emit-technology/emit-cross/bindings/ethereum/Bridge"
	erc20Handler "github.com/emit-technology/emit-cross/bindings/ethereum/ERC20Handler"
	connection "github.com/emit-technology/emit-cross/connections/ethereum"
	"github.com/emit-technology/emit-cross/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	Client() *ethclient.Client
	EnsureHasBytecode(address common.Address) error
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

func InitializeChain(chainCfg *core.ChainConfig, chainDB *chains.ChainDB, logger log15.Logger, sysErr chan<- error, m *metrics.ChainMetrics) (*Chain, core.ProposalState, error) {
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
	conn := connection.NewConnection(cfg.endpoint, cfg.http, kp, logger, cfg.gasLimit, cfg.maxGasPrice)
	err = conn.Connect()
	if err != nil {
		return nil, nil, err
	}
	err = conn.EnsureHasBytecode(cfg.bridgeContract)
	if err != nil {
		return nil, nil, err
	}
	err = conn.EnsureHasBytecode(cfg.erc20HandlerContract)
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

	erc20HandlerContract, err := erc20Handler.NewERC20Handler(cfg.erc20HandlerContract, conn.Client())
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
	writer.setContract(bridgeContract, erc20HandlerContract)
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
		next, err := chainDB.GetNextPollBlockNum(uint8(cfg.id), kp.Address())
		if err != nil {
			return err
		}

		if new(big.Int).SetUint64(next).Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = new(big.Int).SetUint64(next)
		}

		lastId, err := chainDB.GetLastBatchVotesId(uint8(cfg.id), kp.Address())
		if err != nil {
			return err
		}

		if new(big.Int).SetUint64(lastId+1).Cmp(cfg.commitVotesStartSeq) == 1 {
			cfg.commitVotesStartSeq = new(big.Int).SetUint64(lastId + 1)
		}

	}
	return nil
}

func (c *Chain) Start() error {
	c.listener.log.Info("staring...")
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
