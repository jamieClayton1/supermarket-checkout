package checkout

import (
	"errors"
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/repository/mock"
	"supermarket-checkout/internal/service/basket"
	"supermarket-checkout/internal/service/item"
	"supermarket-checkout/internal/util"
	"testing"

	"go.uber.org/mock/gomock"
	"gotest.tools/assert"
)

// When we scan an item, no error is returned and a basketId is returned
func TestCheckoutService_ScanItem_SKUFoundAndIdNotNil(t *testing.T) {
	basketId := "valid"
	sku := "A"
	i := entity.Item{}

	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().AddBasketItem(gomock.Eq(&i), gomock.Eq(&basketId)).Return(basketId, nil)
	basketService := basket.NewBasketService(basketRepository)

	itemRepository := mock.NewMockItemRepository(gomock.NewController(t))
	itemRepository.EXPECT().FetchItem(gomock.Eq(sku)).Return(&i, nil)
	itemService := item.NewItemService(itemRepository)

	checkoutService := NewCheckoutService(&itemService, &basketService)

	expected := basketId
	response, err := checkoutService.ScanItem(sku, &basketId)

	assert.NilError(t, err)
	assert.DeepEqual(t, expected, response)
}

// When we scan an item with no basketId, no error and a basketId is returned
func TestCheckoutService_ScanItem_IdNil(t *testing.T) {
	basketId := "new"
	sku := "A"
	i := entity.Item{}

	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().AddBasketItem(gomock.Eq(&i), nil).Return(basketId, nil)
	basketService := basket.NewBasketService(basketRepository)

	itemRepository := mock.NewMockItemRepository(gomock.NewController(t))
	itemRepository.EXPECT().FetchItem(gomock.Eq(sku)).Return(&i, nil)
	itemService := item.NewItemService(itemRepository)

	checkoutService := NewCheckoutService(&itemService, &basketService)

	expected := basketId
	response, err := checkoutService.ScanItem(sku, nil)

	assert.NilError(t, err)
	assert.DeepEqual(t, expected, response)
}

// When we scan an item with an invalid SKU, an error is returned
func TestCheckoutService_ScanItem_SkuNotFoundErrors(t *testing.T) {
	expectedErr := errors.New("test error")
	basketId := "new"
	sku := "A"

	itemRepository := mock.NewMockItemRepository(gomock.NewController(t))
	itemRepository.EXPECT().FetchItem(gomock.Eq(sku)).Return(nil, expectedErr)
	itemService := item.NewItemService(itemRepository)

	checkoutService := NewCheckoutService(&itemService, nil)
	_, err := checkoutService.ScanItem(sku, &basketId)

	assert.Error(t, err, err.Error())
}

// When we fetch the parice of a basket with a valid basket id, the price is returned
func TestCheckoutService_FetchPrice_ValidBasketId(t *testing.T) {
	basketId := "valid"
	bskt := entity.Basket{
		Items: []*entity.Item{
			{
				SKU:       "A",
				UnitPrice: 100,
			},
		},
	}

	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().FetchBasket(gomock.Eq(basketId)).Return(&bskt, nil)
	basketService := basket.NewBasketService(basketRepository)

	checkoutService := NewCheckoutService(nil, &basketService)

	expected := 100
	response, err := checkoutService.FetchPrice(basketId)

	assert.NilError(t, err)
	assert.DeepEqual(t, expected, response)
}

// When we fetch the price of an invalid basket id, an error is returned
func TestCheckoutService_FetchPrice_InvalidBasketIdErrors(t *testing.T) {
	basketId := "invalid"

	basketRepository := mock.NewMockBasketRepository(gomock.NewController(t))
	basketRepository.EXPECT().FetchBasket(gomock.Eq(basketId)).Return(nil, errors.New("test error"))
	basketService := basket.NewBasketService(basketRepository)

	checkoutService := NewCheckoutService(nil, &basketService)

	_, err := checkoutService.FetchPrice(basketId)

	assert.Error(t, err, err.Error())
}

// Calculates the correct price for a batch of items with a batch price and batch size
func TestBatchPriceWithBatchItems(t *testing.T) {
	price := batchPrice(10, 20, 50, 5)
	expected := 200
	if price != expected {
		t.Errorf("Incorrect price, got: %d, want: %d", price, expected)
	}
}

// Calculates the correct price for a batch of items with a batch price and batch size, when there are regular items left over
func TestBatchPriceWithBatchItemsAndRegularItemsLeft(t *testing.T) {
	price := batchPrice(10, 20, 50, 5)
	expected := 200
	if price != expected {
		t.Errorf("Incorrect price, got: %d, want: %d", price, expected)
	}
}

// Calculates the correct price for a batch of items with a batch price and batch size, when there are no regular items left over
func TestBatchPriceWithNoRegularItems(t *testing.T) {
	price := batchPrice(10, 20, 50, 20)
	expected := 50
	if price != expected {
		t.Errorf("Incorrect price, got: %d, want: %d", price, expected)
	}
}

// Calculates the correct price for a single item without a batch price or batch size
func TestBatchPriceWithoutBatch(t *testing.T) {
	price := batchPrice(10, 1, 0, 0)
	expected := 10
	if price != expected {
		t.Errorf("Incorrect price, got: %d, want: %d", price, expected)
	}
}

// Calculates the correct price for multiple items with batch pricing
func TestCalculatePrice_MultipleItemsWithBatchPricing(t *testing.T) {
	// Arrange
	items := []*entity.Item{
		{
			SKU:        "A",
			UnitPrice:  100,
			BatchSize:  util.Pointer(3),
			BatchPrice: util.Pointer(250),
		},
		{
			SKU:        "A",
			UnitPrice:  100,
			BatchSize:  util.Pointer(3),
			BatchPrice: util.Pointer(250),
		},
		{
			SKU:        "A",
			UnitPrice:  100,
			BatchSize:  util.Pointer(3),
			BatchPrice: util.Pointer(250),
		},
		{
			SKU:        "B",
			UnitPrice:  50,
			BatchSize:  util.Pointer(2),
			BatchPrice: util.Pointer(80),
		},
		{
			SKU:        "B",
			UnitPrice:  50,
			BatchSize:  util.Pointer(2),
			BatchPrice: util.Pointer(80),
		},
		{
			SKU:       "C",
			UnitPrice: 30,
		},
	}

	// Act
	price, err := calculatePrice(items)

	// Assert
	assert.NilError(t, err)
	assert.Equal(t, 360, price)
}

// Calculates the correct price for multiple items with no batch pricing
func TestCalculatePrice_MultipleItemsNoBatchPricing(t *testing.T) {
	// Arrange
	items := []*entity.Item{
		{
			SKU:       "A",
			UnitPrice: 100,
		},
		{
			SKU:       "A",
			UnitPrice: 100,
		},
		{
			SKU:       "A",
			UnitPrice: 100,
		},
		{
			SKU:       "B",
			UnitPrice: 50,
		},
		{
			SKU:       "B",
			UnitPrice: 50,
		},
		{
			SKU:       "C",
			UnitPrice: 30,
		},
	}

	// Act
	price, err := calculatePrice(items)

	// Assert
	assert.NilError(t, err)
	assert.Equal(t, 430, price)
}

// Calculates the correct price for a single item with no batch pricing
func TestCalculatePrice_SingleItemNoBatchPricing(t *testing.T) {
	// Arrange
	items := []*entity.Item{
		{
			SKU:       "A",
			UnitPrice: 100,
		},
	}

	// Act
	price, err := calculatePrice(items)

	// Assert
	assert.NilError(t, err)
	assert.Equal(t, 100, price)
}

// Calculates the correct price for a single item with batch pricing
func TestCalculatePrice_SingleItemWithBatchPricing(t *testing.T) {
	// Arrange
	items := []*entity.Item{
		{
			SKU:        "A",
			UnitPrice:  100,
			BatchSize:  util.Pointer(1),
			BatchPrice: util.Pointer(400),
		},
	}

	// Act
	price, err := calculatePrice(items)

	// Assert
	assert.NilError(t, err)
	assert.Equal(t, 400, price)
}

// Handles items with a zero unit price correctly
func TestCalculatePrice_ZeroUnitPrice(t *testing.T) {
	// Arrange
	items := []*entity.Item{
		{
			SKU:       "A",
			UnitPrice: 0,
		},
	}

	// Act
	price, err := calculatePrice(items)

	// Assert
	assert.NilError(t, err)
	assert.Equal(t, 0, price)
}
