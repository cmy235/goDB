package node

import "fmt"

// BstNode binary search tree node
type BstNode struct {
	Value int
	Left  *BstNode
	Right *BstNode
}

// InsertNode inserts node into place from receiver node
func (node *BstNode) InsertNode(nodeToAdd *BstNode) {
	if node.Value == nodeToAdd.Value {
		return
	}

	if node.Value > nodeToAdd.Value {
		if node.Left == nil {
			node.Left = nodeToAdd
			fmt.Printf("Added %v to tree", nodeToAdd.Value)
		} else {
			// look recursively down Left side until we find nil
			node.Left.InsertNode(nodeToAdd)
		}
	} else {
		if node.Right == nil {
			node.Right = nodeToAdd
			fmt.Printf("Added %v to tree\n", nodeToAdd.Value)
		} else {
			// look down Right
			node.Right.InsertNode(nodeToAdd)
		}
	}
}

// FindNode finds node with passed value starting from receiver node
func (node *BstNode) FindNode(val int) *BstNode {
	if node == nil {
		return nil
	}

	if node.Value == val {
		fmt.Printf("Found node: %v \n", node.Value)
		return node
	}

	if val > node.Value {
		return node.Right.FindNode(val)
	}

	return node.Left.FindNode(val)
}

func (node *BstNode) RemoveValue(val int) *BstNode {
	if val < node.Value {
		node.Left = node.Left.RemoveValue(val)
		return node
	} else if val > node.Value {
		node.Right = node.Right.RemoveValue(val)
		return node
	}

	return node.deleteNode()
}

func (node *BstNode) deleteNode() *BstNode {
	fmt.Printf("Removed %v from tree \n", node.Value)
	if node.Left == nil && node.Right == nil {
		// no children
		return nil
	} else if node.Left == nil {
		// one child
		return node.Right
	} else if node.Right == nil {
		return node.Left
	} else if node.Right != nil && node.Left != nil {
		// two children
		min := node.Right.findMinNode()
		node.Value = min.Value
		node.Right = node.Right.RemoveValue(min.Value)
		return node
	}
	return nil
}

// find minimum Value beginning at given node
func (node *BstNode) findMinNode() *BstNode {
	if node.Left == nil {
		return node
	}
	return node.Left.findMinNode()
}
