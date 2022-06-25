// Package treemap
// Based on https://github.com/Tufin/asciitree
package treemap

import (
	"fmt"
	"io"
	"strings"
)

// Tree can be any map with:
// 1. Key that has method 'String() string'
// 2. Value is Tree itself
// You can replace this with your own tree
type Tree map[string]Tree

func (tree Tree) Add(path string) {
	frags := strings.Split(path, "/")
	tree.add(frags)
}

func (tree Tree) add(frags []string) {
	if len(frags) == 0 {
		return
	}

	nextTree, ok := tree[frags[0]]
	if !ok {
		nextTree = Tree{}
		tree[frags[0]] = nextTree
	}

	nextTree.add(frags[1:])
}

func (tree Tree) Delete(path string) {
	frags := strings.Split(path, "/")
	tree.delete(frags)
}

func (tree Tree) delete(frags []string) {
	if len(frags) == 0 {
		return
	}

	nextTree, ok := tree[frags[0]]
	if !ok {
		return
	}

	if len(frags) == 1 {
		delete(tree, frags[0])
	}

	nextTree.delete(frags[1:])
}

func (tree Tree) Search(path string) Tree {
	frags := strings.Split(path, "/")
	return tree.search(frags)
}

func (tree Tree) search(frags []string) Tree {
	if len(frags) == 0 {
		return nil
	}

	nextTree, ok := tree[frags[0]]
	if !ok {
		return nil
	}

	if len(frags) == 1 {
		return tree
	}

	return nextTree.search(frags[1:])
}

func (tree Tree) Fprint(w io.Writer, root bool, padding string) {
	if tree == nil {
		return
	}

	index := 0
	for k, v := range tree {
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, len(tree))), k)
		v.Fprint(w, false, padding+getPadding(root, getBoxTypeExternal(index, len(tree))))
		index++
	}
}

type BoxType int

const (
	Regular BoxType = iota
	Last
	AfterLast
	Between
)

func (boxType BoxType) String() string {
	switch boxType {
	case Regular:
		return "\u251c" // ├
	case Last:
		return "\u2514" // └
	case AfterLast:
		return " "
	case Between:
		return "\u2502" // │
	default:
		panic("invalid box type")
	}
}

func getBoxType(index int, len int) BoxType {
	if index+1 == len {
		return Last
	} else if index+1 > len {
		return AfterLast
	}
	return Regular
}

func getBoxTypeExternal(index int, len int) BoxType {
	if index+1 == len {
		return AfterLast
	}
	return Between
}

func getPadding(root bool, boxType BoxType) string {
	if root {
		return ""
	}

	return boxType.String() + " "
}
