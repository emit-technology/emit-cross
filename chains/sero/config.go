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
	"errors"
	"fmt"
	"github.com/emit-technology/emit-cross/core"
	"github.com/emit-technology/emit-cross/types"
	"math/big"

	"github.com/sero-cash/go-sero/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 10000000000
const DefaultBlockConfirmations = 12

// Chain specific options
var (
	AccountEndpointOpt    = "accountEndpoint"
	BridgeOpt             = "bridge"
	src20HandlerOpt       = "src20Handler"
	commitNodeOpt         = "commitNode"
	MaxGasPriceOpt        = "maxGasPrice"
	GasLimitOpt           = "gasLimit"
	HttpOpt               = "http"
	StartBlockOpt         = "startBlock"
	BlockConfirmationsOpt = "blockConfirmations"
	signatureColletorOpt  = "signatureColletor"
)

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	name                     string        // Human-readable chain name
	id                       types.ChainId // ChainID
	endpoint                 string        // url for rpc endpoint
	accountEndpoint          string
	from                     string // address of key to use
	keystorePath             string // Location of keyfiles
	blockstorePath           string
	freshStart               bool // Disables loading from blockstore at start
	bridgeContract           common.Address
	src20HandlerContract     common.Address
	signatureColletorContact common.Address
	gasLimit                 *big.Int
	maxGasPrice              *big.Int
	http                     bool // Config for type of connection
	startBlock               *big.Int
	blockConfirmations       *big.Int
	commitNode               bool
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:                     chainCfg.Name,
		id:                       chainCfg.Id,
		endpoint:                 chainCfg.Endpoint,
		from:                     chainCfg.From,
		keystorePath:             chainCfg.KeystorePath,
		freshStart:               chainCfg.FreshStart,
		bridgeContract:           common.Address{},
		src20HandlerContract:     common.Address{},
		signatureColletorContact: common.Address{},
		gasLimit:                 big.NewInt(DefaultGasLimit),
		maxGasPrice:              big.NewInt(DefaultGasPrice),
		http:                     false,
		startBlock:               big.NewInt(0),
		blockConfirmations:       big.NewInt(0),
	}

	if contract, ok := chainCfg.Opts[BridgeOpt]; ok && contract != "" {
		config.bridgeContract = common.Base58ToAddress(contract)
		delete(chainCfg.Opts, BridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for sero config")
	}

	if accountEndpoint, ok := chainCfg.Opts[AccountEndpointOpt]; ok && accountEndpoint != "" {
		config.accountEndpoint = accountEndpoint
		delete(chainCfg.Opts, AccountEndpointOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.accountEndpoint field for sero config")
	}

	config.src20HandlerContract = common.Base58ToAddress(chainCfg.Opts[src20HandlerOpt])
	delete(chainCfg.Opts, src20HandlerOpt)
	if contract, ok := chainCfg.Opts[signatureColletorOpt]; ok && contract != "" {
		config.signatureColletorContact = common.Base58ToAddress(contract)
		delete(chainCfg.Opts, signatureColletorOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.signatureColletor field for sero config")
	}
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

	if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "true" {
		config.http = true
		delete(chainCfg.Opts, HttpOpt)
	} else if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "false" {
		config.http = false
		delete(chainCfg.Opts, HttpOpt)
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
