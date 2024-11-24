package tests

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/inmemorydb"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/services"
)

func setupService() services.ReceiptService {
	// Instantiate the in-memory store
	store := inmemorydb.NewMemoryStore()
	// Create the service implementation with the store
	return services.NewReceiptServiceImpl(store)
}
