package list_double

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleLinkedList_PushFront(t *testing.T) {
	list := NewDoubleLinkedList()

	list.PushFront("value01")

	assert.Equal(t, list.Len(), 1)
	assert.Equal(t, list.isEmpty(), false)
	assert.Equal(t, list.Front(), list.Back())

	list.PushFront("value2")
	assert.Equal(t, list.Len(), 2)
	assert.NotEqual(t, list.Front(), list.Back())
	assert.Equal(t, list.Front(), list.Back().Prev)
	assert.Equal(t, list.Front().Next, list.Back())

	val3 := "value3"
	list.PushFront(val3)
	assert.Equal(t, list.Len(), 3)
	assert.Equal(t, list.Front().Next, list.Back().Prev)
	assert.Equal(t, list.Front().value, val3)
}
func TestDoubleLinkedList_PushBack(t *testing.T) {
	list := NewDoubleLinkedList()

	val1, val2, val3 := "value01", "value02", "value03"
	n1 := list.PushBack(val1)

	assert.Equal(t, list.Len(), 1)
	assert.Equal(t, list.isEmpty(), false)
	assert.Equal(t, list.Front(), list.Back())
	assert.Equal(t, n1, list.Back())
	assert.Equal(t, list.Front(), n1)

	list.PushBack( val2)
	assert.Equal(t, list.Len(), 2)
	assert.NotEqual(t, list.Front(), list.Back())
	assert.Equal(t, list.Front(), list.Back().Prev)
	assert.Equal(t, list.Front().Next, list.Back())

	list.PushBack(val3)
	assert.Equal(t, list.Len(), 3)
	assert.Equal(t, list.Front().Next, list.Back().Prev)
	assert.Equal(t, list.Back().value, val3)
	assert.Equal(t, list.Front().value, val1)
}

func TestDoubleLinkedList_MoveToFront(t *testing.T) {
	list := NewDoubleLinkedList()

	n1 := list.PushFront(1)
	n2 := list.PushFront(2)
	n3 := list.PushFront(3)

	assert.Equal(t, list.Len(), 3)
	assert.NotEqual(t, list.Front(), n2)
	assert.Equal(t, n2.Next, n1)
	assert.Equal(t, n2.Prev, n3)
	list.MoveToFront(n2)

	assert.Equal(t, list.Front(), n2)
	assert.Equal(t, n2.Next, n3)
	assert.Nil(t, n2.Prev)
}

func TestDoubleLinkedList_Remove(t *testing.T) {
	list := NewDoubleLinkedList()

	n1 := list.PushBack(1)
	n2 := list.PushBack(2)
	n3 := list.PushBack(3)

	assert.Equal(t, list.Len(), 3)
	assert.Equal(t, n1.Next, n2)
	assert.Equal(t, n3.Prev, n2)

	list.Remove(n2)

	assert.Equal(t, list.Len(), 2)
	assert.Equal(t, n1.Next, n3)
	assert.Equal(t, n3.Prev, n1)
}