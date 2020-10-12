/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-4-29 下午3:33
 */

package metricUtil

import (
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Observed interface{
	labels() []string
}

func CalcTimeUsedMicro(startTime time.Time) float64 {
	return float64(time.Since(startTime) / time.Microsecond)
}

func CalcTimeUsedMilli(startTime time.Time) float64 {
	return float64(time.Since(startTime) / time.Millisecond)
}

func ServeMetrics(metricsServerAddress string) error {
	go func() {
		muxProm := http.NewServeMux()
		muxProm.Handle("/metrics", promhttp.Handler())
		var err error
		if err = http.ListenAndServe(metricsServerAddress, muxProm); err != nil {
			glog.Errorf("failed to listen prometheus, err: %v", err)
		}
	}()

	return nil
}
