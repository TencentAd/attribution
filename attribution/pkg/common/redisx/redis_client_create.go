package redisx

import (
	"github.com/go-redis/redis"
)

type Option struct {
	ClusterOptions *redis.ClusterOptions
	Options        *redis.Options
	IsCluster      bool
}

func CreateRedisClient(option *Option) (redis.Cmdable, error) {
	var ret redis.Cmdable
	if option.IsCluster {
		ret = redis.NewClusterClient(option.ClusterOptions)
	} else {
		ret = redis.NewClient(option.Options)
	}

	_, err := ret.Ping().Result()
	return ret, err
}
