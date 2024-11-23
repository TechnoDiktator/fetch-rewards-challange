package handlers

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/services"
	"github.com/go-playground/validator/v10"
)

// ReceiptHandler represents the handler for receipt-related actions
type ReceiptHandler struct {
	Service   services.ReceiptService
	Validator *validator.Validate
}

// NewReceiptHandler creates a new instance of ReceiptHandler
func NewReceiptHandler(service services.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{
		Service:   service,
		Validator: validator.New(),
	}
}
