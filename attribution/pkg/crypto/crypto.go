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
	"sync"

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
	if len(data) == 0 {
		return "", nil
	}

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
	if len(data) == 0 {
		return "", nil
	}

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

type Parallel struct {
	wg sync.WaitGroup
	err error
}

func (p *Parallel) AddTask(fun func(string, string) (string, error), groupId string, data string, result *string) {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		if p.err != nil {
			return
		}
		var err error
		*result, err = fun(groupId, data)
		if err != nil {
			p.err = err
		}
	}()
}

func (p *Parallel) WaitAndCheck() error {
	p.wg.Wait()
	return p.err
}
