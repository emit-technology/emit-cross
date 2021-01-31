// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"github.com/emit-technology/emit-cross/bindings/ethereum/Bridge"
	"github.com/emit-technology/emit-cross/bindings/ethereum/ERC20Handler"
	log "github.com/emit-technology/emit-cross/log"
	metrics "github.com/emit-technology/emit-cross/metrics/types"
)

var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4

type writer struct {
	cfg                  Config
	conn                 Connection
	bridgeContract       *Bridge.Bridge // instance of bound receiver bridgeContract
	erc20HandlerContract *ERC20Handler.ERC20Handler
	log                  log.Logger
	stop                 <-chan int
	sysErr               chan<- error // Reports fatal error to core
	metrics              *metrics.ChainMetrics
}

// NewWriter creates and returns writer
func NewWriter(conn Connection, cfg *Config, log log.Logger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *writer {
	return &writer{
		cfg:     *cfg,
		conn:    conn,
		log:     log,
		stop:    stop,
		sysErr:  sysErr,
		metrics: m,
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
