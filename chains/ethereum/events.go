// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (l *listener) handleErc20DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling Erc20 fungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)

	record, err := l.erc20HandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return msg.Message{}, err
	}

	return msg.NewFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.Amount,
		record.ResourceID,
		record.DestinationRecipientAddress,
	), nil
}
