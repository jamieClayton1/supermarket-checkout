package item

import (
	"errors"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/repository/mock"
	"testing"

	"go.uber.org/mock/gomock"
	"gotest.tools/assert"
)

// When a valid basket if is provided, a basket is returned
func TestItemService_FetchItem_ValidId(t *testing.T) {
	item := entity.Item{}
	sku := "A"
	itemRepository := mock.NewMockItemRepository(gomock.NewController(t))
	itemRepository.EXPECT().FetchItem(gomock.Eq(sku)).Return(&item, nil)
	service := NewItemService(itemRepository)

	expected := item
	response, err := service.FetchItem(sku)

	assert.NilError(t, err)
	assert.DeepEqual(t, expected, *response)
}

// When an invalid id is provided, an error is returned
func TestItemService_FetchItem_InvalidId(t *testing.T) {
	sku := "A"
	itemRepository := mock.NewMockItemRepository(gomock.NewController(t))
	itemRepository.EXPECT().FetchItem(gomock.Eq(sku)).Return(nil, errors.New("test error"))
	service := NewItemService(itemRepository)
	_, err := service.FetchItem(sku)

	assert.Error(t, err, err.Error())
}
