package entity

type SKU = string

type Item struct {
	SKU        SKU
	UnitPrice  int
	BatchSize  *int
	BatchPrice *int
}
