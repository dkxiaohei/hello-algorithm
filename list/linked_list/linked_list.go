package linked_list

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	headNode *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func InitLinkedList(value int) *LinkedList {
	l := NewLinkedList()
	l.headNode = NewNode(value)
	return l
}

func (l *LinkedList) AddToHead(value int) {
	node := NewNode(value)
	if l.headNode != nil {
		node.next = l.headNode
	}
	l.headNode = node
}

func (l *LinkedList) NodeWithValue(value int) *Node {
	for node := l.headNode; node != nil; node = node.next {
		if node.value == value {
			return node
		}
	}
	return nil
}

func (l *LinkedList) AddAfter(nodeValue, newValue int) {
	nodeWithValue := l.NodeWithValue(nodeValue)
	if nodeWithValue == nil {
		return
	}
	newNode := NewNode(newValue)
	newNode.next = nodeWithValue.next
	nodeWithValue.next = newNode
}

func (l *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := l.headNode; node != nil; node = node.next {
		if node.next == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (l *LinkedList) AddToTail(value int) {
	node := NewNode(value)
	lastNode := l.LastNode()
	if lastNode == nil {
		l.headNode = node
		return
	}
	lastNode.next = node
}

func (l *LinkedList) Iterate() {
	l.IterateWithFunc(func(value int) {
		fmt.Println(value)
	})
}

func (l *LinkedList) IterateWithFunc(f func(int)) {
	for node := l.headNode; node != nil; node = node.next {
		f(node.value)
	}
}

