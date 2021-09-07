package queue

import (
	"container/list"
)

type iQueue interface {
	Enqueue(el interface{})
	Dequeue() interface{}
}

type Queue []interface{}

func (q *Queue) Enqueue(el interface{}) {
	*q = append(*q, el)
}

func (q *Queue) Dequeue() interface{} {
	el := (*q)[0]
	*q = (*q)[1:]
	return el
}

type QueueOptimized struct {
	list.List
}

func (q *QueueOptimized) Enqueue(el interface{}) {
	q.PushBack(el)
}

func (q *QueueOptimized) Dequeue() interface{} {
	el := q.Front()
	q.Remove(el)
	return el.Value
}
