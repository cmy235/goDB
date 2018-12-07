package main

import (
	"fmt"

	"github.com/cmy235/goDB/bst"
)

func main() {
	myTree := bst.Tree{root: nil}

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
