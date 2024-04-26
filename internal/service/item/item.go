package item

import (
	"fmt"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/repository"
)

// Item service, containing an item repository
type ItemService struct {
	ItemRepository repository.ItemRepository
}

// Construct a new item service with an item repository
func NewItemService(itemRepository repository.ItemRepository) ItemService {
	return ItemService{
		ItemRepository: itemRepository,
	}
}

// Fetch an item from the item service
func (itemService *ItemService) FetchItem(sku string) (*entity.Item, error) {
	res, err := itemService.ItemRepository.FetchItem(sku)
	if err != nil {
		return nil, fmt.Errorf("fetching item: %s", err)
	}
	return res, nil
}
