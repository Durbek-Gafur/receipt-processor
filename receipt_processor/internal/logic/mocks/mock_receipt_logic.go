// Code generated by MockGen. DO NOT EDIT.
// Source: definition.go

// Package mock_receipt_logic is a generated GoMock package.
package mock_receipt_logic

import (
	schema "receipt_processor/internal/schema"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReceiptLogic is a mock of ReceiptLogic interface.
type MockReceiptLogic struct {
	ctrl     *gomock.Controller
	recorder *MockReceiptLogicMockRecorder
}

// MockReceiptLogicMockRecorder is the mock recorder for MockReceiptLogic.
type MockReceiptLogicMockRecorder struct {
	mock *MockReceiptLogic
}

// NewMockReceiptLogic creates a new mock instance.
func NewMockReceiptLogic(ctrl *gomock.Controller) *MockReceiptLogic {
	mock := &MockReceiptLogic{ctrl: ctrl}
	mock.recorder = &MockReceiptLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiptLogic) EXPECT() *MockReceiptLogicMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockReceiptLogic) Get(id string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReceiptLogicMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReceiptLogic)(nil).Get), id)
}

// Process mocks base method.
func (m *MockReceiptLogic) Process(receipt schema.Receipt) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", receipt)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Process indicates an expected call of Process.
func (mr *MockReceiptLogicMockRecorder) Process(receipt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockReceiptLogic)(nil).Process), receipt)
}
