package list_single

import (
	"errors"
	"fmt"
)

type List interface {
	GetHead() *Node
	GetTail() *Node
	GetNode(at int) *Node
	Add(node *Node)
	Insert(at int, node *Node) (bool, error)
	Delete(node *Node) (bool, error)
	Len() int
	Display()
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type Node struct {
	value interface{}
	link  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) GetTail() *Node {
	return l.tail
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) GetNode(at int) *Node {
	if at >= l.Len() {
		return nil
	}

	n := l.head
	if at == 0 {
		return n
	}

	for i := 0; i < at; i++ {
		n = n.link
	}

	return n
}

func (l *LinkedList) Add(node *Node) {
	if l.Len() == 0 {
		l.head = node
	} else {
		l.tail.link = node
	}

	l.tail = node
	l.len++
}

func (l *LinkedList) Delete(nodeDel *Node) (bool, error) {
	n := l.head

	if n.value == nodeDel.value {
		l.head = nodeDel.link
		l.len--
		return true, nil
	}

	for n != nil {
		if n.link == nodeDel {
			n.link = nodeDel.link
			l.len--
			return true, nil
		}
		n = n.link
	}
	return false, errors.New("node not found")
}

func (l *LinkedList) Insert(at int, newNode *Node) (bool, error) {
	if at > l.Len() || at < 0 {
		return false, errors.New("out of boundary")
	}

	if at == 0 {
		newNode.link = l.head
		l.head = newNode
		l.len++
		return true, nil
	}

	node := l.GetNode(at)
	newNode.link = node
	prevNode := l.GetNode(at - 1)
	prevNode.link = newNode
	l.len++
	return true, nil
}

func (l *LinkedList) Display() {
	if l.Len() == 0 {
		fmt.Println("No nodes in the list")
	} else {
		node := l.head
		for node != nil {
			fmt.Println("Node -> ", node.value)
			node = node.link
		}
	}
}
