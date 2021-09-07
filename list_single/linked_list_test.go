package list_single

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	list := NewLinkedList()
	n1 := NewNode("node_one")
	list.Add(n1)

	assert.Equal(t, list.Len(), 1)
	assert.Equal(t, list.GetHead(), n1)
	assert.Equal(t, list.GetTail(), n1)

	n2 := NewNode("node_two")
	list.Add(n2)

	assert.Equal(t, list.Len(), 2)
	assert.Equal(t, list.GetHead(), n1)
	assert.Equal(t, list.GetTail(), n2)
}

func TestLinkedList_Insert(t *testing.T) {
	list := NewLinkedList()
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	list.Add(n1)
	ok, err := list.Insert(0, n2)
	ok, err = list.Insert(1, n3)

	assert.Equal(t, list.Len(), 3)
	assert.Equal(t, ok, true)
	assert.Equal(t, err, nil)
	assert.Equal(t, list.GetNode(0), n2)
	assert.Equal(t, list.GetNode(1), n3)
	assert.Equal(t, list.GetNode(2), n1)

	ok, err = list.Insert(10, NewNode(4))

	assert.Equal(t, ok, false)
	assert.EqualError(t, err, "out of boundary")
	assert.Equal(t, list.Len(), 3)
}

func TestLinkedList_Delete(t *testing.T) {
	list := NewLinkedList()
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)

	list.Add(n1)
	list.Add(n2)
	list.Add(n3)

	assert.Equal(t, list.Len(), 3)

	ok, err := list.Delete(n2)

	assert.Equal(t, ok, true)
	assert.Equal(t, err, nil)
	assert.Equal(t, list.Len(), 2)

	ok, err = list.Delete(n4)
	assert.Equal(t, ok, false)
	assert.EqualError(t, err, "node not found")
	assert.Equal(t, list.Len(), 2)
}