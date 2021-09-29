package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/data_structures/heap"
)

func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	minHeap := heap.NewMinHeap(len(inputArray))
	for _, v := range inputArray {
		minHeap.Insert(v)
	}
	fmt.Println(minHeap)
	fmt.Println("----DELIMETER----")
	minHeap.BuildMinHeap()
	fmt.Println(minHeap)
	fmt.Println("----DELIMETER----")
	fmt.Println(minHeap.Remove())
	fmt.Println("----DELIMETER----")
	fmt.Println(minHeap)
	fmt.Println(minHeap.Remove())
	fmt.Println("----DELIMETER----")
	fmt.Println(minHeap)

}
