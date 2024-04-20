package provider

import (
	"supermarket-checkout/internal/repository/item"
	"supermarket-checkout/internal/service"
)

type ServiceProvider struct {
	CheckoutService *service.CheckoutService
}

func NewApiServiceProvider() ServiceProvider {
	itemRepository := item.NewLocalItemRepository()
	itemService := service.NewItemService(&itemRepository)
	checkoutService := service.NewCheckoutService(&itemService)

	return ServiceProvider{
		CheckoutService: &checkoutService,
	}
}