package util

import (
	"testing"

	"gotest.tools/assert"
)

// Test pointer simple functionality, i.e a string
func TestPointer(t *testing.T) {
	val := "testing"
	expected := &val
	res := Pointer(val)

	assert.DeepEqual(t, expected, res)
}

// MORE TESTS FOR DIFFERENT TYPES. Slice, Struct etc..
