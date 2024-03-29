package treemap

import "strings"

type TreeMap struct {
	Node     *TreeNode
	NextTree map[string]*TreeMap
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
		Node:     NewTreeNode(),
		NextTree: make(map[string]*TreeMap),
	}
}

func NewTreeMapWithPaths(paths []string) *TreeMap {
	root := NewTreeMap()
	root.AddPaths(paths)
	return root
}

func (t *TreeMap) Init(url, value string) {
	t.Node.Init(url, value)
}

func (t *TreeMap) Purge() {
	t.Node = NewTreeNode()
	t.NextTree = make(map[string]*TreeMap)
}

func (t *TreeMap) Add(path string) {
	frags := strings.Split(path, "/")
	if len(frags) == 0 {
		return
	}

	if len(frags) == 1 {
		t.Init(frags[0], frags[0])
		return
	}
	t.Init(``, frags[0])

	t.add(path, frags[1:])
}

func (t *TreeMap) add(path string, frags []string) {
	if len(frags) == 0 {
		return
	}

	nextTree, ok := t.NextTree[frags[0]]
	if !ok {
		nextTree = NewTreeMap()
		if len(frags) == 1 {
			nextTree.Init(path, frags[0])
		} else {
			nextTree.Init(``, frags[0])
		}
		t.NextTree[frags[0]] = nextTree
	}

	nextTree.add(path, frags[1:])
}

func (t *TreeMap) AddPaths(paths []string) {
	if len(paths) == 0 {
		return
	}
	for _, path := range paths {
		t.Add(path)
	}
}

func (t *TreeMap) Delete(path string) {
	frags := strings.Split(path, "/")
	if len(frags) == 0 {
		return
	}

	if len(frags) == 1 {
		// delete the root of the TreeMap
		t.Purge()
		return
	}

	t.delete(frags[1:])
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
	if len(frags) == 0 {
		return nil
	}

	if len(frags) == 1 {
		if t.Node.Value == frags[0] {
			return t
		}
		return nil
	}

	return t.search(frags[1:])
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

func (t *TreeMap) LeafNodesWithUrl() []string {
	c := make(chan string)
	go t.treeNodesWalk(c)

	var urls []string
	for url := range c {
		urls = append(urls, url)
	}
	return urls
}

func (t *TreeMap) treeNodesWalk(c chan<- string) {
	t.treeNodesWalker(c)
	// close channel so that range can finish
	close(c)
}

func (t *TreeMap) treeNodesWalker(c chan<- string) {
	if len(t.NextTree) == 0 && t.Node.Url != `` {
		c <- t.Node.Url
	}

	for _, v := range t.NextTree {
		v.treeNodesWalker(c)
	}
}

func (t *TreeMap) SearchUrls(path string) []string {
	treeMap := t.Search(path)
	if treeMap == nil {
		return nil
	}
	return treeMap.LeafNodesWithUrl()
}

