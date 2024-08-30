package logic

import (
	"fmt"
	"math"
	"receipt_processor/internal/schema"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Rule function type
type Rule func(receipt schema.Receipt) (int, error)

// Option function type
type Option func(*ReceiptLogicOptions)

// registerRules adds all default rules to the ReceiptLogicOptions.
func (r *receiptLogic) registerRules() {
	r.Options.Rules = []Rule{
		alphanumericNamePoints,
		roundDollarAmountPoints,
		multipleOfQuarterPoints,
		itemsPoints,
		descriptionLengthPoints,
		oddDayPoints,
		timePoints,
	}
}

func alphanumericNamePoints(receipt schema.Receipt) (points int, err error) {
	for _, ch := range receipt.Retailer {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			points++
		}
	}
	return points, nil
}

func roundDollarAmountPoints(receipt schema.Receipt) (points int, err error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return points, err
	}
	if total == math.Floor(total) {
		return 50, nil
	}
	return 0, nil
}

func multipleOfQuarterPoints(receipt schema.Receipt) (points int, err error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return points, err
	}
	if math.Mod(total*100, 25) == 0 {
		return 25, nil
	}
	return 0, nil
}

func itemsPoints(receipt schema.Receipt) (points int, err error) {
	return 5 * (len(receipt.Items) / 2), nil
}

func descriptionLengthPoints(receipt schema.Receipt) (points int, err error) {
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 && len(desc) != 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				fmt.Println("Error converting string to float:", err)
				return points, err
			}
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points, nil
}

func oddDayPoints(receipt schema.Receipt) (points int, err error) {
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	day := purchaseDate.Day()
	if day%2 != 0 {
		return 6, nil
	}
	return 0, nil
}

func timePoints(receipt schema.Receipt) (points int, err error) {
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	hour := purchaseTime.Hour()
	if hour >= 14 && hour < 16 {
		return 10, nil
	}
	return 0, nil
}
