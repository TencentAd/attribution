/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/5/20, 4:44 PM
 */

package redis

import (
	"flag"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/crypto/freq/info"
	"github.com/go-redis/redis"
	"github.com/golang/glog"
)

var (
	minuteFreqRedisAddress        = flag.String("minute_freq_redis_address", "", "")
	minuteFreqRedisServerPassword = flag.String("minute_freq_redis_password", "", "minuteFreqRedis server password")
	minuteFreqRedisWriteTimeout   = flag.Int("minute_freq_redis_write_timeout_ms", 50, "minuteFreqRedis write timeout")
	minuteFreqRedisReadTimeout    = flag.Int64("minute_freq_redis_read_timeout_ms", 50, "minuteFreqRedis read timeout")
	minuteFreqRedisIsCluster      = flag.Bool("minute_freq_redis_is_cluster", false, "")
	minuteFreqRedisExpiration     = flag.Int("minute_freq_redis_expiration_second", 3600, "")
)

const (
	minuteFreqPrefix = "minute_freq::"
)

type MinuteFreqRedis struct {
	redisClient redis.Cmdable
}

func NewMinuteFreqRedis() interface{} {
	redisClient, err := redisx.CreateRedisClient(&redisx.Option{
		ClusterOptions: &redis.ClusterOptions{
			Addrs:        []string{*minuteFreqRedisAddress},
			Password:     *minuteFreqRedisServerPassword,
			WriteTimeout: time.Millisecond * time.Duration(*minuteFreqRedisWriteTimeout),
			ReadTimeout:  time.Millisecond * time.Duration(*minuteFreqRedisReadTimeout),
		},
		Options: &redis.Options{
			Addr:         *minuteFreqRedisAddress,
			Password:     *minuteFreqRedisServerPassword,
			WriteTimeout: time.Millisecond * time.Duration(*minuteFreqRedisWriteTimeout),
			ReadTimeout:  time.Millisecond * time.Duration(*minuteFreqRedisReadTimeout),
		},
		IsCluster: *minuteFreqRedisIsCluster,
	})

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
	_, err = c.redisClient.Set(
		c.formatKey(resource),
		str,
		time.Duration(*minuteFreqRedisExpiration)*time.Second).Result()
	return err
}

func (c *MinuteFreqRedis) Get(resource string) (*info.MinuteFreqInfo, error) {
	val, err := c.redisClient.Get(c.formatKey(resource)).Result()

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
