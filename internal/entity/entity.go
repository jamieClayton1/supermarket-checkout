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

// DTO for providing configuration to a fetch item function
type FetchItemConfig struct {
	SKU SKU
}

// DTO for the resulting value from a fetch item function
type FetchItemResult struct {
	Item *Item
}

// DTO for providing confiuguration to a fetch price function
type FetchPriceConfig struct {
	ItemSKUs []SKU
}

// DTO for the resulting value from a fetch price function
type FetchPriceResult struct {
	Price int
}

// Request body representation of a fetch checkout price request
type FetchCheckoutPriceRequest struct {
	ItemSKUs []SKU `json:"item_skus"`
}
