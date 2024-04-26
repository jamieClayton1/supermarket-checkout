package checkout

import (
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/util"
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
			SKU:        "A",
			UnitPrice:  100,
		},
		{
			SKU:        "A",
			UnitPrice:  100,
		},
		{
			SKU:        "A",
			UnitPrice:  100,
		},
		{
			SKU:        "B",
			UnitPrice:  50,
		},
		{
			SKU:        "B",
			UnitPrice:  50,
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
