/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/12/20, 10:08 AM
 */

package keymanager

import (
	"flag"
	"math/big"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/keymanager/hdfs"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/keymanager/redis"
)

var (
	keyManagerName = flag.String("key_manager_name", "hdfs", "support [redis, hdfs]")

	keyManagerFactory = factory.NewFactory("key_manager")
)

func init() {
	keyManagerFactory.Register("redis", redis.NewRedisKeyManager)
	keyManagerFactory.Register("hdfs", hdfs.NewHDFSKeyManager)
}

// 接口，实现类有redis的KeyManager和hdfs的KeyManager
type KeyManager interface {
	GetEncryptKey(string) (*big.Int, error)
	GetDecryptKey(string) (*big.Int, error)
}

func CreateKeyManager() (KeyManager, error) {
	if s, err := keyManagerFactory.Create(*keyManagerName); err != nil {
		return nil, err
	} else {
		return s.(KeyManager), nil
	}
}
