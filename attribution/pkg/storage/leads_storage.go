/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 5:05 PM
 */

package storage

import (
	"flag"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/leads/pull/protocal"
	"github.com/TencentAd/attribution/attribution/pkg/storage/redis"
)

var (
	leadsStorageName    = flag.String("leads_storage_name", "redis", "")
	leadsStorageFactory = factory.NewFactory("leads_storage")
)

func init() {
	leadsStorageFactory.Register("redis", redis.NewLeadsRedisStorage)
}

// 线索存储
type LeadsStorage interface {
	Store(leads *protocal.LeadsInfo) error
}

func CreateLeadsStorage() (LeadsStorage, error) {
	if s, err := leadsStorageFactory.Create(*leadsStorageName); err != nil {
		return nil, err
	} else {
		return s.(LeadsStorage), nil
	}
}
