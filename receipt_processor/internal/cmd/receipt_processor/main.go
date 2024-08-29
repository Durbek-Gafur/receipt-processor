package main

import (
	"log"
	"receipt_processor/internal/server"
)

func main() {
	srv, err := server.InitializeServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	log.Println("Starting server on :8081")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
