package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipt_processor/internal/schema"

	"github.com/gorilla/mux"
)

// processReceipt handles the receipt processing request.
func (s *server) processReceipt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt schema.Receipt
		if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		id, err := s.ReceiptLogic.ProcessReceipt(receipt)
		if err != nil {
			http.Error(w, "Failed to ProcessReceipt", http.StatusInternalServerError)
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
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			http.Error(w, "Missing receipt ID", http.StatusBadRequest)
			return
		}
		points, err := s.ReceiptLogic.GetPointByReceiptID(id)
		if err != nil {
			http.Error(w, "Failed to Get", http.StatusInternalServerError)
			return
		}

		// Respond with the points
		response := map[string]string{"points": fmt.Sprintf("%d", points)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// handleMissingID handles the request with missing id
func (s *server) handleMissingID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Missing receipt ID", http.StatusBadRequest)
	}
}
