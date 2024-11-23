package tests

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models/storemodels"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestCalculateRetailerPoints(t *testing.T) {
	service := setupService()

	receipt := storemodels.Receipt{Retailer: "Target"}
	expectedPoints := 6 // "Target" has 6 alphanumeric characters.

	points := service.CalculateRetailerPoints(receipt)

	if points != expectedPoints {
		t.Errorf("expected %d points, got %d", expectedPoints, points)
	}
}

func TestCalculateTotalIsRoundDollar(t *testing.T) {
	service := setupService()

	tests := []struct {
		total    string
		expected int
	}{
		{"35.00", 50}, // Round dollar
		{"35.50", 0},  // Not a round dollar
		{"0.00", 50},  // Edge case
	}

	for _, test := range tests {
		receipt := storemodels.Receipt{Total: test.total}
		points := service.CalculateTotalIsRoundDollar(receipt)
		if points != test.expected {
			t.Errorf("for total %s, expected %d points, got %d", test.total, test.expected, points)
		}
	}
}

// Test for ProcessReceipt with correct time.Time for PurchaseDate
func TestProcessReceipt(t *testing.T) {
	service := setupService()

	// Parse PurchaseDate into a time.Time object
	purchaseDate, err := time.Parse("2006-01-02", "2022-03-20")
	if err != nil {
		t.Fatalf("failed to parse purchase date: %v", err)
	}

	receipt := storemodels.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: purchaseDate, // Assigning time.Time here
		PurchaseTime: "14:33",
		Items: []storemodels.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}

	// Process the receipt and retrieve its ID
	id, err := service.ProcessReceipt(receipt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Ensure the ID is not empty and valid
	if id == "" {
		t.Error("expected a valid receipt ID, got empty string")
	}

	// Validate the ID format using Google's uuid package
	if _, err := uuid.Parse(id); err != nil {
		t.Errorf("invalid receipt ID format: %v", err)
	}
}
