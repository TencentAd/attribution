package redis

import (
    "context"
    "time"

    "github.com/TencentAd/attribution/attribution/pkg/impression/kv"
    "github.com/go-redis/redis/v8"
)

var (
    DefaultClientTimeOut = 30*time.Second
)

type Redis struct {
    client *redis.Client
    option *kv.Option
}

func (r *Redis) Has(key string) (bool, error) {
    ctx, _ := context.WithTimeout(context.Background(), DefaultClientTimeOut)
    _, err := r.client.Get(ctx, kv.Prefix + key).Result()
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
    return r.client.Set(ctx, kv.Prefix + key, nil, r.option.Expiration).Err()
}

func New(option *kv.Option) (*Redis, error) {
    client := redis.NewClient(&redis.Options{
        Addr:     option.Address,
        Password: option.Password,
    })

    return &Redis{
        client: client,
        option: option,
    }, nil
}