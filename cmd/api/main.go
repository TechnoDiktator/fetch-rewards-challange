package main

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/handlers"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/repository"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/services"
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set up the store and the service
	store := repository.NewMemoryStore()
	service := services.NewReceiptServiceImpl(store)

	// Set up the Gin router and handlers
	r := gin.Default()
	handler := handlers.NewReceiptHandler(service)

	// Define routes and link them to handlers
	r.POST("/receipts/process", handler.ProcessReceipt)
	r.GET("/receipts/:id/points", handler.GetPoints)

	// Run the server
	err := r.Run(":8080")
	if err != nil {
		logger.Log.Fatalf("Failed to start server: %v", err)
	}
}
