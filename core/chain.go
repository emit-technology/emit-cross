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

package core

import (
	metrics "github.com/emit-technology/emit-cross/metrics/types"
	"github.com/emit-technology/emit-cross/types"
)

type Chain interface {
	Start() error // Start chain
	Id() types.ChainId
	LatestBlock() metrics.LatestBlock
	SetState(state map[types.ChainId]ProposalState)
	Name() string
	Stop()
}

type ChainConfig struct {
	Name         string            // Human-readable chain name
	Id           types.ChainId     // ChainID
	Endpoint     string            // url for rpc endpoint
	From         string            // address of key to use
	KeystorePath string            // Location of key files
	Insecure     bool              // Indicated whether the test keyring should be used
	FreshStart   bool              // If true, blockstore is ignored at start.
	LatestBlock  bool              // If true, overrides blockstore or latest block in config and starts from current block
	Opts         map[string]string // Per chain options
}
