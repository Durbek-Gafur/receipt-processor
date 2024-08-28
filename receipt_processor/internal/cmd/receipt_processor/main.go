package main

import (
	"log"
	"net/http"
	"receipt_processor/internal/api"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {

    // Initialize a cache with a default expiration time of 5 minutes, and purging expired items every 10 minutes
    c := cache.New(5*time.Minute, 10*time.Minute)
    
    http.HandleFunc("/", api.Handler(c))
    log.Println("Starting server on :8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}

 