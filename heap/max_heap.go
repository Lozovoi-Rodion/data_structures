package heap

import "fmt"

type maxheap struct {
	heapArray []int
	size      int
	maxSize   int
}

func NewMaxHeap(maxSize int) *maxheap {
	return &maxheap{maxSize: maxSize}
}

func (h *maxheap) String() string {
	return fmt.Sprintf("Heap size is %d.\n Heap maxSize is %d.\n Heap elements: %v \n", h.size, h.maxSize, h.heapArray)
}

func (h *maxheap) isLeaf(idx int) bool {
	return idx >= h.size/2 && idx <= h.size
}

func (h *maxheap) leftChild(idx int) int {
	return idx*2 + 1
}

func (h *maxheap) rightChild(idx int) int {
	return idx*2 + 2
}

func (h *maxheap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *maxheap) swap(first, second int) {
	h.heapArray[first], h.heapArray[second] = h.heapArray[second], h.heapArray[first]
}

func (h *maxheap) upHeapify(idx int) {
	for h.heapArray[idx] > h.heapArray[h.parent(idx)] {
		h.swap(idx, h.parent(idx))
		idx = h.parent(idx)
	}
}

func (h *maxheap) downHeapify(currentIdx int) {
	if h.isLeaf(currentIdx) {
		return
	}

	largest := currentIdx
	leftChildIdx := h.leftChild(currentIdx)
	rightChildIdx := h.rightChild(currentIdx)

	if leftChildIdx < h.size && h.heapArray[leftChildIdx] > h.heapArray[largest] {
		largest = leftChildIdx
	}
	if rightChildIdx < h.size && h.heapArray[rightChildIdx] > h.heapArray[largest] {
		largest = rightChildIdx
	}
	if largest != currentIdx {
		h.swap(currentIdx, largest)
		h.downHeapify(largest)
	}
	return
}

func (h *maxheap) Get() int {
	return h.heapArray[0]
}

func (h *maxheap) Insert(el int) error {
	if h.size >= h.maxSize {
		return fmt.Errorf("heap is full")
	}

	h.heapArray = append(h.heapArray, el)
	h.size++
	h.upHeapify(h.size - 1)
	return nil
}

func (h *maxheap) Remove() int {
	root := h.heapArray[0]
	h.heapArray[0] = h.heapArray[h.size-1]
	h.heapArray = h.heapArray[:h.size-1]
	h.size--
	h.downHeapify(0)
	return root
}

func (h *maxheap) BuildMaxHeap() {
	for index := (h.size / 2) - 1; index >= 0; index-- {
		h.downHeapify(index)
	}
}
