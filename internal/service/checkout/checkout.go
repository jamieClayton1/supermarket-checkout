package checkout

import (
	"errors"
	"fmt"
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

type FetchItemFunc = func(*entity.FetchItemConfig) (*entity.FetchItemResult, error)

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

// Calculate price of a map of SKU counts, given a map of SKUs and 
// a function to fetch the relevant Item data
func calculatePrice(skus map[entity.SKU]int, fetchItemFunc FetchItemFunc) (int, error) {
	var price int
	for sku, units := range skus {
		item, err := fetchItemFunc(&entity.FetchItemConfig{
			SKU: sku,
		})
		if err != nil {
			return 0, fmt.Errorf("error when pricing items: %s", err)
		}
		if item.BatchSize != nil && item.BatchPrice != nil {
			price += batchPrice(item.UnitPrice, units, *item.BatchPrice, *item.BatchSize)
		} else {
			price += item.UnitPrice * units
		}
	}
	return price, nil
}