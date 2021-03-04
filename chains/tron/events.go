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
	"github.com/emit-technology/emit-cross/types"
)

func (l *listener) handleTrc20DepositedEvent(blockNumer uint64, destId types.ChainId, nonce types.Nonce) (types.TransferMsg, error) {
	l.log.Info("Handling Trc20 fungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)

	resourcId, recipeint, amount, err := l.state[l.cfg.id].GetDepositFTRecord(uint64(nonce), uint8(destId))

	if err != nil {
		l.log.Error("Error Unpacking SRC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return types.TransferMsg{}, err
	}
	return types.NewFungibleTransferMsg(blockNumer, l.cfg.id, destId, nonce, amount, types.ResourceId(resourcId), recipeint), nil
}
