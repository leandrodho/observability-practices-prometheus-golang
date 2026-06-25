package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"observability-practices-prometheus-golang/blind"
	"observability-practices-prometheus-golang/observed"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Seed random generator for runtime environment simulations
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== OBSERVABILITY PRACTICES IN GOLANG ===")
	fmt.Println()

	// =========================================================================
	// 1. BLIND / UNMONITORED ROUTE
	// =========================================================================
	// This endpoint executes completely in the dark with no diagnostic visibility
	http.HandleFunc("/api/v1/blind", blind.LegacyProcessHandler)

	// =========================================================================
	// 2. OBSERVED ROUTE (Wrapped with Prometheus Metrics Middleware)
	// =========================================================================
	// This endpoint tracks volume, tracking anomalies and real-time response durations
	http.HandleFunc("/api/v1/observed", observed.ObserveMetrics(observed.CleanProcessHandler))

	// =========================================================================
	// 3. PROMETHEUS METRICS SCRAPE ENDPOINT
	// =========================================================================
	// Exposing the required gateway path for external systems to gather data logs
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server successfully instantiated on port :2112")
	fmt.Println("👉 Test Blind Target:        http://localhost:2112/api/v1/blind")
	fmt.Println("👉 Test Observed Target:     http://localhost:2112/api/v1/observed")
	fmt.Println("📊 View Live Prometheus Data: http://localhost:2112/metrics")
	fmt.Println()

	if err := http.ListenAndServe(":2112", nil); err != nil {
		fmt.Printf("Server critical startup failure: %v\n", err)
	}
}