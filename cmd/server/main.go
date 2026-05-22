package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	port := ":8080"
	publicDir := "./public"

	// Production Practice: Create a custom ServeMux (router) rather than changing the global default mux
	mux := http.NewServeMux()

	// http.FileServer maps browser URL routes directly onto local file folders
	fileServer := http.FileServer(http.Dir(publicDir))
	mux.Handle("/", fileServer)

	// Production Practice: Always configure timeouts explicitly instead of using default server configs.
	// This protects your system allocations against hung connection limits.
	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server actively deployment-ready on http://localhost%s\n", port)
	log.Println("Press Ctrl+C to terminate...")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Critical server error encountered: %v", err)
	}
}
