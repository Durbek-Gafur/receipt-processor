package db

//go:generate mockgen -source=definition.go -destination=./mocks/mock_db.go -package=mock_db
type DB interface {
	SavePoints(id string, points int) error
	GetPointByReceiptID(id string) (interface{}, error)
}
