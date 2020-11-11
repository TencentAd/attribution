/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 7:21 PM
 */

package main

import (
	"flag"

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/handler/file/click"
	"github.com/TencentAd/attribution/attribution/pkg/handler/file/conv"
	"github.com/TencentAd/attribution/attribution/pkg/storage/attribution"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"

	"github.com/golang/glog"
)

var (
	clickDataPath      = flag.String("click_data_path", "", "")
	conversionDataPath = flag.String("conversion_data_path", "", "")
)

func run() error {
	clickIndex, err := clickindex.CreateClickIndex()
	if err != nil {
		return err
	}

	clickFileHandler := click.NewClickFileHandle(*clickDataPath, clickIndex)
	if err := clickFileHandler.Run(); err != nil {
		glog.Errorf("failed to process click log, err: %v", err)
		return err
	}

	attrStore, err := attribution.CreateAttributionStore()
	if err != nil {
		return err
	}
	convFileHandler := conv.NewConvFileHandle(*conversionDataPath, clickIndex, attrStore)
	if err := convFileHandler.Run(); err != nil {
		glog.Errorf("failed to process conversion data, err: %v", err)
		return err
	}

	return nil
}

func main() {
	flagx.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}
