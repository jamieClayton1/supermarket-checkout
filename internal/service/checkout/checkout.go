package checkout

import (
	"errors"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/service/item"
)

type CheckoutService struct {
	ItemService *item.ItemService
}

func NewCheckoutService(itemService *item.ItemService) CheckoutService {
	return CheckoutService{
		ItemService: itemService,
	}
}

func (checkoutService *CheckoutService) FetchPrice(config *FetchPriceConfig) (*FetchPriceResult, error) {
	return nil, errors.New("method not implemented")
}

type FetchPriceConfig struct {
	ItemSKUs []entity.SKU
}

type FetchPriceResult struct {
	Price int
}
