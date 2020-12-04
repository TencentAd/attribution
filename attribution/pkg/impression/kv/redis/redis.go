package redis

import (
	"context"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/impression/kv/opt"
	"github.com/go-redis/redis/v8"
)

var (
	DefaultClientTimeOut = 30 * time.Second
)

type Redis struct {
	client *redis.Client
	option *opt.Option
}

func (r *Redis) Get(key string) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), DefaultClientTimeOut)
	return r.client.Get(ctx, opt.Prefix+key).Result()
}

func (r *Redis) Set(key string, value string) error {
	ctx, _ := context.WithTimeout(context.Background(), DefaultClientTimeOut)
	return r.client.Set(ctx, opt.Prefix+key, value, r.option.Expiration).Err()
}

func New(option *opt.Option) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     option.Address,
		Password: option.Password,
	})

	return &Redis{
		client: client,
		option: option,
	}, nil
}
