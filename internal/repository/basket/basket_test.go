package basket

import (
	"supermarket-checkout/internal/entity"
	"supermarket-checkout/internal/util"
	"testing"

	"gotest.tools/assert"
)

// When we put an item into the basket, we can fetch it
func TestLocalBasketRepository_PutBasketFetchBasket(t *testing.T) {
	repo := NewLocalBasketRepository()
	item := entity.Item{
		SKU:        "A",
		UnitPrice:  50.00,
		BatchSize:  util.Pointer(3),
		BatchPrice: util.Pointer(130),
	}
	basketId, err := repo.PutBasketItem(&item, nil)
	assert.NilError(t, err)

	expected := &entity.Basket{
		Items: []*entity.Item{&item},
	}
	res, err := repo.FetchBasket(basketId)
	assert.NilError(t, err)
	assert.DeepEqual(t, expected, res)
}

// When basket doesn't exist, an error is returned
func TestLocalBasketRepository_FetchBasket_NotExistErrors(t *testing.T) {
	repo := NewLocalBasketRepository()
	_, err := repo.FetchBasket("invalid")

	assert.Error(t, err, err.Error())
}

// When a nil item is provided, an error is returned
func TestLocalBasketRepository_PutBasketItem_NilItemErrors(t *testing.T) {
	repo := NewLocalBasketRepository()
	_, err := repo.PutBasketItem(nil, nil)

	assert.Error(t, err, err.Error())
}

// When a nil basketId is provided, we generate a new one
func TestLocalBasketRepository_PutBasketItem_NilIdGeneratesId(t *testing.T) {
	repo := NewLocalBasketRepository()
	item := entity.Item{
		SKU:        "A",
		UnitPrice:  50.00,
		BatchSize:  util.Pointer(3),
		BatchPrice: util.Pointer(130),
	}
	id, err := repo.PutBasketItem(&item, nil)

	assert.NilError(t, err)
	if id == "" {
		t.Error("empty id returned")
	}

}
