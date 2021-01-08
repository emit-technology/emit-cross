// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/sero-cash/go-sero/crypto"
)

func Hash(data []byte) (ret [32]byte) {
	h := crypto.Keccak256Hash(data)
	copy(ret[:], h[:])
	return
}
