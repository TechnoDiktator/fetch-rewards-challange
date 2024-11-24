package handlers

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPoints handles GET /points/{user_id}
// @Summary Get points for a user
// @Description Get the total points a user has earned based on their user ID
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} gin.H{"points": "integer"} "Total points"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /points/{user_id} [get]
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
