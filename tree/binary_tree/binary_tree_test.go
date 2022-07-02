package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTreeNode(t *testing.T) {
	root := CreateTreeNode(42)

	assert.NotNil(t, root)
	assert.Equal(t, 42, root.Data)
	assert.Nil(t, root.Left)
	assert.Nil(t, root.Right)
}

func TestGetTreeNodeNumWithEmptyTree(t *testing.T) {
	var root *TreeNode

	assert.Equal(t, 0, root.GetTreeNodeNum())
}

func TestGetTreeNodeNum(t *testing.T) {
	root := InitTree()

	assert.Equal(t, 7, root.GetTreeNodeNum())
}

func TestGetTreeDegreeWithEmptyTree(t *testing.T) {
	var root *TreeNode

	assert.Equal(t, 0, root.GetTreeDegree())
}

func TestGetTreeDegree(t *testing.T) {
	root := InitTree()

	assert.Equal(t, 3, root.GetTreeDegree())
}

func TestInsertTreeNodeWithEmptyTree(t *testing.T) {
	var root *TreeNode
	root.InsertTreeNode(42)

	assert.Nil(t, root)
}

func TestInsertTreeNode(t *testing.T) {
	root := CreateTreeNode(42)
	root.InsertTreeNode(21)

	assert.NotNil(t, root.Left)
	assert.Equal(t, 21, root.Left.Data)
}

func TestSearchWithEmptyTree(t *testing.T) {
	var root *TreeNode
	result := root.Search(42)

	assert.Nil(t, result)
}

func TestSearchWhenHit(t *testing.T) {
	root := InitTree()
	result := root.Search(3)

	assert.NotNil(t, result)
	assert.Equal(t, 3, result.Data)
}

func TestSearchWhenNotHit(t *testing.T) {
	root := InitTree()
	result := root.Search(42)

	assert.Nil(t, result)
}

