package services

import (
	"fmt"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/repository"
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"github.com/sirupsen/logrus"
)

// ReceiptServiceImpl implements the ReceiptService interface
type ReceiptServiceImpl struct {
	store repository.ReceiptStore
}

func NewReceiptServiceImpl(store repository.ReceiptStore) ReceiptService {
	return &ReceiptServiceImpl{store: store}
}

func (s *ReceiptServiceImpl) ProcessReceipt(receipt models.Receipt) (string, error) {

	logger.Log.WithFields(logrus.Fields{
		"retailer": receipt.Retailer,
		"total":    receipt.Total,
		"items":    len(receipt.Items),
	}).Infof("Processing receipt")

	id := s.store.AddReceipt(receipt)
	logger.Log.WithFields(logrus.Fields{
		"id": id,
	}).Info("Receipt processed successfully")

}

func (s *ReceiptServiceImpl) GetPoints(id string) (int, error) {

	logger.Log.WithFields(logrus.Fields{
		"id": id,
	}).Infoln("Get points for receipt")
	receipt, exists := s.store.GetReceiptByID(id)
	if !exists {
		logger.Log.WithFields(logrus.Fields{
			"id": id,
		}).Error("Receipt not found")
		return 0, fmt.Errorf("receipt not found")
	}
	// Log the receipt data for debugging
	logger.Log.Infof("======== Get points for receipt with id %s =========", id)
	var points int
	points = 0

	// Rule 1: Points for the retailer name

	points += len(receipt.Retailer)
	logger.Log.Infoln("Points For Retailer Name %d", points)

}
