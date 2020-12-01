/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 4:44 PM
 */

package redis

import (
	"context"
	"flag"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/encrypt/freq/info"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

var (
	minuteFreqRedisConfig        = flag.String("minute_freq_redis_config", "{}", "")
	minuteFreqRedisExpiration     = flag.Int("minute_freq_redis_expiration_second", 3600, "")
)

const (
	minuteFreqPrefix = "minute_freq::"
)

type MinuteFreqRedis struct {
	redisClient redis.Cmdable
}

func NewMinuteFreqRedis() interface{} {
	redisClient, err := redisx.CreateRedisClientV2(*minuteFreqRedisConfig)

	if err != nil {
		glog.Errorf("failed to create counter redis, err: %v", err)
		return nil
	}

	return &MinuteFreqRedis{
		redisClient: redisClient,
	}
}

func (c *MinuteFreqRedis) Set(resource string, info *info.MinuteFreqInfo) error {
	str, err := info.String()
	if err != nil {
		return err
	}
	_, err = c.redisClient.Set(context.Background(),
		c.formatKey(resource),
		str,
		time.Duration(*minuteFreqRedisExpiration)*time.Second).Result()
	return err
}

func (c *MinuteFreqRedis) Get(resource string) (*info.MinuteFreqInfo, error) {
	val, err := c.redisClient.Get(context.Background(), c.formatKey(resource)).Result()

	if err == redis.Nil {
		return &info.MinuteFreqInfo{
			LastMinute: 0,
			Count:      0,
		}, nil
	}

	if err != nil {
		return nil, err
	}

	return info.ParseMinuteFreqInfo([]byte(val))
}

func (c *MinuteFreqRedis) formatKey(resource string) string {
	return minuteFreqPrefix + resource
}
