package binary_tree

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeNode(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

