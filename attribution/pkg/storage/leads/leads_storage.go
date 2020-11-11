/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/10/20, 5:16 PM
 */

package leads

import (
	"flag"
	"time"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/leads/pull/protocal"
	"github.com/TencentAd/attribution/attribution/pkg/storage/leads/redis"
)

var (
	leadsStorageName    = flag.String("leads_storage_name", "redis", "support [redis]")
	leadsStorageFactory = factory.NewFactory("leads_storage")
)

func init() {
	leadsStorageFactory.Register("redis", redis.NewLeadsRedis)
}

// 线索存储
type Storage interface {
	Store(leads *protocal.LeadsInfo, expire time.Duration) error
}

func CreateLeadsStorage() (Storage, error) {
	if s, err := leadsStorageFactory.Create(*leadsStorageName); err != nil {
		return nil, err
	} else {
		return s.(Storage), nil
	}
}
