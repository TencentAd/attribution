package redis

import (
	"context"
	"flag"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

var (
	keyStorageRedisConfig = flag.String("key_storage_redis_config", "{}", "")
)

type KeyStorageRedis struct {
	client redis.Cmdable
}

// 该方法用于构造一个新的RedisStorage并返回
func NewKeyStorageRedis() (*KeyStorageRedis, error) {
	client, err := redisx.CreateRedisClientV2(*keyStorageRedisConfig)
	if err != nil {
		return nil, err
	}

	return &KeyStorageRedis{
		client: client,
	}, nil
}

func (r *KeyStorageRedis) Store(groupId string, encryptKey string) error {
	_, err := r.client.Set(context.Background(), r.formatKey(groupId), encryptKey, 0).Result()
	if err != nil {
		glog.Errorf("fail to set key, err[%v]", err)
		return err
	}
	return nil
}

func (r *KeyStorageRedis) Load(groupId string) (string, error) {
	value, err := r.client.Get(context.Background(), r.formatKey(groupId)).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		glog.Errorf("fail to get encrypt key, err[%v]", err)
		return "", err
	}
	return value, nil
}

const keyPrefix = "key::"
func (r *KeyStorageRedis) formatKey(key string) string {
	return keyPrefix + key
}
