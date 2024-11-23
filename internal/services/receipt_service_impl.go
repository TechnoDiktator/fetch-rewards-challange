package services

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/db"
)

// ReceiptServiceImpl implements the ReceiptService interface
type ReceiptServiceImpl struct {
	store *db.ReceiptStore
}

// NewReceiptServiceImpl returns the implementing object of the ReceiptService interface
func NewReceiptServiceImpl(store db.ReceiptStore) ReceiptService {
	return &ReceiptServiceImpl{store: &store}
}
