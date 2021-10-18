package list_double

type List interface {
	Len() int
	Front() *Node
	Back() *Node
	PushFront(value interface{}) *Node
	PushBack(value interface{}) *Node
	Remove(n *Node)
	MoveToFront(n *Node)
	isEmpty() bool
}

type Node struct {
	Prev  *Node
	value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

type doubleLinkedList struct {
	head *Node
	tail *Node
	len  int
}

func NewDoubleLinkedList() *doubleLinkedList {
	return &doubleLinkedList{}
}

func (l *doubleLinkedList) isEmpty() bool {
	return l.len == 0
}

func (l *doubleLinkedList) Len() int {
	return l.len
}

func (l *doubleLinkedList) Front() *Node {
	return l.head
}

func (l *doubleLinkedList) Back() *Node {
	return l.tail
}

func (l *doubleLinkedList) PushFront(value interface{}) *Node {
	n := NewNode(value)
	if l.isEmpty() {
		l.head = n
		l.tail = n
	} else {
		n.Next = l.head
		l.head.Prev = n
		l.head = n
	}
	l.len++
	return n
}

func (l *doubleLinkedList) PushBack(value interface{}) *Node {
	n := NewNode(value)
	if l.isEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.tail.Next = n
		n.Prev = l.tail
		l.tail = n
	}
	l.len++
	return n
}

func (l *doubleLinkedList) Remove(n *Node) {
	if l.head == n {
		if l.len == 1 {
			l.head = nil
			l.tail = nil
		} else {
			n.Next.Prev = nil
			l.head = n.Next
		}
	} else if l.tail == n {
		n.Prev.Next = nil
		l.tail = n.Prev
	} else {
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}
	l.len--
	n.Prev = nil
	n.Next = nil
}

func (l *doubleLinkedList) MoveToFront(n *Node) {
	if n == l.head || l.len == 1 && n == l.tail {
		return
	}
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev

	n.Next = l.head
	n.Prev = nil
	l.head.Prev = n
	l.head = n
}
