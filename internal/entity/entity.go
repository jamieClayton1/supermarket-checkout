package entity

// Custom SKU type
type SKU = string

// A checkout item, containing SKU, price and batch size/price configuration
type Item struct {
	SKU        SKU
	UnitPrice  int
	BatchSize  *int
	BatchPrice *int
}

type Basket struct {
	Items []*Item
}

// Request body representation of a fetch checkout price request
type FetchCheckoutPriceRequest struct {
	BasketId string `json:"basket_id"`
}

// Request body representation of a scan item request
type ScanItemRequest struct {
	BasketId *string `json:"basket_id"`
	SKU string `json:"sku"`
}
