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
	"os"

	"attribution/pkg/common/flagx"
	metricUtil "attribution/pkg/common/metric-util"
	"attribution/pkg/handler/http/click"
	"attribution/pkg/handler/http/conv"
	"attribution/pkg/parser"
	"attribution/pkg/storage"
	_ "attribution/pkg/storage/all"

	"github.com/golang/glog"
)

var (
	serverAddress        = flag.String("server_address", ":9081", "")
	metricsServerAddress = flag.String("metric_server_address", ":8005", "")
)

func serveHttp() error {
	clickIndex, err := storage.CreateClickIndex()
	if err != nil {
		return err
	}

	impl := &ServerImpl{
		clickIndex: clickIndex,
	}

	if err := impl.initClickHandle(); err != nil {
		return err
	}
	glog.Info("init click handle done")

	if err := impl.initConvHandle(); err != nil {
		return err
	}
	glog.Info("init conv handle done")

	glog.Info("server init done, going to start")
	return http.ListenAndServe(*serverAddress, nil)
}

type ServerImpl struct {
	clickIndex storage.ClickIndex
}

func (s *ServerImpl) initClickHandle() error {
	clickParser, err := parser.CreateClickParser()
	if err != nil {
		return err
	}
	clickHttpHandle := click.NewClickHttpHandle().
		WithParser(clickParser).
		WithClickIndex(s.clickIndex)
	if err := clickHttpHandle.Init(); err != nil {
		return err
	}
	http.Handle("/click", clickHttpHandle)
	return nil
}

func (s *ServerImpl) initConvHandle() error {
	convParser, err := parser.CreateConvParser()
	if err != nil {
		return err
	}

	attributionStores, err := storage.CreateAttributionStore()

	convHttpHandle := conv.NewConvHttpHandle().
		WithParser(convParser).
		WithClickIndex(s.clickIndex).
		WithAttributionStore(attributionStores)
	if err := convHttpHandle.Init(); err != nil {
		return err
	}
	http.Handle("/conv", convHttpHandle)
	return nil
}

func main() {
	flagx.Parse()
	metricUtil.ServeMetrics(*metricsServerAddress)
	if err := serveHttp(); err != nil {
		glog.Errorf("failed to start server, err: %v", err)
		os.Exit(1)
	}
}
