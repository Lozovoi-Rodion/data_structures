package heap

import "fmt"

type minheap struct {
	heapArray []int
	size      int
	maxsize   int
}

func (m *minheap) String() string {
	return fmt.Sprintf("Heap size is %d.\n Heap maxSize is %d.\n Heap elements: %v \n", m.size, m.maxsize, m.heapArray)
}

func NewMinHeap(maxsize int) *minheap {
	return &minheap{maxsize: maxsize}
}

func (m *minheap) isLeaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *minheap) parent(index int) int {
	return (index - 1) / 2
}

func (m *minheap) leftChild(index int) int {
	return 2*index + 1
}

func (m *minheap) rightChild(index int) int {
	return 2*index + 2
}

func (m *minheap) Insert(item int) error {
	if m.size >= m.maxsize {
		return fmt.Errorf("heap is full")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}

func (m *minheap) Remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:m.size-1]
	m.size--
	m.downHeapify(0)
	return top
}

func (m *minheap) swap(first, second int) {
	m.heapArray[first], m.heapArray[second] = m.heapArray[second], m.heapArray[first]
}

func (m *minheap) upHeapify(index int) {
	for m.heapArray[index] < m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *minheap) downHeapify(current int) {
	if m.isLeaf(current) {
		return
	}

	smallest := current
	leftChildIndex := m.leftChild(current)
	rightChildIndex := m.rightChild(current)

	if leftChildIndex < m.size && m.heapArray[leftChildIndex] < m.heapArray[smallest] {
		smallest = leftChildIndex
	}
	if rightChildIndex < m.size && m.heapArray[rightChildIndex] < m.heapArray[smallest] {
		smallest = rightChildIndex
	}

	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest)
	}
}

func (m *minheap) BuildMinHeap() {
	for index := m.size/2 - 1; index >= 0; index-- {
		m.downHeapify(index)
	}
}
