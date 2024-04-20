package entity

type SKU = string

type Item struct {
	SKU        SKU
	UnitPrice  int
	BatchSize  *int
	BatchPrice *int
}

type FetchItemConfig struct {
	SKU
}

type FetchItemResult struct {
	*Item
}
