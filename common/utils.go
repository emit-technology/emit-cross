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

package common

import (
	"encoding/binary"
	"github.com/emit-technology/emit-cross/crypto/secp256k1"
	"github.com/emit-technology/emit-cross/types"
	"github.com/sero-cash/go-czero-import/cpt"
	"github.com/sero-cash/go-sero/common/address"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sero-cash/go-czero-import/c_superzk"
	"github.com/sero-cash/go-czero-import/c_type"
	"github.com/sero-cash/go-czero-import/superzk"
	seroCommon "github.com/sero-cash/go-sero/common"
)

func init() {
	cpt.ZeroInit_NoCircuit()
}
func encodeNumber(number uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, number)
	return enc
}

func GenSeroTK(privKey []byte) address.TKAddress {
	var tk address.TKAddress
	var seed c_type.Uint256
	copy(seed[:], privKey)
	sk := superzk.Seed2Sk(&seed, 1)
	pubBytes, _ := superzk.Sk2Tk(&sk)
	copy(tk[:], pubBytes[:])
	return tk
}

func GenMainPkr(kp *secp256k1.Keypair) *c_type.PKr {
	privKey := crypto.FromECDSA(kp.PrivateKey())

	tk := GenSeroTK(privKey)

	salt := encodeNumber(1)
	random := append(tk[:], salt...)
	r := crypto.Keccak256Hash(random)
	var rand c_type.Uint256
	copy(rand[:], r[:])
	pk := tk.ToPk().ToUint512()
	return superzk.Pk2PKr(&pk, &rand).NewRef()
}

func GenCommonAddress(kp *secp256k1.Keypair) seroCommon.Address {
	pkr := GenMainPkr(kp)
	var result seroCommon.Address
	result.SetBytes(pkr[:])
	return result
}

func GenContractAddress(addr seroCommon.Address) seroCommon.ContractAddress {
	var conctractaddr seroCommon.ContractAddress
	pkr := new(c_type.PKr)
	copy(pkr[:], addr[:])
	hash := c_superzk.HashPKr(pkr)
	conctractaddr.SetBytes(hash[:])
	return conctractaddr
}

/*
func GenSeroKey(priKey *ecdsa.PrivateKey) *keystore.Key {
	tk := GenSeroTK(priKey)
	return &keystore.Key{
		Id:         uuid.NewUUID(),
		Address:    tk.ToPk(),
		Tk:         tk,
		PrivateKey: priKey,
		At:         0,
		Version:    1,
	}
}*/

func IDAndNonce(srcId types.ChainId, nonce types.Nonce) *big.Int {
	var data []byte
	data = append(data, nonce.Big().Bytes()...)
	data = append(data, uint8(srcId))
	return big.NewInt(0).SetBytes(data)
}
