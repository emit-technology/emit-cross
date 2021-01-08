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

import "C"
import (
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sero-cash/go-sero/common"
	"math/big"
	"sync"
)

func CommonHash(data []byte) (ret common.Hash) {
	h := crypto.Keccak256Hash(data)
	copy(ret[:], h[:])
	return
}

type VoteSet map[common.Hash][]byte

func (self VoteSet) copy() (ret VoteSet) {
	ret = make(VoteSet)
	for k, v := range self {
		ret[k] = v
	}
	return
}

type PendingVote struct {
	pendingVoteMu sync.RWMutex
	pendingVote   map[common.Hash]VoteSet
}

func NewPendingVote() (ret PendingVote) {
	ret.pendingVote = make(map[common.Hash]VoteSet)
	return ret
}

func (self *PendingVote) Add(m msg.Message) bool {
	self.pendingVoteMu.Lock()
	defer self.pendingVoteMu.Unlock()

	key := MsgToHash(m)

	newProposal := false

	var vs VoteSet
	if _, ok := self.pendingVote[key]; !ok {
		vs = make(VoteSet)
		self.pendingVote[key] = vs
		newProposal = true
	} else {
		vs = self.pendingVote[key]
	}

	sign := m.Payload[2].([]byte)
	signHash := CommonHash(sign)
	vs[signHash] = sign
	return newProposal

}

func (self *PendingVote) GetPending(m msg.Message) (ret VoteSet) {
	self.pendingVoteMu.RLock()
	defer self.pendingVoteMu.RUnlock()
	key := MsgToHash(m)
	if votes, ok := self.pendingVote[key]; ok {
		ret = votes.copy()
	}
	return
}

func (self *PendingVote) Delete(m msg.Message) {
	self.pendingVoteMu.Lock()
	defer self.pendingVoteMu.Unlock()
	key := MsgToHash(m)
	delete(self.pendingVote, key)
}

func MsgToHash(m msg.Message) common.Hash {

	var data []byte

	data = append(data, new(big.Int).SetUint64(uint64(m.Source)).Bytes()...)
	data = append(data, new(big.Int).SetUint64(uint64(m.Destination)).Bytes()...)
	data = append(data, new(big.Int).SetUint64(uint64(m.DepositNonce)).Bytes()...)
	data = append(data, m.ResourceId[:]...)
	data = append(data, m.Payload[1].([]byte)...)
	data = append(data, m.Payload[0].([]byte)...)

	return CommonHash(data)

}
