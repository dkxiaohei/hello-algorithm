package treemap

import (
	"os"
	"testing"
)

func TestDelete_EmptyTree(t *testing.T) {
	tree := Tree{}
	tree.Fprint(os.Stdout, true, "")

	tree.Delete("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")
}

func TestDelete_LeafNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	tree.Delete("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")
}

func TestDelete_NonLeafNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	tree.Delete("a/b/c")
	tree.Fprint(os.Stdout, true, "")
}

func TestDelete_NonExistentNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	tree.Delete("a/b/c/e")
	tree.Fprint(os.Stdout, true, "")
}

func TestSearch_EmptyTree(t *testing.T) {
	tree := Tree{}
	tree.Fprint(os.Stdout, true, "")

	resultTree := tree.Search("a/b/c/e")
	resultTree.Fprint(os.Stdout, true, "")
}

func TestSearch_LeafNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	resultTree := tree.Search("a/b/c/d")
	resultTree.Fprint(os.Stdout, true, "")
}

func TestSearch_NonLeafNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Add("a/b/c/e")
	tree.Fprint(os.Stdout, true, "")

	resultTree := tree.Search("a/b/c")
	resultTree.Fprint(os.Stdout, true, "")
}

func TestSearch_NonExistentNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	resultTree := tree.Search("a/b/c/e")
	resultTree.Fprint(os.Stdout, true, "")
}

func TestFprint_EmptyTree(t *testing.T) {
	tree := Tree{}

	tree.Fprint(os.Stdout, true, "")
}

func TestFprint_OneLevel(t *testing.T) {
	tree := Tree{}
	tree.Add("/root")

	tree.Fprint(os.Stdout, true, "")
}

func TestFprint_TwoLevels(t *testing.T) {
	tree := Tree{}
	tree.Add("/root/sibling")

	tree.Fprint(os.Stdout, true, "")
}

func TestFprint_TwoSiblings(t *testing.T) {
	tree := Tree{}
	tree.Add("/root/sibling1")
	tree.Add("/root/sibling2")

	tree.Fprint(os.Stdout, true, "")
}

func TestFprint_TwoRoots(t *testing.T) {
	tree := Tree{}
	tree.Add("/root/sibling1")
	tree.Add("/root/sibling2/sibling1")
	tree.Add("/root/sibling2/sibling2")

	tree.Fprint(os.Stdout, false, "")
}
