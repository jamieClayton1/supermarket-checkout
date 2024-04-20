package repository

import "supermarket-checkout/internal/entity"

type ItemRepository interface {
	ItemFetcher
}

type ItemFetcher interface {
	FetchItem(*FetchItemConfig) (*FetchItemResult, error)
}

type FetchItemConfig struct {
	entity.SKU
}

type FetchItemResult struct {
	*entity.Item
}
