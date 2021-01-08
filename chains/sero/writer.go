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
	"github.com/ChainSafe/chainbridge-utils/core"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/emit-technology/emit-cross/bindings/sero/Bridge"
	src20Handler "github.com/emit-technology/emit-cross/bindings/sero/SRC20Handler"
	"github.com/emit-technology/emit-cross/chains"
)

var _ core.Writer = &writer{}

// https://github.com/ChainSafe/chainbridge-solidity/blob/b5ed13d9798feb7c340e737a726dd415b8815366/contracts/Bridge.sol#L20
var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4

type writer struct {
	cfg                  Config
	conn                 Connection
	bridgeContract       *Bridge.Bridge // instance of bound receiver bridgeContract
	src20HandlerContract *src20Handler.SRC20Handler
	log                  log15.Logger
	stop                 <-chan int
	sysErr               chan<- error // Reports fatal error to core
	metrics              *metrics.ChainMetrics
	pendingVote          chains.PendingVote
	state                map[uint8]chains.ProposalState
}

// NewWriter creates and returns writer
func NewWriter(conn Connection, cfg *Config, log log15.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *writer {
	return &writer{
		cfg:         *cfg,
		conn:        conn,
		log:         log,
		stop:        stop,
		sysErr:      sysErr,
		metrics:     m,
		pendingVote: chains.NewPendingVote(),
	}
}

func (w *writer) start() error {
	w.log.Debug("Starting sero writer...")
	return nil
}

// setContract adds the bound receiver bridgeContract to the writer
func (w *writer) setContract(bridge *Bridge.Bridge, handler *src20Handler.SRC20Handler) {
	w.bridgeContract = bridge
	w.src20HandlerContract = handler
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *writer) ResolveMessage(m msg.Message) bool {
	w.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())

	switch m.Type {
	case msg.FungibleTransfer:
		return w.createErc20Proposal(m)
	default:
		w.log.Error("Unknown message type received", "type", m.Type)
		return false
	}
}
