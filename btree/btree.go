package btree

// Bplustree B+ Tree implementation
type Bplustree struct {
	Root   *Node
	Order  int
	KeyMax int
}

func main() {
	myTree := Bplustree{}
	order := 3
	myTree.MakeTree(order)
}

// Node is a B+ Tree node
type Node struct {
	Keys     []int
	Parent   *Node
	Pointers []*Node
	KeyCount int
	IsLeaf   bool
}

// MakeTree adds params to struct
func (tree *Bplustree) MakeTree(order int) {
	tree.Root = &Node{}
	tree.Order = order
	tree.KeyMax = order - 1
}

// Insert value into tree
func (tree *Bplustree) Insert(key int) {
	if tree.Root == nil {
		tree.makeRoot(key)
	}
	// find the leaf at this key we want to add to
	leaf := tree.findLeaf(key)

	if leaf.KeyCount < (tree.Order - 1) {
		insertValueToLeaf(leaf, key)
	} else {
		// we hit our max, split the node, insert the key
		tree.splitAndInsert(leaf, key)
	}
}

func (tree *Bplustree) splitAndInsert(leaf *Node, key int) {
	insertValueToLeaf(leaf, key)

	mid := getMid(tree.Order)

	newLeaf := &Node{
		Parent:   leaf.Parent,
		KeyCount: tree.Order - 1 - mid,
		IsLeaf:   true,
	}

	// pointer issues here?
	j := 0
	for idx := mid; idx < tree.Order; idx++ {
		newLeaf.Keys[j] = leaf.Keys[idx]
		newLeaf.Pointers[j] = leaf.Pointers[idx]
		j++
	}

	newKeys := make([]int, mid)
	newPointers := make([]*Node, mid)
	copy(newKeys, leaf.Keys)
	copy(newPointers, leaf.Pointers)
	leaf.Keys = newKeys
	leaf.Pointers = newPointers
	leaf.KeyCount = mid - 1

	// make minimum of right leaf the keyToPromote
	keyToPromote := leaf.Keys[0]

	// newLeaf (left), leaf (right)
	tree.promoteKeyToParent(keyToPromote, newLeaf, leaf)
}

func (tree *Bplustree) promoteKeyToParent(key int, left, right *Node) error {
	// cases
	// 1 there is no parent of left, right => makeNewRoot
	if left.Parent == nil {
		return tree.makeNewRootAndInsert(key, left, right)
	}

	// 2 there is a parent, and the parents keys are < the order - 1 limit  => insert into node
	if left.Parent.KeyCount+1 < tree.Order-1 {
		insertValueToLeaf(left.Parent, key)
	}

	// 3 there is a parent, and the parents keys are now >= the order - 1 limit => split the parent and then insert key into one of the new parents
	tree.splitAndInsertIntoNode(key, left, right)
	return nil
}

func (tree *Bplustree) splitAndInsertIntoNode(key int, left, right *Node) error {
	// TODO
	// do the steps for splitting and inserting at leaf
	// also ensure parent node points to left, right
	// split left.Parent into two
	return nil
}

func (tree *Bplustree) makeNewRootAndInsert(key int, left, right *Node) error {
	// make rootNode
	tree.Root = &Node{
		KeyCount: 1,
		Parent:   nil,
	}
	tree.Root.Keys[0] = key
	tree.Root.Pointers[0] = left
	tree.Root.Pointers[1] = right

	// set left, right parents to root
	left.Parent = tree.Root
	right.Parent = tree.Root

	return nil
}

func getMid(order int) int {
	maxKeys := order - 1
	if maxKeys%2 == 0 {
		return maxKeys / 2
	}
	return (maxKeys / 2) + 1
}

func (tree *Bplustree) findLeaf(key int) *Node {
	if tree.Root == nil {
		return nil
	}

	currentNode := tree.Root
	i := 0

	// while the currentnode is NOT a leaf
	for !currentNode.IsLeaf {
		i = 0
		for i < currentNode.KeyCount {
			if key >= currentNode.Keys[i] {
				i++
			} else {
				break
			}
		}
		// reset currentNode
		currentNode = currentNode.Pointers[i]
	}
	return currentNode
}

func insertValueToLeaf(leaf *Node, key int) {
	idx := 0
	for leaf.Keys[idx] < key {
		idx++
	}

	// reset all of the keys GREATER than the idx where you're adding
	leaf.Keys = append(leaf.Keys, 0)
	copy(leaf.Keys[idx+1:], leaf.Keys[idx:])
	leaf.Keys[idx] = key
	leaf.KeyCount++
}

func (tree *Bplustree) makeRoot(key int) {
	tree.Root.Keys[0] = key
	tree.Root.Pointers = nil
	tree.Root.Parent = nil
	tree.Root.KeyCount++
	tree.Root.IsLeaf = true
}

// func Delete

// notes
// check if key already exists/duplicates in tree
// don't worry about making a linked list at the bottom of leaf nodes
