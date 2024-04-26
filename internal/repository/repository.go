package repository

import "supermarket-checkout/internal/entity"

// Item repository interface
type ItemRepository interface {
	FetchItem(sku string) (*entity.Item, error)
}

// Basket repository interface
type BasketRepository interface {
	FetchBasket(basketId string) (*entity.Basket, error)
	AddBasketItem(item *entity.Item, basketId *string) (string, error)
}
