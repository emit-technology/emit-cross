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
	"encoding/json"
	"errors"
	"github.com/emit-technology/emit-cross/types"
	"github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"math/big"
	"strconv"
)

func (w *writer) GetDepositRecord(nonce uint64, dest uint8) (resourceID [32]byte, destinationRecipientAddress []byte, amount *big.Int, err error) {
	params := []abi.Param{
		{"uint64": strconv.FormatUint(nonce, 10)},
		{"uint8": strconv.FormatUint(uint64(dest), 10)},
	}
	paramsJson, err := json.Marshal(params)

	if err != nil {
		w.log.Error("Failed to Marshal params", "err", err)
		return [32]byte{}, nil, nil, err
	}

	result, err := w.conn.Client().TriggerConstantContract("", w.cfg.trc20HandlerContract, "getDepositRecord(uint64,uint8)", string(paramsJson))
	if err != nil {
		w.log.Error("Failed to TriggerConstantContract", "err", err)
		return [32]byte{}, nil, nil, err
	}
	if result.Result.Code > 0 || result.Result.Message != nil {
		w.log.Error("Failed to TriggerConstantContract", "err", string(result.Result.Message))

		return [32]byte{}, nil, nil, errors.New(string(result.Result.Message))
	}
	recrod, err := w.trc20HandlerAbi.GetDepositRecord(result.GetConstantResult()[0])
	if err != nil {
		return [32]byte{}, nil, nil, err

	}
	copy(resourceID[:], recrod.ResourceID[:])
	destinationRecipientAddress = recrod.DestinationRecipientAddress
	amount = recrod.Amount
	return
}

func (w *writer) PropsalDataHash(recipient []byte, amount *big.Int) [32]byte {
	return ConstructTrc20ProposalDataHash(w.cfg.trc20HandlerContract, recipient, amount)
}
func (w *writer) GetProposalStatus(source uint8, nonce uint64, dataHash [32]byte) (uint8, error) {
	return w.prosalStatus(types.ChainId(source), types.Nonce(nonce), dataHash)
}
func (w *writer) GetBridgeAddress() []byte {
	addr, _ := address.Base58ToAddress(w.cfg.bridgeContract)
	return addr[:]
}
func (w *writer) IsWithCollector() bool {
	return true
}
