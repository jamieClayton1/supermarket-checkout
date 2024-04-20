package checkout

import "testing"

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
