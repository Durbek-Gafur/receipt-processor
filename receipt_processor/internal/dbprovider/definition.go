package db

import "receipt_processor/internal/schema"

//go:generate mockgen -source=definition.go -destination=./mocks/mock_db.go -package=mock_db
type DB interface {
	SaveReceipt(id string, receipt schema.Receipt) error
	GetPointByID(id string) (interface{}, error)
}