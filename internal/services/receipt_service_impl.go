package services

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/inmemorydb"
)

// ReceiptServiceImpl implements the ReceiptService interface
type ReceiptServiceImpl struct {
	store *inmemorydb.ReceiptStore
}

// NewReceiptServiceImpl returns the implementing object of the ReceiptService interface
func NewReceiptServiceImpl(store inmemorydb.ReceiptStore) ReceiptService {
	return &ReceiptServiceImpl{store: &store}
}
