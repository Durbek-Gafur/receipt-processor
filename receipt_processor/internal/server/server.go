package server

import (
	"net/http"
	"receipt_processor/internal/logic"
)
type Server interface{
	ListenAndServe() error
}
// server implements the Server interface.
type server struct {
	ReceiptLogic logic.ReceiptLogic
}

// NewServer creates a new Server instance.
func NewServer(receiptLogic logic.ReceiptLogic) Server {
	return &server{ReceiptLogic: receiptLogic}
}

// setupRoutes sets up the routes for the server.
func (s *server) setupRoutes() {
	http.HandleFunc("/receipts/process", s.processReceipt())
	http.HandleFunc("/receipts/", s.getPoints())
}

// ListenAndServe starts the HTTP server.
func (s *server) ListenAndServe() error {
	s.setupRoutes()
	return http.ListenAndServe(":8081", nil)
}
