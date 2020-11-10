package storage

import (
	"context"
	"github.com/TencentAd/attribution/attribution/pkg/crypto/conf"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
	"time"
)

type RedisStorage struct {
	client    redislock.RedisClient
	redisMode conf.RedisMode
}

// 该方法用于构造一个新的RedisStorage并返回
func NewRedisStorage(redisConf *conf.RedisConf) *RedisStorage {
	var client redislock.RedisClient
	if redisConf.RedisMode == conf.RedisSingleNode {
		client = redis.NewClient(&redis.Options{
			Addr:         redisConf.NodeAddresses[0],
			DialTimeout:  10 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		})
	}
	if redisConf.RedisMode == conf.RedisCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        redisConf.NodeAddresses,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		})
	}
	return &RedisStorage{
		client:    client,
		redisMode: redisConf.RedisMode,
	}
}

func (r *RedisStorage) GetLock() *redislock.Client {
	return redislock.New(r.client)
}

func (r *RedisStorage) Storage(groupId string, encryptKey string) error {
	ctx := context.Background()
	_, err := r.client.(redis.Cmdable).Set(ctx, groupId, encryptKey, 0).Result()
	if err != nil {
		glog.Errorf("fail to set key, err[%v]", err)
		return err
	}
	return nil
}

func (r *RedisStorage) Fetch(groupId string) (string, error) {
	ctx := context.Background()
	value, err := r.client.(redis.Cmdable).Get(ctx, groupId).Result()
	if err != nil {
		glog.Errorf("fail to get encrypt key, err[%v]", err)
		return "", err
	}
	return value, nil
}
