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

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/common/metricutil"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/click"
	"github.com/TencentAd/attribution/attribution/pkg/handler/http/conv"
	"github.com/TencentAd/attribution/attribution/pkg/parser"
	"github.com/TencentAd/attribution/attribution/pkg/storage/attribution"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"

	"github.com/golang/glog"
)

var (
	serverAddress        = flag.String("server_address", ":9081", "")
	metricsServerAddress = flag.String("metric_server_address", ":8005", "")
)

func serveHttp() error {
	clickIndex, err := clickindex.CreateClickIndex()
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
	clickIndex clickindex.ClickIndex
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

	attributionStores, err := attribution.CreateAttributionStore()
	if err != nil {
		return err
	}

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
	if err := flagx.Parse(); err != nil {
		panic(err)
	}
	_ = metricutil.ServeMetrics(*metricsServerAddress)
	if err := serveHttp(); err != nil {
		glog.Errorf("failed to start server, err: %v", err)
		os.Exit(1)
	}
}
