package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode(42)

	assert.Equal(t, 42, node.value)
	assert.Nil(t, node.next)
}

func TestNewLinkedList(t *testing.T) {
	l := NewLinkedList()

	assert.Equal(t, &LinkedList{}, l)
	assert.Nil(t, l.headNode)
}

func TestInitLinkedList(t *testing.T) {
	l := InitLinkedList(42)
	assert.NotNil(t, l.headNode)
	assert.Equal(t, 42, l.headNode.value)
	assert.Nil(t, l.headNode.next)
}

func TestLinkedList_AddToHead(t *testing.T) {
	l := InitLinkedList(42)
	l.AddToHead(21)

	assert.Equal(t, 21, l.headNode.value)
	assert.Equal(t, 42, l.headNode.next.value)
}

func TestLinkedList_NodeWithValue(t *testing.T) {
	l := InitLinkedList(42)
	nodeWithValue := l.NodeWithValue(42)

	assert.Equal(t, nodeWithValue, l.headNode)
}

func TestLinkedList_AddAfter(t *testing.T) {
	l := InitLinkedList(42)
	l.AddToHead(21)
	l.AddAfter(21, 84)

	assert.Equal(t, 84, l.headNode.next.value)
}

func TestLinkedList_LastNode(t *testing.T) {
	l := InitLinkedList(42)
	l.AddToHead(21)

	assert.Equal(t, 42, l.LastNode().value)
}

func TestLinkedList_AddToTail(t *testing.T) {
	l := InitLinkedList(42)
	l.AddToTail(84)

	assert.Equal(t, 42, l.headNode.value)
	assert.Equal(t, 84, l.headNode.next.value)
}

func TestLinkedList_Iterate(t *testing.T) {
	l := InitLinkedList(42)
	l.AddToHead(21)
	l.AddToTail(84)

	l.Iterate()
}

