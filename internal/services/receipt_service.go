package services

import (
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models"
)

type ReceiptService interface {
	ProcessReceipt(receipt models.Receipt) (string, error)
	GetPoints(id string) (int, error)
}
