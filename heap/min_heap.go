package heap

import "fmt"

type minheap struct {
	heapArray []int
	size      int
	maxSize   int
}

func NewMinHeap(maxSize int) *minheap {
	return &minheap{maxSize: maxSize}
}

func (h *minheap) String() string {
	return fmt.Sprintf("Heap size is %d.\n Heap maxSize is %d.\n Heap elements: %v \n", h.size, h.maxSize, h.heapArray)
}

func (h *minheap) isLeaf(idx int) bool {
	return idx >= h.size/2 && idx <= h.size
}

func (h *minheap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *minheap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *minheap) rightChild(idx int) int {
	return idx*2 + 2
}

func (h *minheap) swap(childIdx, parentIdx int) {
	h.heapArray[childIdx], h.heapArray[parentIdx] = h.heapArray[parentIdx], h.heapArray[childIdx]
}

func (h *minheap) upHeapify(idx int) {
	for h.heapArray[idx] < h.heapArray[h.parent(idx)] {
		h.swap(idx, h.parent(idx))
		idx = h.parent(idx)
	}
}

func (h *minheap) downHeapify(currentIdx int) {
	if h.isLeaf(currentIdx) {
		return
	}

	smallestIdx := currentIdx
	leftChildIdx := h.leftChild(currentIdx)
	rightChildIdx := h.rightChild(currentIdx)

	if leftChildIdx < h.size && h.heapArray[smallestIdx] > h.heapArray[leftChildIdx] {
		smallestIdx = leftChildIdx
	}

	if rightChildIdx < h.size && h.heapArray[smallestIdx] > h.heapArray[rightChildIdx] {
		smallestIdx = rightChildIdx
	}

	if smallestIdx != currentIdx {
		h.swap(smallestIdx, currentIdx)
		h.downHeapify(smallestIdx)
	}
}

func (h *minheap) Get() int {
	return h.heapArray[0]
}

func (h *minheap) Insert(el int) error {
	if h.size >= h.maxSize {
		return fmt.Errorf("heap is full")
	}
	h.heapArray = append(h.heapArray, el)
	h.size++
	h.upHeapify(h.size - 1)
	return nil
}

func (h *minheap) Remove() int {
	root := h.heapArray[0]
	h.heapArray[0] = h.heapArray[h.size-1]
	h.heapArray = h.heapArray[:h.size-1]
	h.size--
	h.downHeapify(0)
	return root
}

func (h *minheap) BuildMinHeap() {
	for index := (h.size / 2) - 1; index >= 0; index-- {
		h.downHeapify(index)
	}
}
