package checkout

import (
	"errors"
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

// Calculates the correct price for a single item with no batch price or size
func TestCalculatePrice(t *testing.T) {
	fetchItemFunc := func(sku string) (*entity.Item, error) {
		return &entity.Item{
				SKU:       sku,
				UnitPrice: 10,
		}, nil
	}
	skus := map[entity.SKU]int{
		"A": 1,
	}
	price, err := calculatePrice(skus, fetchItemFunc)
	assert.NilError(t, err)
	assert.Equal(t, 10, price)
}

// Calculates the correct price for more than one item with no batch price or size
func TestCalculatePriceMultipleSKUs(t *testing.T) {
	fetchItemFunc := func(sku string) (*entity.Item, error) {
		return &entity.Item{
			SKU:       sku,
			UnitPrice: 10,
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

// Calculates the correct price for more than one item with different prices with no batch price or size
func TestCalculatePriceMultipleSKUsDifferentPrices(t *testing.T) {
	fetchItemFunc := func(sku string) (*entity.Item, error) {
		res := &entity.Item{
			SKU: sku,
		}
		if sku == "A" {
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

// Calculates the correct price for more than one item with different price, duplicates and
// with no batch price or size
func TestCalculatePriceMultipleSKUsDuplicates(t *testing.T) {
	fetchItemFunc := func(sku string) (*entity.Item, error) {
		res := &entity.Item{
			SKU: sku,
		}
		if sku == "A" {
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

// Calculates the correct price for more than one item with different prices & batch pricing
func TestCalculatePriceBatchPricing(t *testing.T) {
	fetchItemFunc := func(sku string) (*entity.Item, error) {
		batchSize := 2
		batchPrice := 5
		res := &entity.Item{
			SKU:        sku,
			BatchSize:  &batchSize,
			BatchPrice: &batchPrice,
		}
		if sku == "A" {
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

// An error is returned when the fetch item func also returns an error
func TestCalculatePriceFetchItemError(t *testing.T) {
	fetchItemFunc := func(sku string) (*entity.Item, error) {
		return nil, errors.New("we returned an error")
	}
	skus := map[entity.SKU]int{
		"A": 1,
		"B": 5,
	}
	_, err := calculatePrice(skus, fetchItemFunc)
	assert.Error(t, err, err.Error())
}

// Count the SKUs from a given list of none duplicated SKUs
func TestCountSKUs(t *testing.T) {
	skus := []entity.SKU{"A", "B", "C"}
	expected := map[entity.SKU]int{
		"A": 1,
		"B": 1,
		"C": 1,
	}
	res := countSKUs(skus)

	assert.DeepEqual(t, expected, res)
}

// Given a list of SKUs with duplicates, the function should return a map with each SKU as a key and the value set to the number of occurrences in the list.
func TestCountSKUs_WithDuplicates(t *testing.T) {
	skus := []entity.SKU{"A", "B", "A", "C", "B", "A"}
	expected := map[entity.SKU]int{
		"A": 3,
		"B": 2,
		"C": 1,
	}
	result := countSKUs(skus)
	assert.DeepEqual(t, expected, result)
}

// Given an empty list of SKUs, the function should return an empty map.
func TestCountSKUs_EmptyList(t *testing.T) {
	skus := []entity.SKU{}
	expected := map[entity.SKU]int{}
	result := countSKUs(skus)
	assert.DeepEqual(t, expected, result)
}
