package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	el, ok := stack.Pop()
	assert.Equal(t, el, "c")
	assert.Equal(t, ok, true)

	el, ok = stack.Pop()
	assert.Equal(t, el, "b")
	assert.Equal(t, ok, true)

	el, ok = stack.Pop()
	assert.Equal(t, el, "a")
	assert.Equal(t, ok, true)

	el, ok = stack.Pop()
	assert.Equal(t, el, "")
	assert.Equal(t, ok, false)
}
