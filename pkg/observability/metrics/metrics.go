package metrics

import (
	"log"
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "endpoint", "status"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for handler.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint", "status"},
	)

	registerOnce sync.Once
)

func RegisterMetric() {
	log.Println("Registering metrics...")
	registerOnce.Do(func() {
		prometheus.MustRegister(RequestCount, RequestDuration)
	})
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
