package bst

import (
	"fmt"

	"github.com/cmy235/goDB/node"
)

type Tree struct {
	root *node.bstNode
}

func (tree *BST) addValue(val int) {
	node := bstNode{value: val, left: nil, right: nil}

	if tree.root == nil {
		tree.root = node
		fmt.Printf("Added %v as root \n", node.value)
	} else {
		tree.root.insertNode(node)
	}
}

// find and return the node in BST with this val
func (tree *BST) find(val int) *bstNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.findNode(val)
}

func (tree *BST) remove(val int) {
	if tree.root != nil {
		tree.root = tree.root.removeValue(val)
	}
}
