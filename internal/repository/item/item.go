package item

import (
	"errors"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/repository"
	"supermarket-checkout/internal/util"
)

// Local item repository implementation
type LocalItemRepository struct {
	store map[entity.SKU]entity.Item
}

// Construct a new local item repository with it's item store
func NewLocalItemRepository() LocalItemRepository {
	return LocalItemRepository{
		map[entity.SKU]entity.Item{
			"A": {
				SKU:        "A",
				UnitPrice:  50.00,
				BatchSize:  util.Pointer(3),
				BatchPrice: util.Pointer(130),
			},
			"B": {
				SKU:        "B",
				UnitPrice:  30.00,
				BatchSize:  util.Pointer(2),
				BatchPrice: util.Pointer(45),
			},
			"C": {
				SKU:       "C",
				UnitPrice: 20.00,
			},
			"D": {
				SKU:       "D",
				UnitPrice: 15.00,
			},
		},
	}
}

// Fetch an item from the local item repository
func (itemRepository *LocalItemRepository) FetchItem(config *repository.FetchItemConfig) (*repository.FetchItemResult, error) {
	if item, ok := itemRepository.store[config.SKU]; ok {
		return &repository.FetchItemResult{
			Item: &item,
		}, nil
	}
	return nil, errors.New("no item found with that SKU")
}
