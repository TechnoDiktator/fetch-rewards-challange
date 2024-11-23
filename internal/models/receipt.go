package models

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"math"
	"strconv"
	"strings"
	"time"
)

type Item struct {

	ShortDescription string `json:"short_description"`
	Price            string `json:"price"`
}
type Receipt struct {
	ID           string    `json:"id"`
	Retailer     string    `json:"retailer"`
	PurchaseDate time.Time `json:"purchase_date"`
	PurchaseTime string    `json:"purchase_time"`
	Items        []Item    `json:"items"`
	Total        string    `json:"total"`
}

/*
These rules collectively define how many points should be awarded to a receipt.

One point for every alphanumeric character in the retailer name.
50 points if the total is a round dollar amount with no cents.
25 points if the total is a multiple of 0.25.
5 points for every two items on the receipt.
If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
6 points if the day in the purchase date is odd.
10 points if the time of purchase is after 2:00pm and before 4:00pm.

*/
// CalculateRetailerPoints calculates points based on retailer name length
func (r *Receipt) CalculateRetailerPoints() int {
	logger.Log.Infof("Calculating points based on retailer: %s", r.Retailer)
	points := 0
	// Rule: One point for every alphanumeric character in the retailer name.
	for _, c := range r.Retailer {
		if isAlphanumeric(c) {
			points++
		}
	}
	logger.Log.Infof("Retailer points calculated: %d", points)
	return points
}

// Helper function to check if a character is alphanumeric
func isAlphanumeric(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

// CalculateTotalPoints checks if the total is a round dollar amount (no cents)
func (r *Receipt) CalculateTotalIsRoundDollar() int {
	logger.Log.Infof("Checking if total is a round dollar: %s", r.Total)
	points := 0
	// Rule: 50 points if total is a round dollar amount with no cents.
	total, err := r.parseTotal(r.Total)
	if err == nil && total == math.Floor(total) {
		points = 50
		logger.Log.Infof("Total is a round dollar, points: %d", points)
	} else {
		logger.Log.Infof("Total is not a round dollar")
	}
	return points
}

// Helper function to parse total as float
func (r *Receipt) parseTotal(total string) (float64, error) {
	return strconv.ParseFloat(total, 64)
}

// CalculateItemPoints calculates points for each item in the receipt
func (r *Receipt) CalculateItemPoints() int {
	logger.Log.Infof("Calculating item points for %d items", len(r.Items))
	points := 0
	// Rule: 5 points for every two items.
	points += (len(r.Items) / 2) * 5
	logger.Log.Infof("Item points calculated: %d", points)
	return points
}


// CalculateItemDescriptionPoints calculates points based on item description length
func (r *Receipt) CalculateItemDescriptionPoints() int {
	logger.Log.Infof("Calculating points based on item descriptions")
	points := 0
	// Rule: Points for item description if length is multiple of 3
	for _, item := range r.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			logger.Log.Infof("Item " , item.)
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}
			itemPoints := int(math.Ceil(price * 0.2)) // Multiply price by 0.2 and round up
			points += itemPoints
			logger.Log.Infof("Item description points for %s: %d", item.ShortDescription, itemPoints)
		}
	}
	logger.Log.Infof("Total item description points: %d", points)
	return points
}
