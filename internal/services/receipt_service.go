package services

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models"
)

//// Rule 1: Points based on the retailer name (1 point for every alphanumeric character)
//points += s.calculateRetailerPoints(receipt.Retailer)
//
//// Rule 2: 50 points if the total is a round dollar amount
//points += s.calculateTotalIsRoundDollar(receipt.Total)
//
//// Rule 3: 25 points if the total is a multiple of 0.25
//points += s.calculateTotalIsMultipleOfQuarter(receipt.Total)
//
//// Rule 4: 5 points for every two items on the receipt
//points += s.calculateItemsPoints(receipt.Items)
//
//// Rule 5: Points based on item description length (multiple of 3)
//points += s.calculateItemDescriptionPoints(receipt.Items)
//
//// Rule 6: 6 points if the purchase date is an odd day
//points += s.calculatePurchaseDatePoints(receipt)
//
//// Rule 7: 10 points if the purchase time is between 2:00pm and 4:00pm
//points += s.calculatePurchaseTimePoints(receipt)

// Log the final points calculated for the receipt

type ReceiptService interface {
	ProcessReceipt(receipt models.Receipt) (string, error)
	GetPoints(id string) (int, error)

	// Points Calculation Methods
	CalculateRetailerPoints(receipt models.Receipt) int

	CalculateTotalIsRoundDollar(receipt models.Receipt) int

	CalculateTotalMultipleOfQuarter(receipt models.Receipt) int

	CalculateItemPoints(receipt models.Receipt) int

	CalculateItemDescriptionPoints(receipt models.Receipt) int

	CalculatePurchaseDatePoints(receipt models.Receipt) int

	CalculatePurchaseTimePoints(receipt models.Receipt) int

	CalculateTotalPoints(receipt models.Receipt) int
}
