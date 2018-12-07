package main

import (
	"fmt"
	"goDB/bst"
)

func main() {
	myTree := bst.Tree{Root: nil}

	myTree.AddValue(14)
	myTree.AddValue(27)

	myTree.Remove(27)

	fmt.Println("Looking for 14 =>")
	myTree.Find(14)
	fmt.Println("Looking for 27 =>")
	myTree.Find(27)
}
