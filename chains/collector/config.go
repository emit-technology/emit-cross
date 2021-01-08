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

package collector

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/sero-cash/go-sero/common"

	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/msg"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 40000000000
const DefaultBlockConfirmations = 10

// Chain specific options
var (
	AccountEndpointOpt    = "accountEndpoint"
	signatureColletorOpt  = "signatureColletor"
	MaxGasPriceOpt        = "maxGasPrice"
	GasLimitOpt           = "gasLimit"
	HttpOpt               = "http"
	StartBlockOpt         = "startBlock"
	BlockConfirmationsOpt = "blockConfirmations"
)

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	name                     string      // Human-readable chain name
	id                       msg.ChainId // ChainID
	endpoint                 string      // url for rpc endpoint
	accountEndpoint          string
	from                     string // address of key to use
	keystorePath             string // Location of keyfiles
	blockstorePath           string
	freshStart               bool // Disables loading from blockstore at start
	signatureColletorContact common.Address
	gasLimit                 *big.Int
	maxGasPrice              *big.Int
	http                     bool // Config for type of connection
	startBlock               *big.Int
	blockConfirmations       *big.Int
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:                     chainCfg.Name,
		id:                       chainCfg.Id,
		endpoint:                 chainCfg.Endpoint,
		from:                     chainCfg.From,
		keystorePath:             chainCfg.KeystorePath,
		blockstorePath:           chainCfg.BlockstorePath,
		freshStart:               chainCfg.FreshStart,
		signatureColletorContact: common.Address{},
		gasLimit:                 big.NewInt(DefaultGasLimit),
		maxGasPrice:              big.NewInt(DefaultGasPrice),
		http:                     false,
		startBlock:               big.NewInt(0),
		blockConfirmations:       big.NewInt(0),
	}

	if accountEndpoint, ok := chainCfg.Opts[AccountEndpointOpt]; ok && accountEndpoint != "" {
		config.accountEndpoint = accountEndpoint
		delete(chainCfg.Opts, AccountEndpointOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.accountEndpoint field for sero config")
	}

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
