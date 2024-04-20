package checkout

import (
	"errors"
	"math"
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

// Calculate batch pricing given the regular price, units purchased, 
// special batch pricing and batch sizing to apply the price at
func batchPrice(price int, units int, batchPrice int, batchSize int) int {
	if batchPrice == 0 && batchSize == 0 {
		return price * units
	}
	batches := int(math.Floor(float64(units) / float64(batchSize)))
	regulars := units % batchSize

	return (batches * batchPrice) + (regulars * price)
}