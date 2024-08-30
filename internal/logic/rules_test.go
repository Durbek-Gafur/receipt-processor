package logic

import (
	"receipt_processor/internal/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphanumericNamePoints(t *testing.T) {
	tests := []struct {
		name      string
		receipt   schema.Receipt
		expected  int
		expectErr bool
	}{
		{
			name: "Valid retailer name with alphanumeric characters",
			receipt: schema.Receipt{
				Retailer: "M&M Corner Market",
			},
			expected: 14,
		},
		{
			name: "Retailer name with special characters",
			receipt: schema.Receipt{
				Retailer: "Shop @ 24/7",
			},
			expected: 7,
		},
		{
			name: "Empty retailer name",
			receipt: schema.Receipt{
				Retailer: "",
			},
			expected: 0,
		},
		{
			name: "Retailer name with spaces only",
			receipt: schema.Receipt{
				Retailer: "     ",
			},
			expected: 0,
		},
		{
			name: "Retailer name with Unicode characters",
			receipt: schema.Receipt{
				Retailer: "Café Déjà Vu",
			},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := alphanumericNamePoints(tt.receipt)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, points)
			}
		})
	}
}

func TestRoundDollarAmountPoints(t *testing.T) {
	tests := []struct {
		name      string
		total     string
		expected  int
		expectErr bool
	}{
		{
			name:     "Round dollar amount",
			total:    "10.00",
			expected: 50,
		},
		{
			name:     "Non-round dollar amount",
			total:    "10.99",
			expected: 0,
		},
		{
			name:     "Zero total",
			total:    "0.00",
			expected: 50,
		},
		{
			name:     "Negative total",
			total:    "-5.00",
			expected: 50,
		},
		{
			name:     "Very large total",
			total:    "999999999999.00",
			expected: 50,
		},
		{
			name:      "Invalid total format",
			total:     "10.0a",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := schema.Receipt{Total: tt.total}
			points, err := roundDollarAmountPoints(receipt)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, points)
			}
		})
	}
}

func TestMultipleOfQuarterPoints(t *testing.T) {
	tests := []struct {
		name      string
		total     string
		expected  int
		expectErr bool
	}{
		{
			name:     "Multiple of 0.25",
			total:    "5.00",
			expected: 25,
		},
		{
			name:     "Not a multiple of 0.25",
			total:    "5.10",
			expected: 0,
		},
		{
			name:     "Edge case near multiple of 0.25 (just below)",
			total:    "1.249999999",
			expected: 0,
		},
		{
			name:     "Negative multiple of 0.25",
			total:    "-0.25",
			expected: 25,
		},
		{
			name:      "Invalid total format",
			total:     "abc",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := schema.Receipt{Total: tt.total}
			points, err := multipleOfQuarterPoints(receipt)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, points)
			}
		})
	}
}

func TestItemsPoints(t *testing.T) {
	tests := []struct {
		name     string
		items    []schema.Item
		expected int
	}{
		{
			name: "Even number of items",
			items: []schema.Item{
				{ShortDescription: "Item 1"},
				{ShortDescription: "Item 2"},
			},
			expected: 5,
		},
		{
			name: "Odd number of items",
			items: []schema.Item{
				{ShortDescription: "Item 1"},
				{ShortDescription: "Item 2"},
				{ShortDescription: "Item 3"},
			},
			expected: 5,
		},
		{
			name:     "No items",
			items:    []schema.Item{},
			expected: 0,
		},
		{
			name: "Large number of items",
			items: []schema.Item{
				{ShortDescription: "Item 1"},
				{ShortDescription: "Item 2"},
				{ShortDescription: "Item 3"},
				{ShortDescription: "Item 4"},
				{ShortDescription: "Item 5"},
				{ShortDescription: "Item 6"},
			},
			expected: 15, // 5 pairs of items
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := schema.Receipt{Items: tt.items}
			points, err := itemsPoints(receipt)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, points)
		})
	}
}

func TestDescriptionLengthPoints(t *testing.T) {
	tests := []struct {
		name      string
		items     []schema.Item
		expected  int
		expectErr bool
	}{
		{
			name: "Description length is a multiple of 3",
			items: []schema.Item{
				{ShortDescription: "  Item12  ", Price: "5.00"},
			},
			expected: 1, // 5.00 * 0.2 = 1
		},
		{
			name: "Description length is not a multiple of 3",
			items: []schema.Item{
				{ShortDescription: "    Item128", Price: "5.00"},
			},
			expected: 0,
		},
		{
			name: "Multiple items with description length as a multiple of 3",
			items: []schema.Item{
				{ShortDescription: "Item12", Price: "5.00"},
				{ShortDescription: "Item56", Price: "10.00"},
			},
			expected: 3, // (5.00 * 0.2 = 1) + (10.00 * 0.2 = 2)
		},
		{
			name:     "Empty description",
			items:    []schema.Item{{ShortDescription: "", Price: "5.00"}},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := schema.Receipt{Items: tt.items}
			points, err := descriptionLengthPoints(receipt)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, points)
			}
		})
	}
}

func TestOddDayPoints(t *testing.T) {
	tests := []struct {
		name      string
		date      string
		expected  int
		expectErr bool
	}{
		{
			name:     "Odd day",
			date:     "2024-08-31",
			expected: 6,
		},
		{
			name:     "Even day",
			date:     "2024-08-30",
			expected: 0,
		},
		{
			name:     "Leap day (odd day)",
			date:     "2024-02-29",
			expected: 6,
		},
		{
			name:     "Leap year (even day)",
			date:     "2024-02-28",
			expected: 0,
		},
		{
			name:      "Invalid date format",
			date:      "31-08-2024",
			expectErr: true,
		},
		{
			name:     "Future date",
			date:     "2100-01-01",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := schema.Receipt{PurchaseDate: tt.date}
			points, err := oddDayPoints(receipt)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, points)
			}
		})
	}
}

func TestTimePoints(t *testing.T) {
	tests := []struct {
		name      string
		time      string
		expected  int
		expectErr bool
	}{
		{
			name:     "Within time range",
			time:     "15:00",
			expected: 10,
		},
		{
			name:     "Outside time range",
			time:     "13:00",
			expected: 0,
		},
		{
			name:     "Boundary time range start",
			time:     "14:00",
			expected: 10,
		},
		{
			name:     "Boundary time range end",
			time:     "16:00",
			expected: 0,
		},
		{
			name:      "Invalid time format",
			time:      "03:00 PM",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receipt := schema.Receipt{PurchaseTime: tt.time}
			points, err := timePoints(receipt)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, points)
			}
		})
	}
}
