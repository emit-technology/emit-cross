// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"github.com/ChainSafe/chainbridge-utils/core"
	metrics "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	"github.com/emit-technology/emit-cross/bindings/ethereum/Bridge"
	"github.com/emit-technology/emit-cross/bindings/ethereum/ERC20Handler"
	"github.com/emit-technology/emit-cross/chains"
	"github.com/sero-cash/go-sero/common/hexutil"
	"math/big"
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
	erc20HandlerContract *ERC20Handler.ERC20Handler
	log                  log15.Logger
	stop                 <-chan int
	sysErr               chan<- error // Reports fatal error to core
	metrics              *metrics.ChainMetrics
	pendingVote          chains.PendingVote
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
func (w *writer) setContract(bridge *Bridge.Bridge, handler *ERC20Handler.ERC20Handler) {
	w.bridgeContract = bridge
	w.erc20HandlerContract = handler
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *writer) ResolveMessage(m msg.Message) bool {

	switch m.Type {
	case msg.FungibleTransfer:
		if w.cfg.commitNode {
			w.log.Info("Attempting to resolve  Erc20Proposal signature", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())
			return w.commitErc20ProposalSign(m)
		} else {

			w.log.Info("received erc20 proposal signature", "src", m.Source,
				"nonce", m.DepositNonce,
				"recipient", hexutil.Encode(m.Payload[1].([]byte)),
				"amount", new(big.Int).SetBytes(m.Payload[0].([]byte)).String())
			return true
		}
	default:
		w.log.Error("Unknown message type received", "type", m.Type)
		return false
	}
}
