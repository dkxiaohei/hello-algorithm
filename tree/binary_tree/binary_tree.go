package binary_tree

import "math"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeNode(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

func InitTree() *TreeNode {
	root := CreateTreeNode(1)
	root.Left = CreateTreeNode(2)
	root.Right = CreateTreeNode(3)
	root.Left.Left = CreateTreeNode(4)
	root.Left.Right = CreateTreeNode(5)
	root.Right.Left = CreateTreeNode(6)
	root.Right.Right = CreateTreeNode(7)
	return root
}

func (root *TreeNode) GetTreeNodeNum() int {
	if root == nil {
		return 0
	}
	return 1 + root.Left.GetTreeNodeNum() + root.Right.GetTreeNodeNum()
}

func (root *TreeNode) GetTreeDegree() int {
	if root == nil {
		return 0
	}
	return int(1 + math.Max(float64(root.Left.GetTreeDegree()), float64(root.Right.GetTreeDegree())))
}

func (root *TreeNode) InsertTreeNode(v int) {
	if root == nil {
		return
	}
	if v <= root.Data {
		if root.Left == nil {
			root.Left = CreateTreeNode(v)
			return
		}
		root.Left.InsertTreeNode(v)
		return
	}
	if root.Right == nil {
		root.Right = CreateTreeNode(v)
		return
	}
	root.Right.InsertTreeNode(v)
}

