package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
