package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/data_structures/heap"
)

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8, 1, 3, 9, 5, 10, 13}
	minHeap := heap.NewMinHeap(len(inputArray))
	maxHeap := heap.NewMaxHeap(len(inputArray))
	for _, v := range inputArray {
		minHeap.Insert(v)
	}
	fmt.Println(minHeap.String())

	for i := 0; i < len(inputArray); i++ {
		fmt.Println(minHeap.Remove())
	}
	fmt.Println(minHeap.String())

	for _, v := range inputArray {
		maxHeap.Insert(v)
	}
	fmt.Println(maxHeap.String())

	for i := 0; i < len(inputArray); i++ {
		fmt.Println(maxHeap.Remove())
	}
	fmt.Println(maxHeap.String())

}
