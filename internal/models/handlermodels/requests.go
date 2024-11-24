package handlermodels

// RequestReceipt represents the structure of the incoming request
type RequestReceipt struct {
	Retailer     string        `json:"retailer" validate:"required"`
	Total        string        `json:"total" validate:"required,numeric"`
	PurchaseDate string        `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string        `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []RequestItem `json:"items" validate:"required,min=1,dive"`
}

// RequestItem represents an item in the incoming request
type RequestItem struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,numeric"`
}
