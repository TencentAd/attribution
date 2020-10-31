package redis

import (
    "context"
    "time"

    "github.com/TencentAd/attribution/attribution/pkg/impression/kv/opt"
    "github.com/go-redis/redis/v8"
)

var (
    DefaultClientTimeOut = 30*time.Second
)

type Redis struct {
    client *redis.Client
    option *opt.Option
}

func (r *Redis) Has(key string) (bool, error) {
    ctx, _ := context.WithTimeout(context.Background(), DefaultClientTimeOut)
    _, err := r.client.Get(ctx, opt.Prefix + key).Result()
    if err == redis.Nil {
        return false, nil
    }

    if err != nil {
        return false, err
    }

    return true, nil
}

func (r *Redis) Set(key string) error {
    ctx, _ := context.WithTimeout(context.Background(), DefaultClientTimeOut)
    return r.client.Set(ctx, opt.Prefix + key, nil, r.option.Expiration).Err()
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