/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:28 AM
 */

package workflow

import "github.com/prometheus/client_golang/prometheus"

var (
	TaskTimeCost = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "adshonor",
			Subsystem: "task",
			Name:      "time_cost",
			Help:      "time_cost",
			Buckets:   []float64{10, 100, 1000, 5000, 10000, 20000, 50000, 100000, 200000, 400000, 1000000},
		},
		[]string{"name"},
	)
)

func init() {
	prometheus.MustRegister(TaskTimeCost)
}
