package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTreeNode(t *testing.T) {
	root := CreateTreeNode(42)

	assert.NotNil(t, root)
	assert.Equal(t, 42, root.value)
	assert.Nil(t, root.Left)
	assert.Nil(t, root.Right)
}

func TestTreeNode_GetTreeNodeNumWithEmptyTree(t *testing.T) {
	var root *TreeNode

	assert.Equal(t, 0, root.GetTreeNodeNum())
}

func TestTreeNode_GetTreeNodeNum(t *testing.T) {
	root := InitTree()

	assert.Equal(t, 7, root.GetTreeNodeNum())
}

func TestTreeNode_GetTreeDegreeWithEmptyTree(t *testing.T) {
	var root *TreeNode

	assert.Equal(t, 0, root.GetTreeDegree())
}

func TestTreeNode_GetTreeDegree(t *testing.T) {
	root := InitTree()

	assert.Equal(t, 3, root.GetTreeDegree())
}

func TestTreeNode_InsertTreeNodeWithEmptyTree(t *testing.T) {
	var root *TreeNode
	root.InsertTreeNode(42)

	assert.Nil(t, root)
}

func TestTreeNode_InsertTreeNode(t *testing.T) {
	root := CreateTreeNode(42)
	root.InsertTreeNode(21)

	assert.NotNil(t, root.Left)
	assert.Equal(t, 21, root.Left.value)
}

func TestTreeNode_SearchWithEmptyTree(t *testing.T) {
	var root *TreeNode
	result := root.Search(42)

	assert.Nil(t, result)
}

func TestTreeNode_SearchWhenHit(t *testing.T) {
	root := InitTree()
	result := root.Search(3)

	assert.NotNil(t, result)
	assert.Equal(t, 3, result.value)
}

func TestTreeNode_SearchWhenNotHit(t *testing.T) {
	root := InitTree()
	result := root.Search(42)

	assert.Nil(t, result)
}

func TestTreeNode_PreOrderTraverse(t *testing.T) {
	root := InitTree()
	root.PreOrderTraverse()
}

func TestTreeNode_MidOrderTraverse(t *testing.T) {
	root := InitTree()
	root.MidOrderTraverse()
}

func TestTreeNode_PostOrderTraverse(t *testing.T) {
	root := InitTree()
	root.PostOrderTraverse()
}

func TestTreeNode_LevelOrderTraverse(t *testing.T) {
	root := InitTree()
	root.LevelOrderTraverse()
}

