package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/thesouldev/goboxd/internal/logger"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte(`{"status":"ok"}`))

	}
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		slog.Debug("new_request", "path", r.URL.Path)
	}
}

func main() {
	// Setup logger
	env := os.Getenv("APP_ENV")        // e.g., "production" or "development"
	logLevel := os.Getenv("LOG_LEVEL") // e.g., "debug", "info", "error"

	// Initialize the app-wide logger
	logger.InitLogger(env, logLevel)

	slog.Info("Application is starting up...")

	// Create a request router
	mux := http.NewServeMux()

	// Register for /healthz endpoint
	mux.HandleFunc("/healthz", healthCheckHandler)

	mux.HandleFunc("/run", runHandler)
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
