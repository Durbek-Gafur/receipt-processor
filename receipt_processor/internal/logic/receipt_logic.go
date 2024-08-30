package logic

import (
	db "receipt_processor/internal/dbprovider"
	"receipt_processor/internal/schema"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ProvideOptions(db db.DB) ReceiptLogicOptions {
	return ReceiptLogicOptions{DB: db}
}

type ReceiptLogicOptions struct {
	DB        db.DB
	Rules     []Rule
	Validator *validator.Validate
}
type receiptLogic struct {
	Options ReceiptLogicOptions
}

func NewReceiptLogic(options ReceiptLogicOptions) ReceiptLogic {
	logic := &receiptLogic{Options: options}
	logic.registerRules()
	logic.Options.Validator = newValidator()
	return logic
}

func (r *receiptLogic) GetPointByReceiptID(i string) (int, error) {
	point, err := r.Options.DB.GetPointByReceiptID(i)
	if err != nil {
		return 0, err
	}
	return point, nil
}

func (r *receiptLogic) ProcessReceipt(receipt schema.Receipt) (string, error) {
	// validate input
	err := r.Options.Validator.Struct(receipt)
	if err != nil {
		// TODO add stacktrace
		return "", err
	}
	// apply rules
	totalPoints := 0
	for _, rule := range r.Options.Rules {
		points, err := rule(receipt)
		if err != nil {
			return "", err
		}
		totalPoints += points
	}

	// TODO if the methods needs to be idemptotent hash receipt and create uuid from hash
	id := uuid.NewString()
	// persist in memory
	err = r.Options.DB.SavePoints(id, totalPoints)
	if err != nil {
		return "", err
	}

	// return id
	return id, nil
}
