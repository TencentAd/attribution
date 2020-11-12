package redis

import (
	"context"
	"math/big"

	"github.com/TencentAd/attribution/attribution/pkg/crypto/keymanager/cache"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/structure"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/util"
	"github.com/bsm/redislock"
	"github.com/golang/glog"
)

// 这个类只负责从生成、获取、存储groupId对应的秘钥（随机数）
type KeyManager struct {
	cache        *cache.CryptoCache
	mutex        *Lock
	redisStorage *KeyStorageRedis
}

func NewRedisKeyManager() interface{} {
	redisStorage, err := NewKeyStorageRedis()
	if err != nil {
		glog.Errorf("failed to create redis key storage, err: %v", err)
		return nil
	}

	lock, err := NewRedisLock()
	if err != nil {
		glog.Errorf("failed to create redis lock, err: %v", err)
	}

	return &KeyManager{
		redisStorage: redisStorage,
		mutex:        lock,
		cache:        cache.NewCryptoCache(),
	}
}

func (keyManager *KeyManager) GetEncryptKey(groupId string) (*big.Int, error) {
	pair, err := keyManager.GetCryptoPair(groupId)
	if err != nil {
		return nil, err
	}

	return pair.EncKey, nil
}

func (keyManager *KeyManager) GetDecryptKey(groupId string) (*big.Int, error) {
	pair, err := keyManager.GetCryptoPair(groupId)
	if err != nil {
		return nil, err
	}

	return pair.DecKey, nil
}

// 这是一个接口，接受用户的cid，在同一批的服务中，cid是相同的，所以我需要生成一个key，然后回传
// 第一次请求这个接口，数据库中是没有这秘钥的，所以需要生成，然后后面的服务都能够获取到这样的一个cid
// 生成的这个key其实和cid无关，但是同一个cid一定要获取到相同的key
func (keyManager *KeyManager) GetCryptoPair(groupId string) (*structure.CryptoPair, error) {
	// 首先查看是否在内存中存在
	pair, ok := keyManager.cache.Load(groupId)
	if ok {
		return pair, nil
	}
	// 否则去redis服务器中拿
	// 这里开始加锁
	lock, err := keyManager.mutex.Obtain(groupId)
	if err == redislock.ErrNotObtained {
		glog.Errorf("could not obtain lock")
		return nil, err
	} else if err != nil {
		glog.Errorf("fail to add a mutex, err[%v]", err)
		return nil, err
	}

	// 延后执行，在方法return之后就会自动解锁
	defer func() {
		if err := lock.Release(context.Background()); err != nil {
			glog.Errorf("failed to release lock for groupId [%s]", groupId)
		}
	}()

	var rVal string
	rVal, err = keyManager.redisStorage.Load(groupId)
	if err != nil {
		return nil, err
	}
	if rVal != "" {
		encKey, err := util.Hex2BigInt(rVal)
		if err != nil {
			glog.Errorf("group[%s] encKey[%s] not valid integer", groupId, rVal)
			return nil, err
		}
		pair = util.CreateCryptoPair(encKey)
		keyManager.cache.Store(groupId, pair)
		return pair, nil
	}

	// 内存中和redis都没有对应的秘钥，说明需要生成一个
	pair, err = util.GenerateCryptoPair(0)
	if err != nil {
		glog.Errorf("fail to generate encrypt key, err[%v]", err)
		return nil, err
	}
	// 生成秘钥成功
	// 存入redis
	err = keyManager.redisStorage.Store(groupId, util.BigInt2Hex(pair.EncKey))
	if err != nil {
		glog.Errorf("fail to storage encrypt key into redis, err[%v]", err)
		//return nil, errors.New("fail to storage encrypt key into redis")
		return nil, err
	}
	// 到这里锁就能够解开了

	// 存入内存
	keyManager.cache.Store(groupId, pair)
	// 存入redis

	return pair, nil
}
