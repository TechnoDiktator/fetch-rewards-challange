package tests

import (
	"bytes"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/handlers"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/utils/constants"
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessReceipt_Success(t *testing.T) {
	// Setup
	logger.InitializeLogger()
	logger.Log.Info("=========================================================================")
	logger.Log.Info()
	gin.SetMode(gin.TestMode) // Use test mode to disable logging
	router := gin.Default()
	receiptService := setupService()
	//purchaseDate, err := time.Parse("2006-01-02", "2024-03-20")

	// Create a handler instance and attach it to the router
	receiptHandler := handlers.ReceiptHandler{Service: receiptService,
		Validator: validator.New()}
	router.POST("/receipts/process", receiptHandler.ProcessReceipt)

	receiptJSON := `{
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
	}`

	//purchaseDate, err := time.Parse("2006-01-02", "2024-03-20")
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, constants.ProcessReceipts, bytes.NewBuffer([]byte(receiptJSON)))
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
	logger.Log.Info("=========================================================================")
	logger.Log.Info()
}

func TestProcessReceipt_InvalidJSON(t *testing.T) {
	//declare logger
	logger.InitializeLogger()
	logger.Log.Info("=========================================================================")
	logger.Log.Info()
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	receiptService := setupService()
	// Create a handler instance and attach it to the router
	receiptHandler := handlers.ReceiptHandler{Service: receiptService,
		Validator: validator.New()}

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

	logger.Log.Info("=========================================================================")
	logger.Log.Info()
}
