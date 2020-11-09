/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 11:14 AM
 */

package safeguard

import "github.com/TencentAd/attribution/attribution/pkg/handler/http/crypto/protocal"

type Parameter struct {
	CryptoRequest  *protocal.CryptoRequest
}

func NewParameter() *Parameter {
	return &Parameter{}
}

func (opt *Parameter) WithCryptoRequest(req *protocal.CryptoRequest) *Parameter {
	opt.CryptoRequest=  req
	return opt
}
