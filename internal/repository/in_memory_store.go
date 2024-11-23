package repository

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models/storemodels"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"sync"
)

// MemoryStore implements the ReceiptStore interface using in-memory storage
type MemoryStore struct {
	mu       sync.Mutex
	receipts map[string]storemodels.Receipt
}

// NewMemoryStore creates a new MemoryStore instance that implements the ReceiptStore Interface
func NewMemoryStore() ReceiptStore {
	logrus.Infof("Getting In Memory Store")
	return &MemoryStore{
		receipts: make(map[string]storemodels.Receipt),
	}
}

// AddReceipt adds a new receipt to the in-memory store and returns the generated ID
func (s *MemoryStore) AddReceipt(receipt storemodels.Receipt) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	logrus.Infof("adding receipt In The In Memory Store")

	id := uuid.New().String()
	receipt.ID = id
	s.receipts[id] = receipt
	return id
}

// GetReceiptByID retrieves a receipt by its ID from the in-memory store
func (s *MemoryStore) GetReceiptByID(id string) (storemodels.Receipt, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	logrus.Infof("getting receipt From In Memory Store")

	receipt, exists := s.receipts[id]
	if !exists {
		return storemodels.Receipt{}, false
	}
	return receipt, exists
}
