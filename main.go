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

	newNode := makeNode(14)
	myTree.add(&newNode)
	secondNode := makeNode(27)
	myTree.add(&secondNode)

	myTree.remove(27)
	myTree.find(14)
}

func (tree *BST) add(node *bstNode) {
	if tree.root == nil {
		tree.root = node
		fmt.Printf("Added %v as root \n", node.value)
	} else {
		tree.insertNode(tree.root, node)
	}
}

func (tree *BST) insertNode(node, nodeToAdd *bstNode) {
	if node.value == nodeToAdd.value {
		return
	}

	if node.value > nodeToAdd.value {
		if node.left == nil {
			node.left = nodeToAdd
			fmt.Printf("Added %v to tree", nodeToAdd.value)
		} else {
			// look recursively down left side until we find nil
			tree.insertNode(node.left, nodeToAdd)
		}
	} else {
		if node.right == nil {
			// set the node child
			node.right = nodeToAdd
			fmt.Printf("Added %v to tree\n", nodeToAdd.value)
		} else {
			// look recursively down right
			tree.insertNode(node.right, nodeToAdd)
		}
	}
}

func makeNode(val int) bstNode {
	node := bstNode{value: val, left: nil, right: nil}
	return node
}

// find and return the node in BST with this val
func (tree *BST) find(val int) *bstNode {
	if tree.root == nil {
		return nil
	}

	return tree.findNode(val, tree.root)
}

func (tree *BST) findNode(val int, node *bstNode) *bstNode {
	if node == nil {
		return nil
	}

	if node.value == val {
		fmt.Printf("Found node: %v \n", node.value)
		return node
	}

	if val > node.value {
		return tree.findNode(val, node.right)
	}

	return tree.findNode(val, node.left)
}

func (tree *BST) remove(val int) {
	if tree.root != nil {
		tree.root = tree.removeNode(tree.root, val)
	}
}

func (tree *BST) removeNode(node *bstNode, val int) *bstNode {
	if val < node.value {
		node.left = tree.removeNode(node.left, val)
		return node
	} else if val > node.value {
		node.right = tree.removeNode(node.right, val)
		return node
	} else {
		// ! found !
		fmt.Printf("Removed %v from tree \n", node.value)
		// no children
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left == nil { // one child
			return node.right
		} else if node.right == nil {
			return node.left
		} else if node.right != nil && node.left != nil { // two children
			// need inorder successor of node we're deleting
			min := tree.findMinNode(node.right)
			node.value = min.value
			// update current node to refleact that node's data
			// remove original min node since it was promoted
			node.right = tree.removeNode(node.right, min.value)
			return node
		} else {
			return nil
		}
	}
}

// find minimum value in tree beginning at given node
func (tree *BST) findMinNode(node *bstNode) bstNode {
	if node.left == nil {
		return *node
	}
	return tree.findMinNode(node.left)
}

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c
}
