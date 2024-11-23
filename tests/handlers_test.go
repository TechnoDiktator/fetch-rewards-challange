package tests

import (
	"bytes"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/db"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/handlers"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models/storemodels"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestService() services.ReceiptService {
	// Mock or initialize the service
	store := db.NewMemoryStore()                 // Or any other store you want to use
	return services.NewReceiptServiceImpl(store) // This function should return a pointer, which is correct for our handler
}

func TestProcessReceipt_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode) // Use test mode to disable logging
	router := gin.Default()
	receiptService := setupTestService()

	// Create a handler instance and attach it to the router
	receiptHandler := handlers.ReceiptHandler{receiptService}
	router.POST("/receipts/process", receiptHandler.ProcessReceipt)

	// Prepare test data
	receipt := storemodels.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []storemodels.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "2.25",
	}

	// Marshal the data into JSON
	receiptJSON := `{
		"retailer": "M&M Corner Market",
		"purchaseDate": "2022-03-20",
		"purchaseTime": "14:33",
		"items": [{"shortDescription": "Gatorade", "price": "2.25"}],
		"total": "2.25"
	}`

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	w := httptest.NewRecorder()

	// Perform the request using the Gin router
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// You can also check the response body if needed
	assert.Contains(t, w.Body.String(), "id")
}

func TestProcessReceipt_InvalidJSON(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	receiptService := setupTestService()

	// Create a handler instance and attach it to the router
	receiptHandler := NewReceiptHandler(receiptService)
	router.POST("/receipts/process", receiptHandler.ProcessReceipt)

	// Invalid JSON data
	invalidJSON := `{"invalid: "data"}`

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer([]byte(invalidJSON)))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	w := httptest.NewRecorder()

	// Perform the request using the Gin router
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Check if error message is in the response body
	assert.Contains(t, w.Body.String(), "error")
}
