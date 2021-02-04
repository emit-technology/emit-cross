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
	"github.com/emit-technology/emit-cross/core"
	"github.com/emit-technology/emit-cross/types"
	"math/big"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 10000000000
const DefaultBlockConfirmations = 12

// Chain specific options
var (
	BridgeOpt              = "bridge"
	trc20HandlerOpt        = "trc20Handler"
	commitNodeOpt          = "commitNode"
	MaxGasPriceOpt         = "maxGasPrice"
	GasLimitOpt            = "gasLimit"
	StartBlockOpt          = "startBlock"
	BlockConfirmationsOpt  = "blockConfirmations"
	CommitVotesStartSeqOpt = "commitVotesStartSeq"
)

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	name                 string        // Human-readable chain name
	id                   types.ChainId // ChainID
	endpoint             string        // url for rpc endpoint
	accountEndpoint      string
	from                 string // address of key to use
	keystorePath         string // Location of keyfiles
	blockstorePath       string
	freshStart           bool // Disables loading from blockstore at start
	bridgeContract       string
	trc20HandlerContract string
	gasLimit             *big.Int
	maxGasPrice          *big.Int
	http                 bool // Config for type of connection
	startBlock           *big.Int
	commitVotesStartSeq  *big.Int
	blockConfirmations   *big.Int
	commitNode           bool
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:               chainCfg.Name,
		id:                 chainCfg.Id,
		endpoint:           chainCfg.Endpoint,
		from:               chainCfg.From,
		keystorePath:       chainCfg.KeystorePath,
		freshStart:         chainCfg.FreshStart,
		gasLimit:           big.NewInt(DefaultGasLimit),
		maxGasPrice:        big.NewInt(DefaultGasPrice),
		http:               false,
		startBlock:         big.NewInt(0),
		blockConfirmations: big.NewInt(0),
	}

	if contract, ok := chainCfg.Opts[BridgeOpt]; ok && contract != "" {
		config.bridgeContract = contract
		delete(chainCfg.Opts, BridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for sero config")
	}

	config.trc20HandlerContract = chainCfg.Opts[trc20HandlerOpt]
	delete(chainCfg.Opts, trc20HandlerOpt)

	if gasPrice, ok := chainCfg.Opts[MaxGasPriceOpt]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.maxGasPrice = price
			delete(chainCfg.Opts, MaxGasPriceOpt)
		} else {
			return nil, errors.New("unable to parse max gas price")
		}
	}

	if gasLimit, ok := chainCfg.Opts[GasLimitOpt]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.gasLimit = limit
			delete(chainCfg.Opts, GasLimitOpt)
		} else {
			return nil, errors.New("unable to parse gas limit")
		}
	}

	if commit, ok := chainCfg.Opts[commitNodeOpt]; ok && commit == "true" {
		config.commitNode = true
		delete(chainCfg.Opts, commitNodeOpt)
	} else if commit, ok := chainCfg.Opts[commitNodeOpt]; ok && commit == "false" {
		config.commitNode = false
		delete(chainCfg.Opts, commitNodeOpt)
	}

	if startBlock, ok := chainCfg.Opts[StartBlockOpt]; ok && startBlock != "" {
		block := big.NewInt(0)
		_, pass := block.SetString(startBlock, 10)
		if pass {
			config.startBlock = block
			delete(chainCfg.Opts, StartBlockOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", StartBlockOpt)
		}
	}

	if commitvotesStartSeq, ok := chainCfg.Opts[CommitVotesStartSeqOpt]; ok && commitvotesStartSeq != "" {
		start := big.NewInt(0)
		_, pass := start.SetString(commitvotesStartSeq, 10)
		if pass {
			config.commitVotesStartSeq = start
			delete(chainCfg.Opts, CommitVotesStartSeqOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", CommitVotesStartSeqOpt)
		}
	}

	if blockConfirmations, ok := chainCfg.Opts[BlockConfirmationsOpt]; ok && blockConfirmations != "" {
		val := big.NewInt(DefaultBlockConfirmations)
		_, pass := val.SetString(blockConfirmations, 10)
		if pass {
			config.blockConfirmations = val
			delete(chainCfg.Opts, BlockConfirmationsOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", BlockConfirmationsOpt)
		}
	}

	if len(chainCfg.Opts) != 0 {
		return nil, fmt.Errorf("unknown Opts Encountered: %#v", chainCfg.Opts)
	}

	return config, nil
}
