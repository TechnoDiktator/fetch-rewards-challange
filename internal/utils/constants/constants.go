package constants

import "time"

const (
	ProcessReceipts string        = "/receipts/process"
	GetPoints       string        = "/receipts/points/:id"
	PORT            int           = 8080
	TIMEOUT         time.Duration = 20 * time.Second
)
