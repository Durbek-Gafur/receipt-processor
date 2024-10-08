// Code generated by MockGen. DO NOT EDIT.
// Source: definition.go
//
// Generated by this command:
//
//	mockgen -source=definition.go -destination=./mocks/mock_receipt_logic.go -package=mock_receipt_logic
//

// Package mock_receipt_logic is a generated GoMock package.
package mock_receipt_logic

import (
	schema "receipt_processor/internal/schema"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
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

// GetPointByReceiptID mocks base method.
func (m *MockReceiptLogic) GetPointByReceiptID(id string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPointByReceiptID", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPointByReceiptID indicates an expected call of GetPointByReceiptID.
func (mr *MockReceiptLogicMockRecorder) GetPointByReceiptID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPointByReceiptID", reflect.TypeOf((*MockReceiptLogic)(nil).GetPointByReceiptID), id)
}

// ProcessReceipt mocks base method.
func (m *MockReceiptLogic) ProcessReceipt(receipt schema.Receipt) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessReceipt", receipt)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessReceipt indicates an expected call of ProcessReceipt.
func (mr *MockReceiptLogicMockRecorder) ProcessReceipt(receipt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessReceipt", reflect.TypeOf((*MockReceiptLogic)(nil).ProcessReceipt), receipt)
}
