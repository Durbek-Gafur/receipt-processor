package server

import (
	"fmt"
	"net/http"
	"receipt_processor/internal/logic"

	"github.com/gorilla/mux"
)
type Server interface{
	ListenAndServe() error
}
// server implements the Server interface.
type server struct {
	Router       *mux.Router
	ReceiptLogic logic.ReceiptLogic
}

// NewServer creates a new Server instance.
func NewServer(receiptLogic logic.ReceiptLogic) Server {
	router := mux.NewRouter()
	server := &server{
		Router:       router,
		ReceiptLogic: receiptLogic,
	}
	server.setupRoutes()
	return server
}

// setupRoutes sets up the routes for the server.
func (s *server) setupRoutes() {
	s.Router.SkipClean(true)
	s.Router.HandleFunc("/receipts/process", s.processReceipt()).Methods("POST")
	s.Router.HandleFunc("/receipts/{id}/points", s.getPoints()).Methods("GET")
	s.Router.HandleFunc("/receipts//points", s.handleMissingID()).Methods("GET")
	s.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Printf("Registered route: %s %v\n", pathTemplate, methods)
		return nil
	})
	
}

// ListenAndServe starts the HTTP server.
func (s *server) ListenAndServe() error {
	return http.ListenAndServe(":8081", s.Router)
}
