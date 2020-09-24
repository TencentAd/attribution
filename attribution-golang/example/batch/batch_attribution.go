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
	"attribution/pkg/handler/file-handler/click"
	"attribution/pkg/handler/file-handler/conv"
	"attribution/pkg/storage"
	"attribution/pkg/storage/hbase"
	"attribution/pkg/storage/native"

	"github.com/golang/glog"
)

var (
	clickDataPath      = flag.String("click_data_path", "", "")
	conversionDataPath = flag.String("conversion_data_path", "", "")
	useHbase           = flag.Bool("use_hbase", false, "")
	useFileStore       = flag.Bool("use_file_store", false, "")
)

func useClickIndex() storage.ClickIndex {
	if *useHbase {
		return hbase.NewClickIndeHbase()
	} else {
		return native.NewClickIndexNative()
	}
}

func useStore() storage.AttributionStore {
	if *useFileStore {
		return native.NewNativeAttributionStore()
	} else {
		return native.NewStdoutAttributionStore()
	}
}

func main() {
	flagx.Parse()

	clickIndex := useClickIndex()

	clickFileHandler := click.NewClickFileHandle(*clickDataPath, clickIndex)
	if err := clickFileHandler.Run(); err != nil {
		glog.Errorf("failed to process click log, err: %v", err)
		return
	}

	attrStore := useStore()
	convFileHandler := conv.NewConvFileHandle(*conversionDataPath, clickIndex, attrStore)
	if err := convFileHandler.Run(); err != nil {
		glog.Errorf("failed to process conversion data, err: %v", err)
		return
	}

}
