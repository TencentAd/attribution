package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/TencentAd/attribution/attribution/pkg/common/metricutil"
)

var (
	HandleErrCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "attribution",
			Subsystem: "imp_attribution",
			Name:      "handle_err_count",
			Help:      "handle err count",
		})
	HandleCost = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "attribution",
			Subsystem: "imp_attribution",
			Name:      "handle_cost",
			Help:      "handle cost",
			Buckets:   []float64{1, 2, 5, 10, 20, 50, 100, 500, 1000, 2000, 5000, 10000},
		})

	ActionErrCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "attribution",
			Subsystem: "imp_attribution",
			Name:      "action_err_count",
			Help:      "action err count",
		},
		[]string{"name"})

	ActionCost = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "attribution",
			Subsystem: "imp_attribution",
			Name:      "action_cost",
			Help:      "action cost",
			Buckets:   []float64{1, 2, 5, 10, 20, 50, 100, 500, 1000, 2000, 5000},
		},
		[]string{"name"})
)

func init() {
	prometheus.MustRegister(HandleErrCount)
	prometheus.MustRegister(HandleCost)
	prometheus.MustRegister(ActionErrCounter)
	prometheus.MustRegister(ActionCost)
}

func CollectHandleMetrics(startTime time.Time, err error) {
	metricutil.CollectMetrics(HandleErrCount, HandleCost, startTime, err)
}

func CollectActionMetrics(name string, startTime time.Time, err error) {
	metricutil.CollectActionMetrics(ActionErrCounter, ActionCost, name, startTime, err)
}
