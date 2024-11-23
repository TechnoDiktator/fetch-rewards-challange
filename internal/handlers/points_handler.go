package handlers

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPoints handles GET /receipts/{id}/points

func (h *ReceiptHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")

	// Fetch the points for the given receipt ID from the service
	points, err := h.Service.GetPoints(id)
	if err != nil {
		logger.Log.Errorf("Failed to get points for receipt ID %s: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	// Respond with the calculated points
	c.JSON(http.StatusOK, gin.H{"points": points})
}
