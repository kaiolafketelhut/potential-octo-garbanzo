package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HealthResponse represents the JSON structure for our health check
type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	mux := http.NewServeMux()

	// Root endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Fieldberry Backend is running!\n")
	})
	
	// Health check API endpoint
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := HealthResponse{Status: "healthy"}
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
