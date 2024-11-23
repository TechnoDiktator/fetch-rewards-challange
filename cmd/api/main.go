package main

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/handlers"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/middlewares"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/repository"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/services"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/utils/constants"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Apply the logging middleware globally
	r.Use(middleware.LogRequest)

	// Initialize the in-memory store for receipt data
	store := repository.NewMemoryStore()

	// Initialize the ReceiptService with the store
	receiptService := services.NewReceiptServiceImpl(store)

	// Initialize the ReceiptHandler with the service
	receiptHandler := handlers.NewReceiptHandler(receiptService)

	// Define routes
	r.POST(constants.ProcessReceipts, receiptHandler.ProcessReceipt)
	r.GET(constants.GetPoints, receiptHandler.GetPoints)

	// Start the Gin server on port 8080
	r.Run(":8080")
}
