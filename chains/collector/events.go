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

package collector

import (
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/emit-technology/emit-cross/chains"
)

func (l *listener) handleSignProposalEvent(source, destId msg.ChainId, nonce msg.Nonce, sign []byte) (msg.Message, error) {
	l.log.Info("Handling Sign Proposal event", "src", source, "dest", destId, "nonce", nonce)

	resourceID, destinationRecipientAddress, amount, err := l.state[uint8(source)].GetDepositRecord(uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "err", err)
		return msg.Message{}, err
	}
	return chains.NewCommitSignature(
		source,
		destId,
		nonce,
		amount,
		resourceID,
		destinationRecipientAddress,
		sign,
	), nil
}
