/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/23/20, 11:39 AM
 */

package modpower

import (
	"fmt"
	"math/big"
)

type ModPower struct {
	Prime  *big.Int
	EncKey *big.Int
	DecKey *big.Int
}

// Requirement
// 1. (prime - 1) / 2 must also be prime
// 2. encKey must be an odd number
func NewModPower(prime *big.Int, encKey *big.Int) *ModPower {
	e := &ModPower{
		Prime:  prime,
		EncKey: encKey,
	}
	e.DecKey = e.findDecKey()
	return e
}

func (e *ModPower) Encrypt(data *big.Int) *big.Int {
	return big.NewInt(0).Exp(data, e.EncKey, e.Prime)
}

func (e *ModPower) Decrypt(data *big.Int) *big.Int {
	return big.NewInt(0).Exp(data, e.DecKey, e.Prime)
}

func (e *ModPower) findDecKey() *big.Int {
	return big.NewInt(0).ModInverse(e.EncKey, big.NewInt(0).Sub(e.Prime, big.NewInt(1)))
}

func Hex2BigInt(hex string) (*big.Int, error) {
	data, ok := big.NewInt(0).SetString(hex, 16)
	if !ok {
		return nil, fmt.Errorf("not valid hex")
	} else {
		return data, nil
	}
}

func MustHex2BigInt(hex string) *big.Int {
	data, _ := big.NewInt(0).SetString(hex, 16)
	return data
}

func Encrypt(data *big.Int, encKey *big.Int, prime *big.Int) *big.Int {
	return big.NewInt(0).Exp(data, encKey, prime)
}
