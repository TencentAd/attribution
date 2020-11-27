/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/6/20, 10:32 AM
 */

package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	ConvEncryptSafeguardErrCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "attribution",
			Subsystem: "conv_encrypt",
			Name:      "safeguard_internal_err",
			Help:      "safeguard internal err",
		},
		[]string{"type"},
	)
	ConvEncryptHandleCost = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "attribution",
		Subsystem: "conv_encrypt",
		Name:      "http_handle_cost",
		Help:      "http handle cost",
		Buckets:   []float64{1, 5, 10, 20, 50, 80, 100, 200, 500, 1000},
	})
	ConvEncryptErrCount = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "attribution",
		Subsystem: "conv_encrypt",
		Name:      "http_err_count",
		Help:      "http err count",
	})
)

func init() {
	prometheus.MustRegister(ConvEncryptSafeguardErrCount)
	prometheus.MustRegister(ConvEncryptHandleCost)
	prometheus.MustRegister(ConvEncryptErrCount)
}
