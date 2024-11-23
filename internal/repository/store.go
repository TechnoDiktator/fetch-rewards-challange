package repository

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models/storemodels"
)

// Interface the defines the methods that the In memory Memory store will implement
type ReceiptStore interface {
	//Adds a new Receipt To the Implementor of this interface
	AddReceipt(receipt storemodels.Receipt) string

	//Gets the Receipt By Id . The implementer will return the receipt object
	GetReceiptByID(id string) (storemodels.Receipt, bool)
}
