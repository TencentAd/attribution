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
