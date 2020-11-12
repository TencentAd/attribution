/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/12/20, 9:44 PM
 */

package crypto

import (
	"math/big"

	"github.com/TencentAd/attribution/attribution/pkg/common/define"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/keymanager"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/util"
	"github.com/golang/glog"
)

var (
	km keymanager.KeyManager
)

func InitCrypto() error {
	var err error
	km, err = keymanager.CreateKeyManager()
	if err != nil {
		glog.Errorf("failed to init key manager, err: %v", err)
		return err
	}
	return nil
}

func Encrypt(groupId string, data string) (string, error) {
	var err error
	var encKey *big.Int
	encKey, err = km.GetEncryptKey(groupId)
	if err != nil {
		return "", err
	}

	var x *big.Int
	x, err = util.Hex2BigInt(data)
	if err != nil {
		return "", err
	}

	return util.ModPow(x, encKey, define.Prime).Text(16), nil
}

func Decrypt(groupId string, data string) (string, error) {
	var err error
	var decKey *big.Int
	decKey, err = km.GetDecryptKey(groupId)
	if err != nil {
		return "", err
	}

	var x *big.Int
	x, err = util.Hex2BigInt(data)
	if err != nil {
		return "", err
	}

	return util.ModPow(x, decKey, define.Prime).Text(16), nil
}
