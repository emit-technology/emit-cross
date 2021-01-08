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

package chains

import (
	"github.com/ChainSafe/chainbridge-utils/msg"
	"math/big"
)

var CollectorChainId = uint8(0)

var SignProposalReq msg.TransferType = "SignProposalReq"

type Proposal struct {
	Source       msg.ChainId // Source where message was initiated
	Destination  msg.ChainId // Destination chain of message
	DepositNonce msg.Nonce   // Nonce for the deposit
	ResourceId   msg.ResourceId
	Recipient    []byte
	Amount       *big.Int
}

func NewCommitSignature(source, dest msg.ChainId, nonce msg.Nonce, amount *big.Int, resourceId msg.ResourceId, recipient []byte, sign []byte) msg.Message {
	return msg.Message{
		Source:       source,
		Destination:  dest,
		Type:         msg.FungibleTransfer,
		DepositNonce: nonce,
		ResourceId:   resourceId,
		Payload: []interface{}{
			amount.Bytes(),
			recipient,
			sign,
		},
	}
}

func NewProposal(m msg.Message) Proposal {
	return Proposal{
		Source:       m.Source,
		Destination:  msg.ChainId(new(big.Int).SetBytes(m.Payload[2].([]byte)).Uint64()),
		DepositNonce: m.DepositNonce,
		ResourceId:   m.ResourceId,
		Recipient:    m.Payload[1].([]byte),
		Amount:       new(big.Int).SetBytes(m.Payload[0].([]byte)),
	}
}
