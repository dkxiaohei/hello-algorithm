package treemap

import "strings"

type TreeMap struct {
	NextTree map[string]*TreeMap
	Node     *TreeNode
}

type TreeNode struct {
	Url   string
	Value string
}

func NewTreeNode() *TreeNode {
	return &TreeNode{}
}

func (n *TreeNode) Init(url, value string) {
	n.Url = url
	n.Value = value
}

func NewTreeMap() *TreeMap {
	return &TreeMap{
		NextTree: make(map[string]*TreeMap),
		Node:     NewTreeNode(),
	}
}

func (t *TreeMap) Add(path, value string) {
	frags := strings.Split(path, "/")
	t.add(path, value, frags)
}

func (t *TreeMap) add(path, value string, frags []string) {
	if len(frags) == 0 {
		return
	}

	nextTree, ok := t.NextTree[frags[0]]
	if !ok {
		nextTree = NewTreeMap()
		if len(frags) == 1 {
			nextTree.Node.Init(path, value)
		}
		t.NextTree[frags[0]] = nextTree
	}

	nextTree.add(path, value, frags[1:])
}

func (t *TreeMap) Delete(path string) {
	frags := strings.Split(path, "/")
	t.delete(frags)
}

func (t *TreeMap) delete(frags []string) {
	if len(frags) == 0 {
		return
	}

	nextTree, ok := t.NextTree[frags[0]]
	if !ok {
		return
	}

	if len(frags) == 1 {
		delete(t.NextTree, frags[0])
		return
	}

	nextTree.delete(frags[1:])
}

func (t *TreeMap) Search(path string) *TreeMap {
	frags := strings.Split(path, "/")
	return t.search(frags)
}

func (t *TreeMap) search(frags []string) *TreeMap {
	if len(frags) == 0 {
		return nil
	}

	nextTree, ok := t.NextTree[frags[0]]
	if !ok {
		return nil
	}

	if len(frags) == 1 {
		return nextTree
	}

	return nextTree.search(frags[1:])
}

