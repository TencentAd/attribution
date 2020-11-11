/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 5:35 PM
 */

package main

import (
	"flag"
	"net/http"

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/common/metricutil"
	"github.com/TencentAd/attribution/attribution/pkg/leads/server/handle"
	"github.com/TencentAd/attribution/attribution/pkg/storage/leads"
	"github.com/golang/glog"
)

var (
	serverAddress        = flag.String("server_address", ":9083", "")
	metricsServerAddress = flag.String("metric_server_address", ":8005", "")
)

func serveHttp() error {
	leadsStorage, err := leads.CreateLeadsStorage()
	if err != nil {
		return err
	}
	receiveLeads := handle.NewReceiveLeadsHandle().
		WithLeadsStorage(leadsStorage)

	http.Handle("/leads", receiveLeads)
	glog.Info("init done")
	return http.ListenAndServe(*serverAddress, nil)
}

func main() {
	if err := flagx.Parse(); err != nil {
		panic(err)
	}
	metricutil.ServeMetrics(*metricsServerAddress)
	if err := serveHttp(); err != nil {
		glog.Errorf("failed to start server, err: %v", err)
	}
}
