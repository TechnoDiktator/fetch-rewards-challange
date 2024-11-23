package models

import (
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
