package provider

import (
	"supermarket-checkout/internal/repository/basket"
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
	basketRepository := basket.NewLocalBasketRepository()

	itemService := service.NewItemService(&itemRepository)
	basketService := service.NewBasketService(&basketRepository)
	checkoutService := service.NewCheckoutService(&itemService, &basketService)

	return ServiceProvider{
		CheckoutService: &checkoutService,
	}
}
