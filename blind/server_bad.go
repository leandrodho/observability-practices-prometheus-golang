package blind

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// LegacyProcessHandler simulates an unmonitored HTTP handler with random latency and failure rates.
func LegacyProcessHandler(w http.ResponseWriter, r *http.Request) {
	// Simulating unpredictable processing latency (0 to 500ms)
	sleepTime := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(sleepTime)

	// Simulating sporadic internal errors (10% failure rate)
	if rand.Float32() < 0.1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error [Blind]")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Process completed successfully [Blind]")
}