package observed

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Define global Prometheus metrics with explicit labels for granular breakdown analysis.
var (
	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed.",
		},
		[]string{"path", "status"}, // Labels allow deep filtering by endpoint and status code
	)

	HttpRequestDurationSeconds = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response latencies for HTTP requests.",
			Buckets: []float64{0.05, 0.1, 0.25, 0.5, 1.0, 2.5}, // Time boundaries defined in seconds
		},
		[]string{"path"},
	)
)

// StatusRecorder helper to intercept and capture the HTTP status code on the fly.
type StatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (rec *StatusRecorder) WriteHeader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

// ObserveMetrics is an instrumentation middleware layer to capture telemetry metrics.
func ObserveMetrics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		path := r.URL.Path

		rec := &StatusRecorder{ResponseWriter: w, StatusCode: http.StatusOK}
		
		next.ServeHTTP(rec, r)

		// Calculate total request execution time
		duration := time.Since(startTime).Seconds()

		// Record observations to the global Prometheus registry
		HttpRequestsTotal.WithLabelValues(path, strconv.Itoa(rec.StatusCode)).Inc()
		HttpRequestDurationSeconds.WithLabelValues(path).Observe(duration)
	}
}