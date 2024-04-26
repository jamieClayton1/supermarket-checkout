package basket

import (
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/repository"
)

type BasketService struct {
	BasketRepository repository.BasketRepository
}

func NewBasketService(basketRepository repository.BasketRepository) BasketService {
	return BasketService{
		BasketRepository: basketRepository,
	}
}

func (basketService *BasketService) FetchBasket(basketId string) (*entity.Basket, error) {
	return basketService.BasketRepository.FetchBasket(basketId)
} 

func (basketService *BasketService) AddBasketItem(item *entity.Item, basketId *string) (string, error) {
	return basketService.BasketRepository.PutBasketItem(item, basketId)
} 