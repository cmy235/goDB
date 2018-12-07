package Node

import "fmt"

type bstNode struct {
	value int
	left  *bstNode
	right *bstNode
}

func (node *bstNode) insertNode(nodeToAdd *bstNode) {
	if node.value == nodeToAdd.value {
		return
	}

	if node.value > nodeToAdd.value {
		if node.left == nil {
			node.left = nodeToAdd
			fmt.Printf("Added %v to tree", nodeToAdd.value)
		} else {
			// look recursively down left side until we find nil
			node.left.insertNode(nodeToAdd)
		}
	} else {
		if node.right == nil {
			node.right = nodeToAdd
			fmt.Printf("Added %v to tree\n", nodeToAdd.value)
		} else {
			// look down right
			node.right.insertNode(nodeToAdd)
		}
	}
}

func (node *bstNode) findNode(val int) *bstNode {
	if node == nil {
		return nil
	}

	if node.value == val {
		fmt.Printf("Found node: %v \n", node.value)
		return node
	}

	if val > node.value {
		return node.right.findNode(val)
	}

	return node.left.findNode(val)
}

func (node *bstNode) removeValue(val int) *bstNode {
	if val < node.value {
		node.left = node.left.removeValue(val)
		return node
	} else if val > node.value {
		node.right = node.right.removeValue(val)
		return node
	}

	return node.deleteNode()
}

func (node *bstNode) deleteNode() *bstNode {
	fmt.Printf("Removed %v from tree \n", node.value)
	if node.left == nil && node.right == nil {
		// no children
		return nil
	} else if node.left == nil {
		// one child
		return node.right
	} else if node.right == nil {
		return node.left
	} else if node.right != nil && node.left != nil {
		// two children
		min := node.right.findMinNode()
		node.value = min.value
		node.right = node.right.removeValue(min.value)
		return node
	}
	return nil
}

// find minimum value beginning at given node
func (node *bstNode) findMinNode() *bstNode {
	if node.left == nil {
		return node
	}
	return node.left.findMinNode()
}
