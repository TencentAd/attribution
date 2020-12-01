/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 10:22 AM
 */

package redis

import (
	"context"
	"flag"
	"strconv"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

var (
	counterRedisConfig = flag.String("counter_redis_config", "", "")
)

const (
	redisCounterPrefix = "counter::"
)

type Counter struct {
	redisClient redis.Cmdable
}

func NewCounter() interface{} {
	redisClient, err := redisx.CreateRedisClientV2(*counterRedisConfig)

	if err != nil {
		glog.Errorf("failed to create counter redis, err: %v", err)
		return nil
	}

	return &Counter{
		redisClient: redisClient,
	}
}

func (c *Counter) Inc(resource string) (int64, error) {
	 return c.redisClient.Incr(context.Background(), c.formatKey(resource)).Result()
}

func (c *Counter) Get(resource string) (int64, error) {
	val, err := c.redisClient.Get(context.Background(), c.formatKey(resource)).Result()
	if err != nil {
		return -1, err
	} else {
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return -1, err
		} else {
			return intVal, nil
		}
	}
}

func (c *Counter) formatKey(resource string) string {
	return redisCounterPrefix + resource
}
