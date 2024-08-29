package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipt_processor/internal/schema"
	"strings"
)

// processReceipt handles the receipt processing request.
func (s *server) processReceipt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt schema.Receipt
		if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		id, err := s.ReceiptLogic.Process(receipt)
		if err != nil {
			http.Error(w, "Failed to Process", http.StatusInternalServerError)
			return
		}

		// Respond with the ID
		response := map[string]string{"id": id}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// getPoints handles the request to get points for a given receipt ID.
func (s *server) getPoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/receipts/")
		points, err := s.ReceiptLogic.Get(id)
		if err != nil {
			http.Error(w, "Failed to Get", http.StatusInternalServerError)
			return
		}

		// Respond with the points
		response := map[string]string{"points": fmt.Sprintf("%f", points)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
