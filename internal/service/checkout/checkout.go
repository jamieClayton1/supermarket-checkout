package checkout

import (
	"fmt"
	"math"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/service/basket"
	"supermarket-checkout/internal/service/item"
)

// Checkout service, containing an Item service
type CheckoutService struct {
	ItemService *item.ItemService
	BasketService *basket.BasketService
}

// Construct a checkout service
func NewCheckoutService(itemService *item.ItemService, basketService *basket.BasketService) CheckoutService {
	return CheckoutService{
		ItemService: itemService,
		BasketService: basketService,
	}
}

func (checkoutService *CheckoutService) ScanItem(sku string, basketId *string) (string, error) {
	item, err := checkoutService.ItemService.FetchItem(sku)
	if err != nil {
		return "", fmt.Errorf("fetching item to scan: %s", err)
	}
	id, err := checkoutService.BasketService.AddBasketItem(item, basketId)
	if err != nil {
		return "", fmt.Errorf("adding item to basket: %s", err)
	}
	return id, nil
}

// Fetch a price from the checkout service
func (checkoutService *CheckoutService) FetchPrice(items []entity.SKU) (price int, err error) {
	price, err = calculatePrice(countSKUs(items), checkoutService.ItemService.FetchItem)
	if err != nil {
		err = fmt.Errorf("fetching price: %s", err)
		return 
	}
	return price, nil
}

// Custom fetch item function type
type FetchItemFunc = func(sku string) (*entity.Item, error)

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
		res, err := fetchItemFunc(sku)
		if err != nil {
			return 0, err
		}
		if res.BatchSize != nil && res.BatchPrice != nil {
			price += batchPrice(res.UnitPrice, units, *res.BatchPrice, *res.BatchSize)
		} else {
			price += res.UnitPrice * units
		}
	}
	return price, nil
}

// Count SKUs from a given slice of SKUs, returning a map with a key of SKU
// and the value of it's corresponding count
func countSKUs(skus []entity.SKU) map[entity.SKU]int {
	items := make(map[entity.SKU]int)
	for _, sku := range skus {
		if count, ok := items[sku]; ok {
			items[sku] = count + 1
		} else {
			items[sku] = 1
		}
	}
	return items
}
