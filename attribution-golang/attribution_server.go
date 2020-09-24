/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:11 AM
 */

package main

import (
	"flag"
	"net/http"

	"attribution/pkg/common/flagx"
	"attribution/pkg/handler/http-handler/click"
	"attribution/pkg/handler/http-handler/conv"
	"attribution/pkg/parser/click_parser"
	"attribution/pkg/parser/conv_parser"
	"attribution/pkg/storage/native"
)

var (
	serverAddress        = flag.String("server_address", "", "")
	metricsServerAddress = flag.String("metric_server_address", "", "")
)

func serveHttp() error {
	// 点击日志索引
	clickIndex := native.NewClickIndexNative()

	// 点击日志handle
	clickHttpHandle := click.NewClickHttpHandle().
		WithParser(&click_parser.ClickParserStruct{}).
		WithClickIndex(clickIndex)
	if err := clickHttpHandle.Init(); err != nil {
		return err
	}
	http.Handle("/click", clickHttpHandle)

	// 转化日志handle
	convHttpHandle := conv.NewConvHttpHandle().
		WithParser(&conv_parser.ConvParserStruct{}).
		WithClickIndex(clickIndex).
		WithAttributionStore(native.NewStdoutAttributionStore())
	if err := convHttpHandle.Init(); err != nil {
		return err
	}
	http.Handle("/conv", convHttpHandle)
	return http.ListenAndServe(*serverAddress, nil)
}

func main() {
	flagx.Parse()
	serveHttp()
}
