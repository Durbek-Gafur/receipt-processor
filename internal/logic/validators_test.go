package logic

import (
	"receipt_processor/internal/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiptValidation(t *testing.T) {
	testValidator := newValidator()
	tests := []struct {
		name    string
		receipt schema.Receipt
		wantErr bool
	}{
		{
			name: "invalid name",
			receipt: schema.Receipt{
				Retailer:     "!M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			wantErr: true,
		},
		{
			name: "valid receipt",
			receipt: schema.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			wantErr: false,
		},
		{
			name: "missing retailer",
			receipt: schema.Receipt{
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			wantErr: true,
		},
		{
			name: "invalid retailer format",
			receipt: schema.Receipt{
				Retailer:     "M&M@Corner Market!",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			wantErr: true,
		},
		{
			name: "invalid purchase date format",
			receipt: schema.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "01/01/2022",
				PurchaseTime: "13:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			wantErr: true,
		},
		{
			name: "invalid purchase time format",
			receipt: schema.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "1:01 PM",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			wantErr: true,
		},
		{
			name: "invalid purchase time format 2",
			receipt: schema.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "25:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.49",
			},
			wantErr: true,
		},
		{
			name: "empty items list",
			receipt: schema.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items:        []schema.Item{},
				Total:        "6.49",
			},
			wantErr: true,
		},
		{
			name: "invalid total format",
			receipt: schema.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				},
				Total: "6.499",
			},
			wantErr: true,
		},
		{
			name: "invalid item description format",
			receipt: schema.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain! Dew 12PK@", Price: "6.49"},
				},
				Total: "6.49",
			},
			wantErr: true,
		},
		{
			name: "invalid item price format",
			receipt: schema.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []schema.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.499"},
				},
				Total: "6.49",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testValidator.Struct(tt.receipt)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
