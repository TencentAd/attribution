package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	DecryptHttpCost = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "attribution",
		Subsystem: "decrypt",
		Name:      "decrypt_http_handle_cost",
		Help:      "decrypt http handle cost",
		Buckets:   []float64{1, 5, 10, 20, 50, 80, 100, 200, 500, 1000},
	})

	DecryptErrCount = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   "attribution",
		Subsystem:   "decrypt",
		Name:        "decrypt_err_count",
		Help:        "decrypt err count",
	})
)

func init()  {
	prometheus.MustRegister(DecryptHttpCost)
	prometheus.MustRegister(DecryptErrCount)
}
