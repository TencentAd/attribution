/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 10:22 AM
 */

package redis

import (
	"flag"
	"strconv"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
)

var (
	counterRedisAddress        = flag.String("counter_redis_address", "", "")
	counterRedisServerPassword = flag.String("counter_redis_password", "", "counterRedis server password")
	counterRedisWriteTimeout   = flag.Int("counter_redis_write_timeout_ms", 50, "counterRedis write timeout")
	counterRedisReadTimeout    = flag.Int64("counter_redis_read_timeout_ms", 50, "counterRedis read timeout")
	counterRedisIsCluster      = flag.Bool("counter_redis_is_cluster", false, "")
)

const (
	redisCounterPrefix = "counter::"
)

type Counter struct {
	redisClient redis.Cmdable
}

func NewCounter() interface{} {
	redisClient, err := redisx.CreateRedisClient(&redisx.Option{
		ClusterOptions: &redis.ClusterOptions{
			Addrs:        []string{*counterRedisAddress},
			Password:     *counterRedisServerPassword,
			WriteTimeout: time.Millisecond * time.Duration(*counterRedisWriteTimeout),
			ReadTimeout:  time.Millisecond * time.Duration(*counterRedisReadTimeout),
		},
		Options: &redis.Options{
			Addr:         *counterRedisAddress,
			Password:     *counterRedisServerPassword,
			WriteTimeout: time.Millisecond * time.Duration(*counterRedisWriteTimeout),
			ReadTimeout:  time.Millisecond * time.Duration(*counterRedisReadTimeout),
		},
		IsCluster: *counterRedisIsCluster,
	})

	if err != nil {
		glog.Errorf("failed to create counter redis, err: %v", err)
		return nil
	}

	return &Counter{
		redisClient: redisClient,
	}
}

func (c *Counter) Inc(resource string) (int64, error) {
	 return c.redisClient.Incr(c.formatKey(resource)).Result()
}

func (c *Counter) Get(resource string) (int64, error) {
	val, err := c.redisClient.Get(c.formatKey(resource)).Result()
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
