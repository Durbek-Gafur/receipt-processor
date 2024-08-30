package logic

import "receipt_processor/internal/schema"

//go:generate mockgen -source=definition.go -destination=./mocks/mock_receipt_logic.go -package=mock_receipt_logic
type ReceiptLogic interface {
	// GetPointByReceiptID gets point for given receipt id
	GetPointByReceiptID(id string) (int, error)
	// ProcessReceipt processes receipt and persists in db
	ProcessReceipt(receipt schema.Receipt) (string, error)
}
