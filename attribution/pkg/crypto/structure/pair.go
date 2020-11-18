/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/12/20, 10:29 AM
 */

package structure

import "math/big"

type CryptoPair struct {
	EncKey *big.Int
	DecKey *big.Int
}
