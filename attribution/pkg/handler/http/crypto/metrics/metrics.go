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
			Namespace:   "attribution",
			Subsystem:   "conv_encrypt",
			Name:        "safeguard_internal_err",
			Help:        "safeguard internal err",
		},
		[]string{"type"},
		)
)

func init()  {
	prometheus.MustRegister(ConvEncryptSafeguardErrCount)
}
