package constants

import "time"

const (
	ProcessReceipts string        = "/receipts/process"
	GetPoints       string        = "/receipts/points/:id"
	PORT            string        = "localhost:4040"
	TIMEOUT         time.Duration = 20 * time.Second
)
