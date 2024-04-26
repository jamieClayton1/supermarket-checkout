package item

import (
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/util"
	"testing"

	"gotest.tools/assert"
)

// Test fetching an item with a SKU that does exist
func TestLocalItemRepositoryFetchItem(t *testing.T) {
	repo := NewLocalItemRepository()
	sku := "A"
	expected := &entity.Item{
		SKU:        "A",
		UnitPrice:  50.00,
		BatchSize:  util.Pointer(3),
		BatchPrice: util.Pointer(130),
	}
	res, err := repo.FetchItem(sku)

	assert.NilError(t, err)
	assert.DeepEqual(t, expected, res)
}

// Test fetching an item with a SKU that doesn't exist
func TestLocalItemRepositoryFetchItemThatDoesntExist(t *testing.T) {
	repo := NewLocalItemRepository()
	sku := "F"

	_, err := repo.FetchItem(sku)

	assert.Error(t, err, err.Error())
}
