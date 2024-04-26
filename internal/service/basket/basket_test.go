package basket

import (
	"fmt"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/repository/mock"
	"testing"

	"go.uber.org/mock/gomock"
	"gotest.tools/assert"
)

func TestBasketService_FetchBasket_ValidId(t *testing.T) {
	basket := entity.Basket{
		Items: []*entity.Item{},
	}
	basketId := "valid"
	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().FetchBasket(gomock.Eq(basketId)).Return(&basket, nil)
	service := NewBasketService(basketRepository)

	expected := basket
	response, err := service.FetchBasket(basketId)

	assert.NilError(t, err)
	assert.DeepEqual(t, expected, *response)
}

func TestBasketService_FetchBasket_InvalidId(t *testing.T) {
	expectedErr := fmt.Errorf("error, invalid")
	basketId := "invalid"
	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().FetchBasket(gomock.Eq(basketId)).Return(nil, expectedErr)
	service := NewBasketService(basketRepository)

	_, err := service.FetchBasket(basketId)
	assert.Error(t, err, expectedErr.Error())
}

func TestBasketService_AddBasketItem_ItemAndIdNotNil(t *testing.T) {
	basketId := "valid"
	item := entity.Item{}
	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().PutBasketItem(gomock.Eq(&item), gomock.Eq(&basketId)).Return(basketId, nil)
	service := NewBasketService(basketRepository)

	expected := basketId
	response, err := service.AddBasketItem(&item, &basketId)
	
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, response)
}

func TestBasketService_AddBasketItem_IdNilGeneratesId(t *testing.T) {
	basketId := "new"
	item := entity.Item{}
	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().PutBasketItem(gomock.Eq(&item), gomock.Nil()).Return(basketId, nil)
	service := NewBasketService(basketRepository)

	expected := basketId
	response, err := service.AddBasketItem(&item, nil)
	
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, response)
}

func TestBasketService_AddBasketItem_ItemNilErrors(t *testing.T) {
	expectedErr := fmt.Errorf("no item")
	basketId := "valid"
	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().PutBasketItem(gomock.Nil(), gomock.Eq(&basketId)).Return(basketId, expectedErr)
	service := NewBasketService(basketRepository)

	_, err := service.AddBasketItem(nil, &basketId)

	assert.Error(t, err, expectedErr.Error())
}

func TestBasketService_AddBasketItem_ErrWithRepoErr(t *testing.T) {
	expectedErr := fmt.Errorf("any err")
	basketId := "valid"
	item := entity.Item{}
	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().PutBasketItem(gomock.Eq(&item), gomock.Eq(&basketId)).Return(basketId, expectedErr)
	service := NewBasketService(basketRepository)

	_, err := service.AddBasketItem(&item, &basketId)
	
	assert.Error(t, err, expectedErr.Error())
}