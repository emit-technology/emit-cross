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
	"math/big"
)

func (l *listener) handleSrc20DepositedEvent(blockNumer uint64, destId types.ChainId, nonce types.Nonce) (types.TransferMsg, error) {

	l.log.Info("Handling Src20 fungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)

	resourcId, recipeint, amount, err := l.state[l.cfg.id].GetDepositFTRecord(uint64(nonce), uint8(destId))

	if err != nil {
		l.log.Error("Error Unpacking SRC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return types.TransferMsg{}, err
	}
	return types.NewFungibleTransferMsg(blockNumer, l.cfg.id, destId, nonce, amount, types.ResourceId(resourcId), recipeint), nil

}

func (l *listener) handleSrc721DepositedEvent(blockNumer uint64, destId types.ChainId, nonce types.Nonce) (types.TransferMsg, error) {
	l.log.Info("Handling Src721 nonFungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)

	resourcId, recipeint, amount, metaData, src20Amount, err := l.state[l.cfg.id].GetDepositNFTRecord(uint64(nonce), uint8(destId))

	if err != nil {
		l.log.Error("Error Unpacking SRC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return types.TransferMsg{}, err
	}
	return types.NewNonFungibleTransferMsg(
		blockNumer,
		l.cfg.id,
		destId,
		nonce,
		amount,
		types.ResourceId(resourcId),
		recipeint,
		metaData,
		src20Amount,
	), nil
}

func (l *listener) handleSrc20ProposalEvent(blockNumer uint64, soruceId types.ChainId, nonce types.Nonce) (types.TransferMsg, error) {

	l.log.Info("Handling passed src20 Proposal event", "src", soruceId, "dest", l.cfg.id, "nonce", nonce)

	resourcId, recipeint, amount, err := l.state[soruceId].GetDepositFTRecord(uint64(nonce), uint8(l.cfg.id))

	if err != nil {
		l.log.Error("Error Unpacking Soruce Deposit Record", "src", soruceId, "dest", l.cfg.id, "nonce", nonce, "err", err)
		return types.TransferMsg{}, err
	}

	return types.NewFungibleTransferMsg(blockNumer, soruceId, l.cfg.id, nonce, amount, types.ResourceId(resourcId), recipeint), nil
}

func (l *listener) handleSrc721ProposalEvent(blockNumer uint64, soruceId types.ChainId, nonce types.Nonce) (types.TransferMsg, error) {

	l.log.Info("Handling passed src721 Proposal event", "src", soruceId, "dest", l.cfg.id, "nonce", nonce)

	resourcId, recipeint, amount, metadata, _, err := l.state[soruceId].GetDepositNFTRecord(uint64(nonce), uint8(l.cfg.id))

	if err != nil {
		l.log.Error("Error Unpacking Soruce Deposit Record", "src", soruceId, "dest", l.cfg.id, "nonce", nonce, "err", err)
		return types.TransferMsg{}, err
	}

	return types.NewNonFungibleTransferMsg(blockNumer, soruceId, l.cfg.id, nonce, amount, types.ResourceId(resourcId), recipeint, metadata, big.NewInt(0)), nil
}

func (l *listener) handleDestFungibleProposalSignEvent(blockNumer uint64, soruceId types.ChainId, destId types.ChainId, nonce types.Nonce) (uint8, types.ProposalSignatures, error) {
	l.log.Info("Handling dest passed SignProposal event", "src", soruceId, "dest", destId, "nonce", nonce)
	resourcId, recipeint, amount, err := l.state[soruceId].GetDepositFTRecord(uint64(nonce), uint8(destId))

	if err != nil {
		l.log.Error("Error Unpacking Soruce Deposit Record", "src", soruceId, "dest", destId, "nonce", nonce, "err", err)
		return 0, types.ProposalSignatures{}, err
	}
	status, signatures, err := l.writer.GetProposalSignatures(uint8(soruceId), uint8(destId), uint64(nonce))

	if err != nil {
		l.log.Error("Error Unpacking GetProposalSignatures", "src", soruceId, "dest", destId, "nonce", nonce, "err", err)

		return 0, types.ProposalSignatures{}, err
	}

	return status, types.NewFungibleProposalSignatures(blockNumer, soruceId, destId, nonce, amount, types.ResourceId(resourcId), recipeint, signatures), nil

}

func (l *listener) handleDestNonFungibleProposalSignEvent(blockNumer uint64, soruceId types.ChainId, destId types.ChainId, nonce types.Nonce) (uint8, types.ProposalSignatures, error) {
	l.log.Info("Handling dest passed SignProposal event", "src", soruceId, "dest", destId, "nonce", nonce)
	resourceId, recipeint, amount, metadata, src20Amount, err := l.state[soruceId].GetDepositNFTRecord(uint64(nonce), uint8(destId))

	if err != nil {
		l.log.Error("Error Unpacking Soruce Deposit Record", "src", soruceId, "dest", destId, "nonce", nonce, "err", err)
		return 0, types.ProposalSignatures{}, err
	}
	status, signatures, err := l.writer.GetProposalSignatures(uint8(soruceId), uint8(destId), uint64(nonce))

	if err != nil {
		l.log.Error("Error Unpacking GetProposalSignatures", "src", soruceId, "dest", destId, "nonce", nonce, "err", err)

		return 0, types.ProposalSignatures{}, err
	}

	return status, types.NewNonFungibleProposalSignatures(blockNumer,
		soruceId,
		destId,
		nonce,
		amount,
		types.ResourceId(resourceId),
		recipeint,
		metadata,
		src20Amount, signatures), nil

}
