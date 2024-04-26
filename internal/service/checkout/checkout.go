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

// Scan an item to an existing basket if a basketId is provided, creating a new basket and returning the basketId if not 
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
func (checkoutService *CheckoutService) FetchPrice(basketId string) (price int, err error) {
	basket, err := checkoutService.BasketService.FetchBasket(basketId)
	if err != nil {
		return
	}
	price, err = calculatePrice(basket.Items)
	if err != nil {
		err = fmt.Errorf("fetching price: %s", err)
		return 
	}
	return
}

// Calculate batch pricing given the regular price, units purchased,
// special batch pricing and batch sizing to apply the price at
func batchPrice(price int, units int, batchPrice int, batchSize int) int {
	// Handle empty values
	if batchPrice == 0 && batchSize == 0 {
		return price * units
	}
	
	// Calculate batch & regular unit counts
	batches := int(math.Floor(float64(units) / float64(batchSize)))
	regulars := units % batchSize

	// Return price for batches plus price for regular units
	return (batches * batchPrice) + (regulars * price)
}

// Calculate price of a slice of items
func calculatePrice(items []*entity.Item) (int, error) {
	var price int
	counts := make(map[entity.SKU]int)
	unq := make([]*entity.Item, 0)

	// Count items & generate unique slice of Items
	for _, item := range items {
		if count, ok := counts[item.SKU]; ok {
			counts[item.SKU] = count + 1
		} else {
			counts[item.SKU] = 1
			unq = append(unq, item)
		}
	}

	// Iterate over unique items, generating either batch price or price
	for _, item := range unq {
		units := counts[item.SKU]
		if item.BatchSize != nil && item.BatchPrice != nil {
			price += batchPrice(item.UnitPrice, units, *item.BatchPrice, *item.BatchSize)
		} else {
			price += item.UnitPrice * units
		}
	}

	return price, nil
}
