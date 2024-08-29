package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	mock_receipt_logic "receipt_processor/internal/logic/mocks"
	"receipt_processor/internal/schema"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)
func TestProcessReceipt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReceiptLogic := mock_receipt_logic.NewMockReceiptLogic(ctrl)
	s := &server{ReceiptLogic: mockReceiptLogic}

	tests := []struct {
		name           string
		inputBody      interface{}
		mockProcess    func(m *mock_receipt_logic.MockReceiptLogic)
		expectedStatus int
		expectedBody    string
	}{
		{
			name:           "Successful Receipt Processing",
			inputBody:      schema.Receipt{ /* populate with valid data */ },
			mockProcess:    func(m *mock_receipt_logic.MockReceiptLogic) { m.EXPECT().Process(gomock.Any()).Return("12345", nil) },
			expectedStatus: http.StatusOK,
			expectedBody:    `{"id":"12345"}`,
		},
		{
			name:           "Invalid Request Payload",
			inputBody:      "invalid payload",
			mockProcess:    func(m *mock_receipt_logic.MockReceiptLogic) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request payload",
		},
		{
			name:           "Failed Receipt Processing",
			inputBody:      schema.Receipt{ /* populate with valid data */ },
			mockProcess:    func(m *mock_receipt_logic.MockReceiptLogic) { m.EXPECT().Process(gomock.Any()).Return("", assert.AnError) },
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to Process",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare the request
			body, _ := json.Marshal(tt.inputBody)
			req := httptest.NewRequest(http.MethodPost, "/receipt", bytes.NewReader(body))
			w := httptest.NewRecorder()

			// Setup the mock
			tt.mockProcess(mockReceiptLogic)

			// Call the handler
			handler := s.processReceipt()
			handler.ServeHTTP(w, req)

			// Check the status code
			assert.Equal(t, tt.expectedStatus, w.Result().StatusCode)

			// Check the response body
			if tt.expectedStatus == http.StatusOK {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			} else {
				assert.Contains(t, w.Body.String(), tt.expectedBody)
			}
		})
	}
}
