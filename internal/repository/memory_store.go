package repository

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models"
	"github.com/google/uuid"
	"sync"
)

// MemoryStore implements the ReceiptStore interface using in-memory storage
type MemoryStore struct {
	mu       sync.Mutex
	receipts map[string]models.Receipt
}

// NewMemoryStore creates a new MemoryStore instance that implements the ReceiptStore Interface
func NewMemoryStore() ReceiptStore {
	return &MemoryStore{
		receipts: make(map[string]models.Receipt),
	}
}

// AddReceipt adds a new receipt to the in-memory store and returns the generated ID
func (s *MemoryStore) AddReceipt(receipt models.Receipt) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.New().String()
	receipt.ID = id
	s.receipts[id] = receipt
	return id
}

// GetReceiptByID retrieves a receipt by its ID from the in-memory store
func (s *MemoryStore) GetReceiptByID(id string) (models.Receipt, bool) {

	s.mu.Lock()
	defer s.mu.Unlock()
	receipt, exists := s.receipts[id]
	if !exists {
		return models.Receipt{}, false
	}
	return receipt, exists
}
