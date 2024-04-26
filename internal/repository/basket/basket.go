package basket

import (
	"fmt"
	"supermarket-checkout/internal/entity"

	"github.com/google/uuid"
)

type LocalBasketRepository struct {
	store map[string]entity.Basket
}

func NewLocalBasketRepository() LocalBasketRepository {
	return LocalBasketRepository{
		store: make(map[string]entity.Basket),
	}
}

func (basketRepository *LocalBasketRepository) FetchBasket(basketId string) (*entity.Basket, error) {
	if basket, ok := basketRepository.store[basketId]; ok {
		return &basket, nil
	} else {
		return nil, fmt.Errorf("no basket %s exists", basketId)
	}
}

func (basketRepository *LocalBasketRepository) PutBasketItem(item *entity.Item, basketId *string) (string, error){
	if item == nil {
		return "", fmt.Errorf("nil item provided")
	}
	if basketId == nil {
		newBasketId := uuid.New().String()
		basketId = &newBasketId
	}
	if basket, ok := basketRepository.store[*basketId]; ok {
		newBasket := basket
		newBasket.Items = append(basket.Items, item)
		basketRepository.store[*basketId] = newBasket
	} else {
		basketRepository.store[*basketId] = entity.Basket{
			Items: []*entity.Item{item},
		}
	}
	return *basketId, nil
}