package item

import (
	"fmt"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/repository"
)

type ItemService struct {
	repository.ItemRepository
}

func NewItemService() ItemService {
	return ItemService{}
}

func (itemService *ItemService) FetchItem(config *FetchItemConfig) (*FetchItemResult, error) {
	res, err := itemService.ItemRepository.FetchItem(&repository.FetchItemConfig{
		SKU: config.SKU,
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching item from item repository: %s", err)
	}
	return &FetchItemResult{
		Item: res.Item,
	}, nil
}

type FetchItemConfig struct {
	entity.SKU
}

type FetchItemResult struct {
	*entity.Item
}
