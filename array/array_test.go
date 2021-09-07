package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_Append(t *testing.T) {
	arr := NewArray(1, 2, 3, 4, 5)

	assert.Equal(t, len(arr), 5)

	arr.Append(6)
	arr.Append(7)

	assert.Equal(t, len(arr), 7)
}

func TestArray_Delete(t *testing.T) {
	arr := NewArray(1,2,3)

	assert.Equal(t, len(arr), 3)

	arr.Delete(0)
	ok, err := arr.Delete(1)

	assert.Equal(t, ok, true)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(arr), 1)

	ok, err = arr.Delete(10)
	assert.Equal(t, ok, false)
	assert.EqualError(t, err, "index exceeds array size")
	assert.Equal(t, len(arr), 1)
}

func TestArray_Insert(t *testing.T) {
	arr := NewArray(1,2,3)

	ok, err := arr.Insert(2, 4)

	assert.Equal(t, ok, true)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(arr), 4)
	assert.Equal(t, arr, Array{1,2,4,3})

	ok, err = arr.Insert(10, 5)
	assert.Equal(t, ok, false)
	assert.EqualError(t, err, "out of boundary")
	assert.Equal(t, len(arr), 4)
}