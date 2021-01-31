package ethereum

import (
	"github.com/emit-technology/emit-cross/types"
)

func (l *listener) handleErc20DepositedEvent(blockNumer uint64, destId types.ChainId, nonce types.Nonce) (types.FTTransfer, error) {
	l.log.Info("Handling Erc20 fungible deposit event", "src", l.cfg.id, "dest", destId, "nonce", nonce)
	resourceId, recipient, amount, err := l.writer.GetDepositRecord(uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "src", l.cfg.id, "dest", destId, "nonce", nonce, "err", err)
		return types.FTTransfer{}, err
	}
	return types.FTTransfer{
		blockNumer,
		l.cfg.id,
		destId,
		nonce,
		types.ResourceId(resourceId),
		recipient,
		amount,
	}, nil

}
