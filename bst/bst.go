package bst

import (
	"fmt"
	"goDB/node"
)

// Tree is a binary search tree
type Tree struct {
	Root *node.BstNode
}

// AddValue adds value to tree
func (tree *Tree) AddValue(val int) {
	node := &node.BstNode{Value: val, Left: nil, Right: nil}

	if tree.Root == nil {
		tree.Root = node
		fmt.Printf("Added %v as Root \n", node.Value)
	} else {
		tree.Root.InsertNode(node)
	}
}

// Find ginds and returns the node in tree with this val
func (tree *Tree) Find(val int) *node.BstNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindNode(val)
}

// Remove removes value from tree
func (tree *Tree) Remove(val int) {
	if tree.Root != nil {
		tree.Root = tree.Root.RemoveValue(val)
	}
}
