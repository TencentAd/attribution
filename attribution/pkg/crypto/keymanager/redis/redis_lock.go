/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/12/20, 10:43 AM
 */

package redis

import (
	"context"
	"flag"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/redisx/v8"
	"github.com/bsm/redislock"
)

var (
	keyRedisLockConfig = flag.String("key_redis_lock_config", "{}", "")
	lockTTLMS          = flag.Int("key_redis_lock_ttl_ms", 10000, "")
)

type Lock struct {
	redisLock *redislock.Client
}

func NewRedisLock() (*Lock, error) {
	client, err := redisx.CreateRedisClientV2(*keyRedisLockConfig)
	if err != nil {
		return nil, err
	}

	return &Lock{
		redisLock: redislock.New(client),
	}, nil
}

func (lock *Lock) Obtain(key string) (*redislock.Lock, error) {
	return lock.redisLock.Obtain(
		context.Background(),
		lock.formatKey(key),
		time.Millisecond*time.Duration(*lockTTLMS),
		&redislock.Options{
			RetryStrategy: redislock.LimitRetry(redislock.LinearBackoff(100*time.Millisecond), 10),
		})
}

const keyLockPrefix = "keyLock::"

func (lock *Lock) formatKey(key string) string {
	return keyLockPrefix + key
}
