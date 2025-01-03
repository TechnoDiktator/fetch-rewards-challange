package handlers

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models/handlermodels"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models/storemodels"
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// ProcessReceipt handles POST /receipts/process
// @Summary Process a receipt
// @Description Process a receipt to calculate rewards and return receipt ID
// @Accept json
// @Produce json
// @Param receipt body handlermodels.RequestReceipt true "Receipt Data"
// @Success 200 {object} gin.H{"id": "string"} "Receipt ID generated"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 400 {object} map[string]string "Validation failed"
// @Failure 400 {object} map[string]string "Invalid date format"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /receipts/process [post]
func (h *ReceiptHandler) ProcessReceipt(c *gin.Context) {
	var req handlermodels.RequestReceipt
	// Bind incoming JSON to RequestReceipt struct
	if err := c.ShouldBindJSON(&req); err != nil {
		// Return error if binding fails
		logger.Log.Errorf("Failed to bind request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"description": "The receipt is invalid", "error": "Invalid request body"})
		return
	}

	// Validate the request using the validator
	if err := h.Validator.Struct(&req); err != nil {
		// Return validation errors
		logger.Log.Errorf("Validation failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"description": "The receipt is invalid", "error": "Validation failed", "details": err.Error()})
		return
	}
	// Parse the purchase date string into a time.Time object
	purchaseDate, err := time.Parse("2006-01-02", req.PurchaseDate)
	if err != nil {
		logger.Log.Errorf("Invalid date format for purchase_date: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"description": "The receipt is invalid", "error": "Invalid date format, expected YYYY-MM-DD"})
		return
	}

	// Convert the request to the internal model
	receipt := storemodels.Receipt{
		Retailer:     req.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: req.PurchaseTime,
		Items:        convertToItemModel(req.Items),
		Total:        req.Total,
	}

	// Call the service to process the receipt
	id, err := h.Service.ProcessReceipt(receipt)
	if err != nil {
		logger.Log.Errorf("Error processing receipt: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"description": "Receipt Not Added", "error": "Failed to process receipt"})
		return
	}

	// Respond with the generated receipt ID
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Helper function to convert request items to model layer items
func convertToItemModel(items []handlermodels.RequestItem) []storemodels.Item {
	var result []storemodels.Item
	for _, item := range items {
		result = append(result, storemodels.Item{
			ShortDescription: item.ShortDescription,
			Price:            item.Price,
		})
	}
	return result
}
