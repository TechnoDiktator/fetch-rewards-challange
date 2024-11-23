package tests

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/repository"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/services"
)

func setupService() *services.ReceiptServiceImpl {
	// Instantiate the in-memory store
	store := repository.NewMemoryStore()
	// Create the service implementation with the store
	return services.NewReceiptServiceImpl(store).(*services.ReceiptServiceImpl)
}
