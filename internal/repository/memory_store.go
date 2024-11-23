package repository

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models"
	"github.com/google/uuid"
	"sync"
)

// This is a simple struct that contains a map and a mutex to serve as the in memory store for us
type MemoryStore struct {
	mu       sync.Mutex
	receipts map[string]models.Receipt
}

func NewMemoryStore() *MemoryStore {
	return
}
