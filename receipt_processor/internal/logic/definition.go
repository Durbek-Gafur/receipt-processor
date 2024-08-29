package logic

import "receipt_processor/internal/schema"

//go:generate mockgen -source=definition.go -destination=./mocks/mock_receipt_logic.go -package=mock_receipt_logic
type ReceiptLogic interface{
	Get(id string) (float64,error)
	Process(receipt schema.Receipt) (string,error)  
}