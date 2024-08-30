package schema

// Receipt represents a receipt to be processed.
type Receipt struct {
	Retailer     string `json:"retailer" validate:"required,text"`
	PurchaseDate string `json:"purchaseDate" validate:"required,timeDate"`
	PurchaseTime string `json:"purchaseTime" validate:"required,timeHour"`
	Items        []Item `json:"items" validate:"required,min=1,dive"`
	Total        string `json:"total" validate:"required,amount"`
}

// Item represents an item on the receipt.
type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required,text"`
	Price            string `json:"price" validate:"required,amount"`
}
