package bst

import (
	"fmt"
	"goDB/node"
)

// Tree is a binary search tree
type Tree struct {
	root *node.BstNode
}

func (tree *Tree) addValue(val int) {
	node := &node.BstNode{Value: val, Left: nil, Right: nil}

	if tree.root == nil {
		tree.root = node
		fmt.Printf("Added %v as root \n", node.Value)
	} else {
		tree.root.InsertNode(node)
	}
}

// find and return the node in BST with this val
func (tree *Tree) find(val int) *node.BstNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.FindNode(val)
}

func (tree *Tree) remove(val int) {
	if tree.root != nil {
		tree.root = tree.root.RemoveValue(val)
	}
}
