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
