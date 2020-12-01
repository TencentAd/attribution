package redisx

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/define"
	"github.com/go-redis/redis/v8"
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

	_, err := ret.Ping(context.Background()).Result()
	return ret, err
}

type ShortOption struct {
	Address        []string `json:"address"`
	Password       string   `json:"password"`
	WriteTimeoutMS int      `json:"write_timeout"`
	ReadTimeMS     int      `json:"read_timeout"`
	IsCluster      int      `json:"is_cluster"`
}

func CreateRedisClientV2(jsonConfig string) (redis.Cmdable, error) {
	var so ShortOption

	var config string
	if jsonConfig == "" || jsonConfig == "{}" {
		config = *define.DefaultRedisConfig
	}

	if err := json.Unmarshal([]byte(config), &so); err != nil {
		return nil, err
	}
	return CreateRedisClientV3(&so)
}

func CreateRedisClientV3(so *ShortOption) (redis.Cmdable, error) {
	if len(so.Address) == 0 {
		return nil, fmt.Errorf("no address configured")
	}

	opt := shortOption2Option(so)

	return CreateRedisClient(opt)
}

func shortOption2Option(so *ShortOption) *Option {
	return &Option{
		IsCluster: so.IsCluster != 0,
		ClusterOptions: &redis.ClusterOptions{
			Addrs:        so.Address,
			Password:     so.Password,
			WriteTimeout: time.Millisecond * time.Duration(so.WriteTimeoutMS),
			ReadTimeout:  time.Millisecond * time.Duration(so.ReadTimeMS),
		},
		Options: &redis.Options{
			Addr:         so.Address[0],
			Password:     so.Password,
			WriteTimeout: time.Millisecond * time.Duration(so.WriteTimeoutMS),
			ReadTimeout:  time.Millisecond * time.Duration(so.ReadTimeMS),
		},
	}
}
