package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define a new counter metric
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_service_requests_total",
			Help: "Total number of requests received",
		},
		[]string{"method", "status"}, // Optional labels
	)
)

func init() {
	// Register the counter metric with Prometheus
	prometheus.MustRegister(requestCount)
}

func main() {
	// Create a custom server with timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      nil, // Default ServeMux, or provide your custom handler
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, // Optional: time to wait for idle connections
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "OK")
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "READY")
	})

	http.HandleFunc("/log", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Printf("log current time: %s\n", time.Now().String())
	})

	// Expose metrics at /metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Increment the counter (for each request)
		requestCount.WithLabelValues(r.Method, "200").Inc()

		// Log and respond to the request
		fmt.Fprintf(w, "Hello, World! Method: %s", r.Method)
	})

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	log.Println("Stopped server")
}
