package repository

import "supermarket-checkout/internal/entity"

// Item repository interface
type ItemRepository interface {
	ItemFetcher
}

// Interface for fetching an Item
type ItemFetcher interface {
	FetchItem(*FetchItemConfig) (*FetchItemResult, error)
}

// DTO for the configuration of a fetch item function
type FetchItemConfig struct {
	entity.SKU
}

// DTO for the resulting value of afetch item function
type FetchItemResult struct {
	*entity.Item
}
