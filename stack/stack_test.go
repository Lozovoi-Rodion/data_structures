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

	tests := []struct {
		el     string
		result bool
	}{
		{"c", true},
		{"b", true},
		{"a", true},
		{"", false},
	}

	for _, tc := range tests {
		tc := tc
		el, ok := stack.Pop()
		assert.Equal(t, el, tc.el)
		assert.Equal(t, ok, tc.result)
	}
}
