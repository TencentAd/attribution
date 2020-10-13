/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 5:08 PM
 */

package redis

import (
	"encoding/json"
	"flag"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx"
	"github.com/TencentAd/attribution/attribution/pkg/leads/pull/protocal"

	"github.com/go-redis/redis"
)

var (
	leadsRedisAddress        = flag.String("leads_redis_address", "", "")
	leadsRedisServerPassword = flag.String("leads_redis_password", "", "leadsRedis server password")
	leadsRedisWriteTimeout   = flag.Int("leads_redis_write_timeout_ms", 50, "leadsRedis write timeout")
	leadsRedisReadTimeout    = flag.Int64("leads_redis_read_timeout_ms", 50, "leadsRedis read timeout")
	leadsRedisIsCluster      = flag.Bool("leads_redis_is_cluster", false, "")
	leadsRedisExpirationHour = flag.Int("leads_redis_expiration_hour", 720, "")
)

type LeadsRedisStorage struct {
	redisClient redis.Cmdable
}

func NewLeadsRedisStorage() *LeadsRedisStorage {
	return &LeadsRedisStorage{}
}

func (s *LeadsRedisStorage) Init() error {
	redisClient, err := redisx.CreateRedisClient(&redisx.Option{
		ClusterOptions: &redis.ClusterOptions{
			Addrs:        []string{*leadsRedisAddress},
			Password:     *leadsRedisServerPassword,
			WriteTimeout: time.Millisecond * time.Duration(*leadsRedisWriteTimeout),
			ReadTimeout:  time.Millisecond * time.Duration(*leadsRedisReadTimeout),
		},
		Options: &redis.Options{
			Addr:         *leadsRedisAddress,
			Password:     *leadsRedisServerPassword,
			WriteTimeout: time.Millisecond * time.Duration(*leadsRedisWriteTimeout),
			ReadTimeout:  time.Millisecond * time.Duration(*leadsRedisReadTimeout),
		},
		IsCluster: *leadsRedisIsCluster,
	})
	if err != nil {
		return err
	}
	s.redisClient = redisClient
	return nil
}

func (s *LeadsRedisStorage) Store(leads *protocal.LeadsInfo) error {
	data, err := json.Marshal(leads)
	if err != nil {
		return err
	}
	_, err = s.redisClient.Set(
		s.FormatKey(leads.LeadsTelephone),
		data,
		time.Hour*time.Duration(*leadsRedisExpirationHour)).Result()
	return err
}

const prefix = "leads::"

func (s *LeadsRedisStorage) FormatKey(key string) string {
	return prefix + key
}
