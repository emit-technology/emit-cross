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
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/emit-technology/emit-cross/chains"
	"github.com/emit-technology/emit-cross/common"
	"github.com/sero-cash/go-sero/accounts/abi/bind"
	"math/big"
)

func (l *listener) handleSrc20DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling Src20 fungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)

	record, err := l.src20HandlerContract.GetDepositRecord(&bind.CallOpts{FromPKr: common.GenMainPkr(l.conn.Keypair())}, uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking SRC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return msg.Message{}, err
	}

	return msg.Message{
		Source:       l.cfg.id,
		Destination:  msg.ChainId(chains.CollectorChainId),
		Type:         chains.SignProposalReq,
		DepositNonce: nonce,
		ResourceId:   record.ResourceID,
		Payload: []interface{}{
			record.Amount.Bytes(),
			record.DestinationRecipientAddress,
			big.NewInt(0).SetUint64(uint64(destId)).Bytes(),
		},
	}, nil

}
