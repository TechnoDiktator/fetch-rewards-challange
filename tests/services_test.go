package tests

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models/storemodels"
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestCalculateRetailerPoints(t *testing.T) {
	service := setupService()
	// Initialize the logger before each test
	logger.InitializeLogger()
	receipt := storemodels.Receipt{Retailer: "Target"}
	expectedPoints := 6 // "Target" has 6 alphanumeric characters.

	points := service.CalculateRetailerPoints(receipt)

	if points != expectedPoints {
		t.Errorf("expected %d points, got %d", expectedPoints, points)
	}
}

func TestCalculateTotalIsRoundDollar(t *testing.T) {
	service := setupService()
	// Initialize the logger before each test
	logger.InitializeLogger()
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
	// Initialize the logger before each test
	logger.InitializeLogger()
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

/*
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
*/

func TestGetPointsReceipt1(t *testing.T) {
	service := setupService()

	logger.InitializeLogger()
	purchaseDate, err := time.Parse("2006-01-02", "2022-01-01")
	// Add receipt to the store
	receipt := storemodels.Receipt{
		Retailer:     "Target",
		PurchaseDate: purchaseDate,
		PurchaseTime: "13:01",
		Items: []storemodels.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},

		Total: "35.35",
	}
	id, _ := service.ProcessReceipt(receipt)

	// Fetch points
	points, err := service.GetPoints(id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Expected points based on business logic
	expectedPoints := 28 // Retailer and item logic
	if expectedPoints == points {
		logger.Log.Info("============= =============== =============")
		logger.Log.Info("============= Bow Wow Wow !!! =============")
		logger.Log.Info("============== Yippie Yo !!! ==============")
		logger.Log.Info("============== Yippie Ye !!! ==============")
		logger.Log.Info("============= =============== =============")
	}

	if points != expectedPoints {
		t.Errorf("expected %d points, got %d", expectedPoints, points)
	}
}

func TestGetPointsReceipt2(t *testing.T) {
	service := setupService()

	logger.InitializeLogger()
	purchaseDate, err := time.Parse("2006-01-02", "2022-03-20")
	// Add receipt to the store
	receipt := storemodels.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: purchaseDate,
		PurchaseTime: "14:33",
		Items: []storemodels.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},

		Total: "9.00",
	}
	id, _ := service.ProcessReceipt(receipt)

	// Fetch points
	points, err := service.GetPoints(id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Expected points based on business logic
	expectedPoints := 109
	if expectedPoints == points {
		logger.Log.Info("============= =============== =============")
		logger.Log.Info("============= Bow Wow Wow !!! =============")
		logger.Log.Info("============== Yippie Yo !!! ==============")
		logger.Log.Info("============== Yippie Ye !!! ==============")
		logger.Log.Info("============= =============== =============")
	}

	if points != expectedPoints {
		t.Errorf("expected %d points, got %d", expectedPoints, points)
	}
}
