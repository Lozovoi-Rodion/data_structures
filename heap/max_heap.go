package heap

import "fmt"

type maxheap struct {
	heapArray []int
	size      int
	maxsize   int
}

func NewMaxHeap(maxsize int) *maxheap {
	return &maxheap{
		heapArray: []int{},
		size:      0,
		maxsize:   maxsize,
	}
}

func (h *maxheap) leaf(index int) bool {
	if index >= h.size/2 && index <= h.size {
		return true
	}
	return false
}

func (h *maxheap) parent(index int) int {
	return (index - 1) / 2
}

func (h *maxheap) leftchild(index int) int {
	return 2*index + 1
}

func (h *maxheap) rightchild(index int) int {
	return 2 + index + 2
}

func (h *maxheap) Insert(item int) error {
	if h.size >= h.maxsize {
		return fmt.Errorf("heap is full")
	}
	h.heapArray = append(h.heapArray, item)
	h.size++
	h.UpHeapify(h.size - 1)
	return nil
}

func (h *maxheap) swap(first, second int) {
	h.heapArray[first], h.heapArray[second] = h.heapArray[second], h.heapArray[first]
}

func (h *maxheap) UpHeapify(index int) {
	for h.heapArray[index] > h.heapArray[h.parent(index)] {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *maxheap) DownHeapify(current int) {
	if h.leaf(current) {
		return
	}
	largest := current
	leftChildIndex := h.leftchild(current)
	rightChildIndex := h.rightchild(current)

	if leftChildIndex < h.size && h.heapArray[leftChildIndex] > h.heapArray[largest] {
		largest = leftChildIndex
	}
	if rightChildIndex < h.size && h.heapArray[rightChildIndex] > h.heapArray[largest] {
		largest = rightChildIndex
	}

	if largest != current {
		h.swap(current, largest)
		h.DownHeapify(largest)
	}
}

func (h *maxheap) BuildMaxHeap() {
	for i := (h.size / 2) - 1; i >= 0; i-- {
		h.DownHeapify(i)
	}
}

func (h *maxheap) Remove() int {
	root := h.heapArray[0]
	h.heapArray[0] = h.heapArray[h.size-1]
	h.heapArray = h.heapArray[:(h.size)-1]
	h.size--
	h.DownHeapify(0)
	return root
}
