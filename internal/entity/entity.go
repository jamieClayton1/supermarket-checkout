package entity

type SKU = string

type Item struct {
	SKU        SKU
	UnitPrice  int
	BatchSize  *int
	BatchPrice *int
}

type FetchItemConfig struct {
	SKU SKU
}

type FetchItemResult struct {
	Item *Item
}

type FetchPriceConfig struct {
	ItemSKUs []SKU
}

type FetchPriceResult struct {
	Price int
}

type FetchCheckoutPriceRequest struct {
	ItemSKUs []SKU `json:"item_skus"`
}
