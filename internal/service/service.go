package service

import (
	"supermarket-checkout/internal/service/checkout"
	"supermarket-checkout/internal/service/item"
)

type CheckoutService = checkout.CheckoutService
type ItemService = item.ItemService

var NewCheckoutService = checkout.NewCheckoutService
var NewItemService = item.NewItemService