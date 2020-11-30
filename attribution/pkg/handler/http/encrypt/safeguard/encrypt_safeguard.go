/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/3/20, 3:15 PM
 */

package safeguard

import (
	"context"
	"errors"
	"flag"
	"strconv"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

var (
	minuteFreqRedisConfig    = flag.String("minute_freq_redis_config", "{}", "")
	convEncryptMaxMinuteFreq = flag.Int("conv_encrypt_max_minute_freq", 10, "")

	ErrConvEncryptSafeguardInternal = errors.New("conv encrypt safeguard internal error")
	ErrConvEncryptExceedFreq        = errors.New("conv encrypt exceed frequency")
)

type ConvEncryptSafeguard struct {
	redisClient redis.Cmdable
}

func NewConvEncryptSafeguard() (*ConvEncryptSafeguard, error) {
	client, err := redisx.CreateRedisClientV2(*minuteFreqRedisConfig)
	if err != nil {
		return nil, err
	}
	return &ConvEncryptSafeguard{
		redisClient: client,
	}, nil
}

func (g *ConvEncryptSafeguard) Against(campaignId int64) error {
	key := g.formatResourceKey(campaignId)

	cnt, err := g.redisClient.Incr(context.Background(), key).Result()
	if err != nil {
		glog.Errorf("failed to incr key, err: %v", err)
		return ErrConvEncryptSafeguardInternal
	}

	if cnt == 1 {
		if _, err := g.redisClient.Expire(context.Background(), key, time.Second*70).Result(); err != nil {
			glog.Errorf("failed to expire key [%s]", key)
		}
	}

	if int(cnt) > *convEncryptMaxMinuteFreq {
		return ErrConvEncryptExceedFreq
	}

	return nil
}

const ConvEncryptResourcePrefix = "convEncrypt::"

func (g *ConvEncryptSafeguard) formatResourceKey(campaignId int64) string {
	return ConvEncryptResourcePrefix + strconv.FormatInt(campaignId, 10) + "_" +
		strconv.FormatInt(time.Now().Unix()/60, 10)
}
