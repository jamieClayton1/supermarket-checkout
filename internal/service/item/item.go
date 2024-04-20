package item

import (
	"fmt"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/repository"
)

type ItemService struct {
	ItemRepository repository.ItemRepository
}

func NewItemService(itemRepository repository.ItemRepository) ItemService {
	return ItemService{
		ItemRepository: itemRepository,
	}
}

func (itemService *ItemService) FetchItem(config *entity.FetchItemConfig) (*entity.FetchItemResult, error) {
	res, err := itemService.ItemRepository.FetchItem(&repository.FetchItemConfig{
		SKU: config.SKU,
	})
	if err != nil {
		return nil, fmt.Errorf("fetching item: %s", err)
	}
	return &entity.FetchItemResult{
		Item: res.Item,
	}, nil
}
