package main

import (
	"fmt"
	"net/http"
	"time"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received : %s %s\n", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Create a request router
	mux := http.NewServeMux()

	// Register for /healthz endpoint
	mux.HandleFunc("/healthz", healthz)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	// Start listening on :8080
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
