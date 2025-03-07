package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

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

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	log.Println("Stopped server")
}
