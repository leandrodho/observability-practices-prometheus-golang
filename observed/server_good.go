package observed

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// CleanProcessHandler represents a stable enterprise business logic service.
func CleanProcessHandler(w http.ResponseWriter, r *http.Request) {
	// Simulating variable system latency (0 to 400ms)
	sleepTime := time.Duration(rand.Intn(400)) * time.Millisecond
	time.Sleep(sleepTime)

	// Simulating a minor 5% controlled failure rate for monitoring purposes
	if rand.Float32() < 0.05 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error [Observed]")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Success [Observed]")
}