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

	"attribution/pkg/common/flagx"
	"attribution/pkg/leads/server/handle"
	"attribution/pkg/storage/redis"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	serverAddress        = flag.String("server_address", ":9083", "")
	metricsServerAddress = flag.String("metric_server_address", ":8005", "")
)

func serveMetrics() error {
	go func() {
		muxProm := http.NewServeMux()
		muxProm.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(*metricsServerAddress, muxProm); err != nil {
			glog.Errorf("failed to listen prometheus, err: %v", err)
		}
	}()

	return nil
}

func serveHttp() error {
	leadsStorage := redis.NewLeadsRedisStorage()
	if err := leadsStorage.Init(); err != nil {
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
	if err := serveMetrics(); err != nil {
		panic(err)
	}
	if err := serveHttp(); err != nil {
		glog.Errorf("failed to start server, err: %v", err)
	}
}
