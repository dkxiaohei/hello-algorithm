package treemap

import (
	"os"
	"testing"
)

func TestTree_DeleteForEmptyTree(t *testing.T) {
	tree := Tree{}
	tree.Fprint(os.Stdout, true, "")

	tree.Delete("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")
}

func TestTree_DeleteForLeafNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	tree.Delete("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")
}

func TestTree_DeleteForNonLeafNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	tree.Delete("a/b/c")
	tree.Fprint(os.Stdout, true, "")
}

func TestTree_DeleteForNonExistentNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	tree.Delete("a/b/c/e")
	tree.Fprint(os.Stdout, true, "")
}

func TestTree_SearchForEmptyTree(t *testing.T) {
	tree := Tree{}
	tree.Fprint(os.Stdout, true, "")

	resultTree := tree.Search("a/b/c/e")
	resultTree.Fprint(os.Stdout, true, "")
}

func TestTree_SearchForLeafNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	resultTree := tree.Search("a/b/c/d")
	resultTree.Fprint(os.Stdout, true, "")
}

func TestTree_SearchForNonLeafNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Add("a/b/c/e")
	tree.Fprint(os.Stdout, true, "")

	resultTree := tree.Search("a/b/c")
	resultTree.Fprint(os.Stdout, true, "")
}

func TestTree_SearchForNonExistentNode(t *testing.T) {
	tree := Tree{}
	tree.Add("a/b/c/d")
	tree.Fprint(os.Stdout, true, "")

	resultTree := tree.Search("a/b/c/e")
	resultTree.Fprint(os.Stdout, true, "")
}

func TestTree_FprintWithEmptyTree(t *testing.T) {
	tree := Tree{}

	tree.Fprint(os.Stdout, true, "")
}

func TestTree_FprintWithOneLevel(t *testing.T) {
	tree := Tree{}
	tree.Add("/root")

	tree.Fprint(os.Stdout, true, "")
}

func TestTree_FprintWithTwoLevels(t *testing.T) {
	tree := Tree{}
	tree.Add("/root/sibling")

	tree.Fprint(os.Stdout, true, "")
}

func TestTree_FprintWithTwoSiblings(t *testing.T) {
	tree := Tree{}
	tree.Add("/root/sibling1")
	tree.Add("/root/sibling2")

	tree.Fprint(os.Stdout, true, "")
}

func TestTree_FprintWithTwoRoots(t *testing.T) {
	tree := Tree{}
	tree.Add("/root/sibling1")
	tree.Add("/root/sibling2/sibling1")
	tree.Add("/root/sibling2/sibling2")

	tree.Fprint(os.Stdout, false, "")
}

