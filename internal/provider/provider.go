package provider

import (
	"supermarket-checkout/internal/repository/item"
	"supermarket-checkout/internal/service"
)

// Base service provider
type ServiceProvider struct {
	CheckoutService *service.CheckoutService
}

// Construct a service provider with API configurations
func NewApiServiceProvider() ServiceProvider {
	itemRepository := item.NewLocalItemRepository()
	itemService := service.NewItemService(&itemRepository)
	checkoutService := service.NewCheckoutService(&itemService)

	return ServiceProvider{
		CheckoutService: &checkoutService,
	}
}
