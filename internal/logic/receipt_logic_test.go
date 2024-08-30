package logic

import (
	"errors"
	mock_db "receipt_processor/internal/dbprovider/mocks"
	"receipt_processor/internal/schema"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProcessReceipt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_db.NewMockDB(ctrl)
	logic := NewReceiptLogic(ProvideOptions(mockDB))

	receipt1 := schema.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []schema.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
		Total: "35.35",
	}

	receipt2 := schema.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []schema.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}

	t.Run("Success - Receipt 1", func(t *testing.T) {
		// Mock database save
		mockDB.EXPECT().SavePoints(gomock.Any(), 28).Return(nil)

		id, err := logic.ProcessReceipt(receipt1)

		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("Success - Receipt 2", func(t *testing.T) {
		// Mock database save
		mockDB.EXPECT().SavePoints(gomock.Any(), 109).Return(nil)

		id, err := logic.ProcessReceipt(receipt2)

		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})

	t.Run("Validation Error", func(t *testing.T) {
		invalidReceipt := schema.Receipt{
			Retailer: "Target",
			Total:    "invalid",
		}

		id, err := logic.ProcessReceipt(invalidReceipt)

		assert.Error(t, err)
		assert.Empty(t, id)
	})

	t.Run("Database Save Error", func(t *testing.T) {
		// Mock database save
		mockDB.EXPECT().SavePoints(gomock.Any(), 28).Return(errors.New("db save error"))

		id, err := logic.ProcessReceipt(receipt1)

		assert.Error(t, err)
		assert.Empty(t, id)
	})
}
