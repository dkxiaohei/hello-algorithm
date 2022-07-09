package binary_tree

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode struct {
	value int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
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

func (root *TreeNode) InsertTreeNode(value int) {
	if root == nil {
		return
	}
	if value <= root.value {
		if root.Left == nil {
			root.Left = CreateTreeNode(value)
			return
		}
		root.Left.InsertTreeNode(value)
		return
	}
	if root.Right == nil {
		root.Right = CreateTreeNode(value)
		return
	}
	root.Right.InsertTreeNode(value)
}

func (root *TreeNode) Search(value int) *TreeNode {
	if root == nil {
		return nil
	}
	if value == root.value {
		return root
	}
	if value < root.value {
		return root.Left.Search(value)
	}
	return root.Right.Search(value)
}

func (root *TreeNode) PreOrderTraverse() {
	if root == nil {
		return
	}
	fmt.Print(root.value, " ")
	root.Left.PreOrderTraverse()
	root.Right.PreOrderTraverse()
}

func (root *TreeNode) MidOrderTraverse() {
	if root == nil {
		return
	}
	root.Left.MidOrderTraverse()
	fmt.Print(root.value, " ")
	root.Right.MidOrderTraverse()
}

func (root *TreeNode) PostOrderTraverse() {
	if root == nil {
		return
	}
	root.Left.PostOrderTraverse()
	root.Right.PostOrderTraverse()
	fmt.Print(root.value, " ")
}

func (root *TreeNode) LevelOrderTraverse() {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		element := queue.Front()
		treeNode := element.Value.(*TreeNode)
		fmt.Print(treeNode.value, " ")
		if treeNode.Left != nil {
			queue.PushBack(treeNode.Left)
		}
		if treeNode.Right != nil {
			queue.PushBack(treeNode.Right)
		}
		queue.Remove(element)
	}
}

