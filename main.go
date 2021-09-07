package main

import (
	"fmt"
	"github.com/Lozovoi-Rodion/data_structures/array"
	"github.com/Lozovoi-Rodion/data_structures/list_single"
	stack2 "github.com/Lozovoi-Rodion/data_structures/stack"
)

func main() {
	arr := array.Array{1, 2, 3}
	arr.Append(3)

	stack := stack2.Stack{"a", "b", "c"}
	stack.Push("d")
	el, ok := stack.Pop()
	if ok {
		fmt.Println(el)
	}

	list := list_single.NewLinkedList()

	n1 := list_single.NewNode(5)
	n2 := list_single.NewNode(6)
	n3 := list_single.NewNode(7)
	list.Add(n1)
	list.Add(n2)
	list.Add(n3)
	list.Display()
	ok, err := list.Delete(n2)
	if ok {
		fmt.Println("Node deleted")
	} else {
		panic(err)
	}
	list.Display()
	n4 := list_single.NewNode(8)
	n5 := list_single.NewNode(9)
	list.Insert(1, n4)
	list.Insert(3, n5)
	fmt.Println("-----")
	list.Display()

}
