/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/11/20, 2:39 PM
 */

package redis

import (
	"context"
	"flag"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/TencentAd/attribution/attribution/pkg/leads/pull/protocal"
	"github.com/TencentAd/attribution/attribution/pkg/storage/leads/util"
	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

var (
	leadsRedisConfig        = flag.String("leads_redis_config", "{}", "")
)

type LeadsRedis struct {
	redisClient redis.Cmdable
}

func NewLeadsRedis() interface{} {
	redisClient, err := redisx.CreateRedisClientV2(*leadsRedisConfig)

	if err != nil {
		glog.Errorf("failed to create counter redis, err: %v", err)
		return nil
	}

	return &LeadsRedis{
		redisClient: redisClient,
	}
}

func (lr *LeadsRedis) Store(leads *protocal.LeadsInfo, expire time.Duration) error {
	indexes, err := util.ExtractLeadsIndex(leads)
	if err != nil {
		return err
	}

	for _, index := range indexes {
		_, err = lr.redisClient.Set(context.Background(), index.Key, index.Value, expire).Result()
		if err != nil {
			return err
		}
	}

	return nil
}
