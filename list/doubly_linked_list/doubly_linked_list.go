package doubly_linked_list

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	headNode *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func InitDoublyLinkedList(value int) *DoublyLinkedList {
	l := NewDoublyLinkedList()
	l.headNode = NewNode(value)
	return l
}

func (l *DoublyLinkedList) AddToHead(value int) {
	node := NewNode(value)
	if l.headNode != nil {
		node.next = l.headNode
		l.headNode.prev = node
	}
	l.headNode = node
}

func (l *DoublyLinkedList) NodeWithValue(value int) *Node {
	for node := l.headNode; node != nil; node = node.next {
		if node.value == value {
			return node
		}
	}
	return nil
}

func (l *DoublyLinkedList) AddAfter(nodeValue, newValue int) {
	nodeWithValue := l.NodeWithValue(nodeValue)
	if nodeWithValue == nil {
		return
	}
	newNode := NewNode(newValue)
	newNode.next = nodeWithValue.next
	if nodeWithValue.next != nil {
		nodeWithValue.next.prev = newNode
	}
	newNode.prev = nodeWithValue
	nodeWithValue.next = newNode
}

func (l *DoublyLinkedList) NodeBetweenValues(firstValue int, secondValue int) *Node {
	for node := l.headNode; node != nil; node = node.next {
		if node.prev != nil && node.next != nil {
			if node.prev.value == firstValue && node.next.value == secondValue {
				return node
			}
		}
	}
	return nil
}

func (l *DoublyLinkedList) LastNode() *Node {
	var lastNode *Node
	for node := l.headNode; node != nil; node = node.next {
		if node.next == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (l *DoublyLinkedList) AddToTail(value int) {
	node := NewNode(value)
	lastNode := l.LastNode()
	if lastNode == nil {
		l.headNode = node
		return
	}
	lastNode.next = node
	node.prev = lastNode
}

func (l *DoublyLinkedList) Iterate() {
	l.IterateWithFunc(func(value int) {
		fmt.Println(value)
	})
}

func (l *DoublyLinkedList) IterateWithFunc(f func(int)) {
	for node := l.headNode; node != nil; node = node.next {
		f(node.value)
	}
}

