package treemap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTreeMapWithPaths(t *testing.T) {
	root := NewTreeMapWithPaths("a/b", "a/c")

	assert.Empty(t, root.Node.Url)
	assert.Equal(t, "a", root.Node.Value)
	assert.Equal(t, 2, len(root.NextTree))

	child1, _ := root.NextTree["b"]
	assert.Equal(t, "a/b", child1.Node.Url)
	assert.Equal(t, "b", child1.Node.Value)

	child2, _ := root.NextTree["c"]
	assert.Equal(t, "a/c", child2.Node.Url)
	assert.Equal(t, "c", child2.Node.Value)
}

func TestTreeMap_Add_WithOneLayer(t *testing.T) {
	root := NewTreeMap()

	root.Add("a")

	assert.Equal(t, "a", root.Node.Url)
	assert.Equal(t, "a", root.Node.Value)
	assert.Empty(t, root.NextTree)
}

func TestTreeMap_Add_WithTwoLayers(t *testing.T) {
	root := NewTreeMap()

	root.Add("a/b")
	root.Add("a/c")

	assert.Empty(t, root.Node.Url)
	assert.Equal(t, "a", root.Node.Value)
	assert.Equal(t, 2, len(root.NextTree))

	child1, _ := root.NextTree["b"]
	assert.Equal(t, "a/b", child1.Node.Url)
	assert.Equal(t, "b", child1.Node.Value)

	child2, _ := root.NextTree["c"]
	assert.Equal(t, "a/c", child2.Node.Url)
	assert.Equal(t, "c", child2.Node.Value)
}

func TestTreeMap_Add_WithThreeLayers(t *testing.T) {
	root := NewTreeMap()

	root.Add("a/b/c")
	root.Add("a/b/d")

	assert.Empty(t, root.Node.Url)
	assert.Equal(t, "a", root.Node.Value)
	assert.Equal(t, 1, len(root.NextTree))

	child1, _ := root.NextTree["b"]
	assert.Empty(t, child1.Node.Url)
	assert.Equal(t, "b", child1.Node.Value)
	assert.Equal(t, 2, len(child1.NextTree))

	child2, _ := child1.NextTree["c"]
	assert.Equal(t, "a/b/c", child2.Node.Url)
	assert.Equal(t, "c", child2.Node.Value)
	assert.Empty(t, child2.NextTree)

	child3, _ := child1.NextTree["d"]
	assert.Equal(t, "a/b/d", child3.Node.Url)
	assert.Equal(t, "d", child3.Node.Value)
	assert.Empty(t, child3.NextTree)
}

func TestTreeMap_AddPaths(t *testing.T) {
	root := NewTreeMap()
	root.AddPaths("a/b", "a/c")

	assert.Empty(t, root.Node.Url)
	assert.Equal(t, "a", root.Node.Value)
	assert.Equal(t, 2, len(root.NextTree))

	child1, _ := root.NextTree["b"]
	assert.Equal(t, "a/b", child1.Node.Url)
	assert.Equal(t, "b", child1.Node.Value)

	child2, _ := root.NextTree["c"]
	assert.Equal(t, "a/c", child2.Node.Url)
	assert.Equal(t, "c", child2.Node.Value)
}

func TestTreeMap_Delete_WithOneLayer(t *testing.T) {
	root := NewTreeMap()
	root.Add("a")

	root.Delete("a")

	assert.Empty(t, root.Node.Url)
	assert.Empty(t, root.Node.Value)
	assert.Empty(t, root.NextTree)
}

func TestTreeMap_Delete_WithTwoLayers(t *testing.T) {
	root := NewTreeMap()
	root.Add("a/b")
	root.Add("a/c")

	root.Delete("a/b")

	assert.Equal(t, 1, len(root.NextTree))
	assert.Nil(t, root.NextTree["b"])

	child, _ := root.NextTree["c"]
	assert.Equal(t, "a/c", child.Node.Url)
	assert.Equal(t, "c", child.Node.Value)
}

func TestTreeMap_Delete_WithTwoLayers_DeleteRoot(t *testing.T) {
	root := NewTreeMap()
	root.Add("a/b")
	root.Add("a/c")

	root.Delete("a")

	assert.Empty(t, root.NextTree)
	assert.Nil(t, root.NextTree["a"])
	assert.Nil(t, root.NextTree["b"])
}

func TestTreeMap_Search_WithOneLayer(t *testing.T) {
	root := NewTreeMap()
	root.Add("a")

	treeMap := root.Search("a")

	assert.Equal(t, "a", treeMap.Node.Url)
	assert.Equal(t, "a", treeMap.Node.Value)
}

func TestTreeMap_Search_WithTwoLayers(t *testing.T) {
	root := NewTreeMap()
	root.Add("a/b")
	root.Add("a/c")

	treeMap := root.Search("a/c")

	assert.Equal(t, "a/c", treeMap.Node.Url)
	assert.Equal(t, "c", treeMap.Node.Value)
}

func TestTreeMap_Search_WithTwoLayers_SearchRoot(t *testing.T) {
	root := NewTreeMap()
	root.Add("a/b")
	root.Add("a/c")

	treeMap := root.Search("a")

	assert.Empty(t, treeMap.Node.Url)
	assert.Equal(t, "a", treeMap.Node.Value)
	assert.NotEmpty(t, treeMap.NextTree)
	assert.NotEmpty(t, treeMap.NextTree["b"])
	assert.NotEmpty(t, treeMap.NextTree["c"])
}

func TestTreeMap_LeafNodesWithUrl_WithOneLayer(t *testing.T) {
	root := NewTreeMap()
	root.Add("a")

	urls := root.LeafNodesWithUrl()
	assert.Equal(t, 1, len(urls))
	assert.Equal(t, "a", urls[0])
}

func TestTreeMap_LeafNodesWithUrl_WithTwoLayers(t *testing.T) {
	root := NewTreeMap()
	root.Add("a/b")
	root.Add("a/c")

	urls := root.LeafNodesWithUrl()
	assert.Equal(t, 2, len(urls))
	assert.True(t, isItemInSlice("a/b", urls))
	assert.True(t, isItemInSlice("a/c", urls))
}

func TestTreeMap_LeafNodesWithUrl_WithThreeLayers(t *testing.T) {
	root := NewTreeMapWithPaths("a/b")
	root.AddPaths("a/b/c", "a/b/d")

	urls := root.LeafNodesWithUrl()
	assert.Equal(t, 2, len(urls))
	assert.True(t, isItemInSlice("a/b/c", urls))
	assert.True(t, isItemInSlice("a/b/d", urls))
}

func isItemInSlice(item string, s []string) bool {
	for _, v := range s {
		if item == v {
			return true
		}
	}
	return false
}

func TestTreeMap_SearchUrls_WithEmptyResult(t *testing.T) {
	root := NewTreeMap()
	root.Add("a/b/c")
	root.Add("a/b/d")

	urls := root.SearchUrls("a/c")

	assert.Empty(t, len(urls))
}

func TestTreeMap_SearchUrls_WithNonEmptyResult(t *testing.T) {
	root := NewTreeMap()
	root.Add("a/b/c")
	root.Add("a/b/d")

	urls := root.SearchUrls("a/b")

	assert.Equal(t, 2, len(urls))
	assert.True(t, isItemInSlice("a/b/c", urls))
	assert.True(t, isItemInSlice("a/b/d", urls))
}

