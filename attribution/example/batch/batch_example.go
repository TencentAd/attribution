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

	"attribution/pkg/common/flagx"
	"attribution/pkg/handler/file/click"
	"attribution/pkg/handler/file/conv"
	"attribution/pkg/storage"
	_ "attribution/pkg/storage/all"

	"github.com/golang/glog"
)

var (
	clickDataPath      = flag.String("click_data_path", "", "")
	conversionDataPath = flag.String("conversion_data_path", "", "")
)

func run() error {
	clickIndex, err := storage.CreateClickIndex()
	if err != nil {
		return err
	}

	clickFileHandler := click.NewClickFileHandle(*clickDataPath, clickIndex)
	if err := clickFileHandler.Run(); err != nil {
		glog.Errorf("failed to process click log, err: %v", err)
		return err
	}

	attrStore, err := storage.CreateAttributionStore()
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
