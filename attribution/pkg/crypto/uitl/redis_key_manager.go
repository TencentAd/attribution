package uitl

import (
	"context"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/conf"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/storage"
	"github.com/bsm/redislock"
	"github.com/golang/glog"
	"math/big"
	"time"
)

// 这个类只负责从生成、获取、存储groupId对应的秘钥（随机数）
type RedisKeyManager struct {
	redisStorage *storage.RedisStorage
	mutex        *redislock.Client
}

func NewRedisKeyManager(redisConf *conf.RedisConf) *RedisKeyManager {
	redisStorage := storage.NewRedisStorage(redisConf)
	lock := redisStorage.GetLock()

	return &RedisKeyManager{
		redisStorage: redisStorage,
		mutex:        lock,
	}
}

// 这是一个接口，接受用户的cid，在同一批的服务中，cid是相同的，所以我需要生成一个key，然后回传
// 第一次请求这个接口，数据库中是没有这秘钥的，所以需要生成，然后后面的服务都能够获取到这样的一个cid
// 生成的这个key其实和cid无关，但是同一个cid一定要获取到相同的key
func (keyManager *RedisKeyManager) GetEncryptKey(groupId string) (*big.Int, error) {
	// 首先查看是否在内存中存在
	value, ok := storage.Fetch(groupId)
	if ok {
		result := new(big.Int)
		result, _ = result.SetString(value, 16)
		return result, nil
	}
	// 否则去redis服务器中拿
	// 这里开始加锁
	ctx := context.Background()
	lock, err := keyManager.mutex.Obtain(ctx, "lock", 100*time.Millisecond, nil)
	if err == redislock.ErrNotObtained {
		glog.Errorf("could not obtain lock")
		//return nil
		return nil, err
	} else if err != nil {
		glog.Errorf("fail to add a mutex, err[%v]", err)
		//return nil, err
		return nil, err
	}

	// 延后执行，在方法return之后就会自动解锁
	defer lock.Release(ctx)

	value, err = keyManager.redisStorage.Fetch(groupId)
	if value != "" {
		// 说明取到了值
		result := new(big.Int)
		result, _ = result.SetString(value, 16)
		// 这列应该存进内存
		storage.Storage(groupId, value)
		//return result, nil
		return result, nil
	} else if err != nil {
		// 如果出现了错误
		return nil, err
	}

	// 内存中和redis都没有对应的秘钥，说明需要生成一个
	encKey, err := GenerateEncKey(224)
	if err != nil {
		glog.Errorf("fail to generate encrypt key, err[%v]", err)
		//return nil, errors.New("fail to generate encrypt key")
		return nil, err
	}
	// 生成秘钥成功
	// 存入redis
	err = keyManager.redisStorage.Storage(groupId, encKey.Text(16))
	if err != nil {
		glog.Errorf("fail to storage encrypt key into redis, err[%v]", err)
		//return nil, errors.New("fail to storage encrypt key into redis")
		return nil, err
	}
	// 到这里锁就能够解开了

	// 存入内存
	storage.Storage(groupId, encKey.Text(16))
	// 存入redis

	return encKey, nil
}
