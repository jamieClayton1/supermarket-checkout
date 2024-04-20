package checkout

import (
	"supermarket-checkout/internal/entity"
	"testing"

	"gotest.tools/assert"
)

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

func TestCalculatePrice(t *testing.T) {
	fetchItemFunc := func(config *entity.FetchItemConfig) (*entity.FetchItemResult, error) {
		return &entity.FetchItemResult{
			Item: &entity.Item{
				SKU:        config.SKU,
				UnitPrice:  10,
				BatchPrice: nil,
			},
		}, nil
	}
	skus := map[entity.SKU]int{
		"A": 1,
	}
	price, err := calculatePrice(skus, fetchItemFunc)
	assert.NilError(t, err)
	assert.Equal(t, 10, price)
}

func TestCalculatePriceMultipleSKUs(t *testing.T) {
	fetchItemFunc := func(config *entity.FetchItemConfig) (*entity.FetchItemResult, error) {
		return &entity.FetchItemResult{
			Item: &entity.Item{
				SKU:        config.SKU,
				UnitPrice:  10,
			},
		}, nil
	}
	skus := map[entity.SKU]int{
		"A": 1,
		"B": 1,
	}
	price, err := calculatePrice(skus, fetchItemFunc)
	assert.NilError(t, err)
	assert.Equal(t, 20, price)
}

func TestCalculatePriceMultipleSKUsDifferentPrices(t *testing.T) {
	fetchItemFunc := func(config *entity.FetchItemConfig) (*entity.FetchItemResult, error) {
		res := &entity.FetchItemResult{
			Item: &entity.Item{
				SKU:        config.SKU,
			},
		}
		if (config.SKU == "A"){
			res.UnitPrice = 10
		} else {
			res.UnitPrice = 20
		}
		return res, nil
	}
	skus := map[entity.SKU]int{
		"A": 1,
		"B": 1,
	}
	price, err := calculatePrice(skus, fetchItemFunc)
	assert.NilError(t, err)
	assert.Equal(t, 30, price)
}

func TestCalculatePriceMultipleSKUsDuplicates(t *testing.T) {
	fetchItemFunc := func(config *entity.FetchItemConfig) (*entity.FetchItemResult, error) {
		res := &entity.FetchItemResult{
			Item: &entity.Item{
				SKU:        config.SKU,
			},
		}
		if (config.SKU == "A"){
			res.UnitPrice = 10
		} else {
			res.UnitPrice = 20
		}
		return res, nil
	}
	skus := map[entity.SKU]int{
		"A": 1,
		"B": 2,
	}
	price, err := calculatePrice(skus, fetchItemFunc)
	assert.NilError(t, err)
	assert.Equal(t, 50, price)
}

func TestCalculatePriceBatchPricing(t *testing.T) {
	fetchItemFunc := func(config *entity.FetchItemConfig) (*entity.FetchItemResult, error) {
		batchSize := 2
		batchPrice := 5
		res := &entity.FetchItemResult{
			Item: &entity.Item{
				SKU:        config.SKU,
				BatchSize: &batchSize,
				BatchPrice: &batchPrice,
			},
		}
		if (config.SKU == "A"){
			res.UnitPrice = 10
		} else {
			res.UnitPrice = 20
		}
		return res, nil
	}
	skus := map[entity.SKU]int{
		"A": 1,
		"B": 5,
	}
	price, err := calculatePrice(skus, fetchItemFunc)
	assert.NilError(t, err)
	assert.Equal(t, 40, price)
}