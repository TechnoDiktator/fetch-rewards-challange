package services

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/repository"
)

// ReceiptServiceImpl implements the ReceiptService interface
type ReceiptServiceImpl struct {
	store repository.ReceiptStore
}

// NewReceiptServiceImpl returns the implementing object of the ReceiptService interface
func NewReceiptServiceImpl(store repository.ReceiptStore) ReceiptService {
	return &ReceiptServiceImpl{store: store}
}
