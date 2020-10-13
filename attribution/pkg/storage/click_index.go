/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 4:26 PM
 */

package storage

import (
	"flag"
	"fmt"

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/proto/click"
	"github.com/TencentAd/attribution/attribution/proto/user"
)

var (
	clickIndexName = flag.String("click_index_name", "native", "support native|redis|hbase")

	ClickIndexFactory = factory.NewFactory()
)

type ClickIndex interface {
	Set(idType user.IdType, key string, click *click.ClickLog) error
	Get(idtype user.IdType, key string) (*click.ClickLog, error)
	Remove(idType user.IdType, key string) error
}

func CreateClickIndex() (ClickIndex, error) {
	clickIndex := ClickIndexFactory.Create(*clickIndexName)
	if clickIndex == nil {
		return nil, fmt.Errorf("click index[%s] not support", *clickIndexName)
	}
	return clickIndex.(ClickIndex), nil
}
