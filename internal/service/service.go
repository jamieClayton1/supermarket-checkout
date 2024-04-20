package service

import (
	"supermarket-checkout/internal/service/checkout"
	"supermarket-checkout/internal/service/item"
)

// Export services
type CheckoutService = checkout.CheckoutService
type ItemService = item.ItemService

// Export constructors
var NewCheckoutService = checkout.NewCheckoutService
var NewItemService = item.NewItemService
