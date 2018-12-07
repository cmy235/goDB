package main

import "fmt"

type bstNode struct {
	value int
	left  *bstNode
	right *bstNode
}

// BST binary search tree structure
type BST struct {
	root *bstNode
}

func main() {
	myTree := BST{root: nil}

	myTree.addValue(14)
	myTree.printTreeOut()
	myTree.addValue(27)

	myTree.printTreeOut()
	myTree.remove(27)
	myTree.printTreeOut()

	fmt.Println("Looking for 14 =>")
	myTree.find(14)
	fmt.Println("Looking for 27 =>")
	myTree.find(27)
}

func (tree *BST) printTreeOut() {
	fmt.Println(tree.root)
	fmt.Println(tree.root.left)
	fmt.Println(tree.root.right)
}

func (tree *BST) addValue(val int) {
	node := makeNode(val)

	if tree.root == nil {
		tree.root = node
		fmt.Printf("Added %v as root \n", node.value)
	} else {
		tree.root.insertNode(node)
	}
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

func makeNode(val int) *bstNode {
	node := bstNode{value: val, left: nil, right: nil}
	return &node
}

// find and return the node in BST with this val
func (tree *BST) find(val int) *bstNode {
	if tree.root == nil {
		return nil
	}

	return tree.root.findNode(val)
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

func (tree *BST) remove(val int) {
	if tree.root != nil {
		tree.root = tree.root.removeValue(val)
	}
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

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}
