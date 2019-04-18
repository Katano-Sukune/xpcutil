// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package xpcutil

const (
	// SatoshiPerBitcent is the number of satoshi in one litecoin cent.
	SatoshiPerBitcent = 1e2

	// SatoshiPerBitcoin is the number of satoshi in one litecoin (1 BTC).
	SatoshiPerBitcoin = 1e4

	// MaxSatoshi is the maximum transaction amount allowed in satoshi.
	MaxSatoshi = 21e10 * SatoshiPerBitcoin
)
