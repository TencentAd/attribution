/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/23/20, 11:39 AM
 */

package util

import "math/big"

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
		// prime是一个素数
		Prime:  prime,
		// 这是生成的那个key
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

func ModPow(data *big.Int, e *big.Int, prime *big.Int) *big.Int {
	return big.NewInt(0).Exp(data, e, prime)
}

func (e *ModPower) findDecKey() *big.Int {
	return big.NewInt(0).ModInverse(e.EncKey, big.NewInt(0).Sub(e.Prime, big.NewInt(1)))
}

func FindDecKey(e *big.Int, p *big.Int) *big.Int {
	return big.NewInt(0).ModInverse(e, big.NewInt(0).Sub(p, big.NewInt(1)))
}