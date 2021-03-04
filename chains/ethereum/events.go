package ethereum

import (
	"github.com/emit-technology/emit-cross/types"
	"math/big"
)

func (l *listener) handleErc20DepositedEvent(blockNumer uint64, destId types.ChainId, nonce types.Nonce) (types.TransferMsg, error) {
	l.log.Info("Handling Erc20 fungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)
	resourceId, recipient, amount, err := l.writer.GetDepositFTRecord(uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return types.TransferMsg{}, err
	}

	return types.NewFungibleTransferMsg(blockNumer, l.cfg.id, destId, nonce, amount, types.ResourceId(resourceId), recipient), nil

}

func (l *listener) handleErc721DepositedEvent(blockNumer uint64, destId types.ChainId, nonce types.Nonce) (types.TransferMsg, error) {
	l.log.Info("Handling Erc721 nonfungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)
	resourceId, recipient, amount, metadata, _, err := l.writer.GetDepositNFTRecord(uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return types.TransferMsg{}, err
	}
	return types.NewNonFungibleTransferMsg(blockNumer, l.cfg.id, destId, nonce, amount, types.ResourceId(resourceId), recipient, metadata, big.NewInt(0)), nil

}
