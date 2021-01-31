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
	"github.com/emit-technology/emit-cross/types"
)

func (l *listener) handleSrc20DepositedEvent(blockNumer uint64, destId types.ChainId, nonce types.Nonce) (types.FTTransfer, error) {
	l.log.Info("Handling Src20 fungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)

	resourcId, recipeint, amount, err := l.state[l.cfg.id].GetDepositRecord(uint64(nonce), uint8(destId))

	if err != nil {
		l.log.Error("Error Unpacking SRC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return types.FTTransfer{}, err
	}

	return types.FTTransfer{
		blockNumer,
		l.cfg.id,
		destId,
		nonce,
		types.ResourceId(resourcId),
		recipeint,
		amount,
	}, nil
}

func (l *listener) handleProposalEvent(blockNumer uint64, soruceId types.ChainId, nonce types.Nonce) (types.FTTransfer, error) {

	l.log.Info("Handling passed Proposal event", "src", soruceId, "dest", l.cfg.id, "nonce", nonce)

	resourcId, recipeint, amount, err := l.state[soruceId].GetDepositRecord(uint64(nonce), uint8(l.cfg.id))

	if err != nil {
		l.log.Error("Error Unpacking Soruce Deposit Record", "src", soruceId, "dest", l.cfg.id, "nonce", nonce, "err", err)
		return types.FTTransfer{}, err
	}

	return types.FTTransfer{
		blockNumer,
		soruceId,
		l.cfg.id,
		nonce,
		types.ResourceId(resourcId),
		recipeint,
		amount,
	}, nil
}

func (l *listener) handleDestProposalEvent(blockNumer uint64, soruceId types.ChainId, destId types.ChainId, nonce types.Nonce) (uint8, types.BatchVotes, error) {
	l.log.Info("Handling dest passed SignProposal event", "src", soruceId, "dest", destId, "nonce", nonce)
	resourcId, recipeint, amount, err := l.state[soruceId].GetDepositRecord(uint64(nonce), uint8(destId))

	if err != nil {
		l.log.Error("Error Unpacking Soruce Deposit Record", "src", soruceId, "dest", destId, "nonce", nonce, "err", err)
		return 0, types.BatchVotes{}, err
	}
	status, signatures, err := l.writer.GetProposalSignatures(uint8(soruceId), uint8(destId), uint64(nonce))

	if err != nil {
		l.log.Error("Error Unpacking GetProposalSignatures", "src", soruceId, "dest", destId, "nonce", nonce, "err", err)

		return 0, types.BatchVotes{}, err
	}

	return status, types.BatchVotes{
		blockNumer,
		soruceId,
		destId,
		nonce,
		types.ResourceId(resourcId),
		recipeint,
		amount,
		signatures,
	}, nil
}
