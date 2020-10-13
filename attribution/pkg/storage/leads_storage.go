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
	"fmt"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/leads/pull/protocal"
)

var (
	leadsStorageName    = flag.String("leads_storage_name", "redis", "")
	LeadsStorageFactory = factory.NewFactory()
)

// 线索存储
type LeadsStorage interface {
	Store(leads *protocal.LeadsInfo) error
}

func CreateLeadsStorage() (LeadsStorage, error) {
	s := LeadsStorageFactory.Create(*leadsStorageName)
	if s == nil {
		return nil, fmt.Errorf("leads storage[%s] not support", *leadsStorageName)
	}
	return s.(LeadsStorage), nil
}
