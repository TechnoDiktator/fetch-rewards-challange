package services

import (
	"fmt"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/models"
	"github.com/TechnoDiktator/fetch-rewards-challange/internal/repository"
	"github.com/TechnoDiktator/fetch-rewards-challange/pkg/logger"
	"github.com/sirupsen/logrus"
	"math"
	"strconv"
	"strings"
	"time"
)

// ReceiptServiceImpl implements the ReceiptService interface
type ReceiptServiceImpl struct {
	store repository.ReceiptStore
}

// NewReceiptServiceImpl returns the implementing object of the ReceiptService interface
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

	return id, nil

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

	points = s.CalculateTotalPoints(receipt)

	logger.Log.Infof("Total points calculated for receipt with ID %s: %d", id, points)

	return points, nil

}

// Points Calculation Methods
func (s *ReceiptServiceImpl) CalculateRetailerPoints(receipt models.Receipt) int {
	logrus.Infof("Calculating points based on retailer: %s", receipt.Retailer)
	points := 0
	for _, c := range receipt.Retailer {
		if isAlphanumeric(c) {
			points++
		}
	}
	logrus.Infof("Retailer points calculated: %d", points)
	return points
}

func isAlphanumeric(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func (s *ReceiptServiceImpl) CalculateTotalIsRoundDollar(receipt models.Receipt) int {
	logrus.Infof("Checking if total is a round dollar: %s", receipt.Total)
	points := 0
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && total == math.Floor(total) {
		points = 50
		logrus.Infof("Total is a round dollar, points: %d", points)
	} else {
		logrus.Infof("Total is not a round dollar")
	}
	return points
}

func (s *ReceiptServiceImpl) CalculateTotalMultipleOfQuarter(receipt models.Receipt) int {
	logrus.Infof("Checking if total is a multiple of 0.25: %s", receipt.Total)
	points := 0
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && math.Mod(total, 0.25) == 0 {
		points = 25
		logrus.Infof("Total is a multiple of 0.25, points: %d", points)
	} else {
		logrus.Infof("Total is not a multiple of 0.25")
	}
	return points
}

func (s *ReceiptServiceImpl) CalculateItemPoints(receipt models.Receipt) int {
	logrus.Infof("Calculating item points for %d items", len(receipt.Items))
	points := 0
	points += (len(receipt.Items) / 2) * 5
	logrus.Infof("Item points calculated: %d", points)
	return points
}

func (s *ReceiptServiceImpl) CalculateItemDescriptionPoints(receipt models.Receipt) int {
	logrus.Infof("Calculating points based on item descriptions")
	points := 0
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			logrus.Infof("Trimmed Item Description is %s and is divisible by 3", item.ShortDescription)
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
			logrus.Infof("Item description points for %s: %d", item.ShortDescription, itemPoints)
		}
	}
	logrus.Infof("Total item description points: %d", points)
	return points
}

func (s *ReceiptServiceImpl) CalculatePurchaseDatePoints(receipt models.Receipt) int {
	logrus.Infof("Checking purchase date for odd day: %s", receipt.PurchaseDate)
	points := 0
	if receipt.PurchaseDate.Day()%2 != 0 {
		points = 6
		logrus.Infof("Purchase day is odd, points: %d", points)
	} else {
		logrus.Infof("Purchase day is not odd")
	}
	return points
}

func (s *ReceiptServiceImpl) CalculatePurchaseTimePoints(receipt models.Receipt) int {
	logrus.Infof("Checking purchase time for between 2:00 PM and 4:00 PM: %s", receipt.PurchaseTime)
	points := 0
	parsedTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		points = 10
		logrus.Infof("Purchase time is between 2:00 PM and 4:00 PM, points: %d", points)
	} else {
		logrus.Infof("Purchase time is not between 2:00 PM and 4:00 PM")
	}
	return points
}

func (s *ReceiptServiceImpl) CalculateTotalPoints(receipt models.Receipt) int {
	logrus.Info("Starting points calculation for receipt")
	points := 0
	points += s.CalculateRetailerPoints(receipt)
	points += s.CalculateTotalIsRoundDollar(receipt)
	points += s.CalculateTotalMultipleOfQuarter(receipt)
	points += s.CalculateItemPoints(receipt)
	points += s.CalculateItemDescriptionPoints(receipt)
	points += s.CalculatePurchaseDatePoints(receipt)
	points += s.CalculatePurchaseTimePoints(receipt)
	logrus.Infof("Total points calculated: %d", points)
	return points
}
