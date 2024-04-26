package service

import (
	"supermarket-checkout/internal/service/basket"
	"supermarket-checkout/internal/service/checkout"
	"supermarket-checkout/internal/service/item"
)

// Export services
type CheckoutService = checkout.CheckoutService
type ItemService = item.ItemService
type BasketService = basket.BasketService

// Export constructors
var NewCheckoutService = checkout.NewCheckoutService
var NewItemService = item.NewItemService
var NewBasketService = basket.NewBasketService
