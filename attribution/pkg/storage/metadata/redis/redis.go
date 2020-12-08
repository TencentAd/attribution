package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	Prefix = "metadata::"
	DefaultClientTimeOut = 30 * time.Second
)

type Redis struct {
	client *redis.Client
}

func (r *Redis) Get(key string) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), DefaultClientTimeOut)
	v, err := r.client.Get(ctx, Prefix+key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return v, err
}

func (r *Redis) Set(key string, value string) error {
	ctx, _ := context.WithTimeout(context.Background(), DefaultClientTimeOut)
	return r.client.Set(ctx, Prefix+key, value, redis.KeepTTL).Err()
}

func New(config map[string]interface{}) (*Redis, error) {
	option := &redis.Options{}
	b, _ := json.Marshal(config)
	if err := json.Unmarshal(b, option); err != nil {
		return nil, err
	}

	client := redis.NewClient(option)
	return &Redis{client: client}, nil
}
