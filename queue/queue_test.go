package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	for _, v := range []int{10, 20, 30} {
		el := queue.Dequeue()
		assert.Equal(t, el, v)
	}
}

func TestQueueOptimized(t *testing.T) {
	queue := QueueOptimized{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	for _, v := range []int{10, 20, 30} {
		el := queue.Dequeue()
		assert.Equal(t, el, v)
	}
}
