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

	"github.com/TencentAd/attribution/attribution/pkg/common/factory"
	"github.com/TencentAd/attribution/attribution/pkg/storage/hbase"
	"github.com/TencentAd/attribution/attribution/pkg/storage/native"
	"github.com/TencentAd/attribution/attribution/proto/click"
	"github.com/TencentAd/attribution/attribution/proto/user"
)

var (
	clickIndexName = flag.String("click_index_name", "native", "support native|redis|hbase")

	clickIndexFactory = factory.NewFactory("click_index")
)

func init() {
	clickIndexFactory.Register("hbase", hbase.NewClickIndexHbase)
	clickIndexFactory.Register("native", native.NewClickIndexNative)
}

type ClickIndex interface {
	Set(idType user.IdType, key string, click *click.ClickLog) error
	Get(idtype user.IdType, key string) (*click.ClickLog, error)
	Remove(idType user.IdType, key string) error
}

func CreateClickIndex() (ClickIndex, error) {
	if clickIndex, err := clickIndexFactory.Create(*clickIndexName); err != nil {
		return nil, err
	} else {
		return clickIndex.(ClickIndex), nil
	}
}
