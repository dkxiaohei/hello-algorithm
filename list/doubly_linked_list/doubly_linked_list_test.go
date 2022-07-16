package doubly_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode(42)

	assert.Equal(t, 42, node.value)
	assert.Nil(t, node.prev)
	assert.Nil(t, node.next)
}

func TestNewDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()

	assert.Equal(t, &DoublyLinkedList{}, l)
	assert.Nil(t, l.headNode)
}

func TestInitDoublyLinkedList(t *testing.T) {
	l := InitDoublyLinkedList(42)

	assert.NotNil(t, l.headNode)
	assert.Equal(t, 42, l.headNode.value)
}

func TestDoublyLinkedList_AddToHead(t *testing.T) {
	l := InitDoublyLinkedList(42)
	l.AddToHead(21)

	assert.Equal(t, 21, l.headNode.value)
	assert.Equal(t, 42, l.headNode.next.value)
	assert.Equal(t, l.headNode, l.headNode.next.prev)
}

func TestDoublyLinkedList_NodeWithValue(t *testing.T) {
	l := InitDoublyLinkedList(42)
	nodeWithValue := l.NodeWithValue(42)

	assert.NotNil(t, nodeWithValue)
	assert.Equal(t, 42, nodeWithValue.value)
}

func TestDoublyLinkedList_AddAfter(t *testing.T) {
	l := InitDoublyLinkedList(42)
	l.AddToHead(21)
	l.AddToHead(84)
	nodeWithValue := l.NodeBetweenValues(84, 42)

	assert.NotNil(t, nodeWithValue)
	assert.Equal(t, 21, nodeWithValue.value)
}

func TestDoublyLinkedList_LastNode(t *testing.T) {
	l := InitDoublyLinkedList(42)
	l.AddToHead(21)
	lastNode := l.LastNode()

	assert.NotNil(t, lastNode)
	assert.Equal(t, 42, lastNode.value)
}

func TestDoublyLinkedList_AddToTail(t *testing.T) {
	l := InitDoublyLinkedList(42)
	l.AddToTail(21)

	assert.Equal(t, 42, l.headNode.value)
	assert.Equal(t, 21, l.headNode.next.value)
}

func TestDoublyLinkedList_Iterate(t *testing.T) {
	l := InitDoublyLinkedList(42)
	l.AddToHead(21)
	l.AddToHead(84)

	l.Iterate()
}

