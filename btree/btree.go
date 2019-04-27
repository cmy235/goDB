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
	// find the leaf node where we want to add the key into
	leaf := tree.findLeaf(key)

	if leaf.KeyCount < (tree.Order - 1) {
		insertValueToLeaf(leaf, key)
	} else {
		// we hit our max, split node, insert the key
		tree.splitAndInsertIntoLeaf(leaf, key)
	}
}

func (tree *Bplustree) splitAndInsertIntoLeaf(leaf *Node, key int) {
	// split this leaf into two
	// make sure two new leaves both have correct parents
	// split in half if even, or give left +1 if they're odd
	// in all cases add key into new right

	mid := getMid(tree.Order)

	newLeaf := &Node{
		Parent:   leaf.Parent,
		KeyCount: tree.Order - 1 - mid,
		IsLeaf:   true,
	}

	j := 0
	for idx := mid; idx < tree.Order; idx++ {
		newLeaf.Keys[j] = leaf.Keys[idx]
		newLeaf.Pointers[j] = leaf.Pointers[idx]
		j++
	}

	// we take the first half of the original leaves/pointers and copy them into new leaves/pointers
	newKeys := make([]int, mid)
	newPointers := make([]*Node, mid)
	copy(newKeys, leaf.Keys)
	copy(newPointers, leaf.Pointers)
	// reset the keys/pointers
	leaf.Keys = newKeys
	leaf.Pointers = newPointers
	leaf.KeyCount = mid - 1

	// make minimum of right leaf the keyToPromote
	keyToPromote := leaf.Keys[0]

	// newLeaf (left), leaf (right)
	tree.promoteKeyToParent(keyToPromote, newLeaf, leaf)
}

func (tree *Bplustree) promoteKeyToParent(key int, left, right *Node) error {

	idxToInsert := findIdxToInsert(left.Parent, left)

	// there is no parent of left, right so makeNewRoot
	if left.Parent == nil {
		return tree.makeNewRootAndInsert(key, left, right)
	}

	// there is a parent, and the parents keys are < the order - 1 limit  => insert key into right node
	// and make sure that the parent also has access to right, so it can set its Pointers to right
	if left.Parent.KeyCount+1 < tree.Order-1 {
		insertValueToNode(left.Parent, idxToInsert, key, right)
		return nil
	}

	// there is a parent, and the parents keys are now >= the order - 1 limit
	// can't just add it into the parent
	// split the parent up in half
	// make sure each half maintains proper parents
	// make sure one of the halves has the new key in it
	tree.splitAndInsertIntoNode(left.Parent, idxToInsert, key, right)
	return nil
}

func insertValueToNode(parent *Node, idxToInsert, key int, right *Node) {
	// go through parent keys
	// re assign them to one index ahead of where they are
	// then add key in at idxToInsert

	for j := parent.KeyCount; j > idxToInsert; j-- {
		parent.Pointers[j+1] = parent.Pointers[j]
		parent.Keys[j] = parent.Keys[j-1]
	}

	parent.Keys[idxToInsert+1] = key
	parent.Pointers[idxToInsert] = right
	// parent now has one more key
	parent.KeyCount++
}

func findIdxToInsert(parent, left *Node) int {
	leftIdx := 0
	// while idx is less than the total keys available && left node is not equal to the parent's child
	// (basically going along until we find the idx right BEFORE the left node)

	for leftIdx < parent.KeyCount && left != parent.Pointers[leftIdx] {
		leftIdx++
	}

	return leftIdx
}

func (tree *Bplustree) splitAndInsertIntoNode(old *Node, idxToInsert, key int, right *Node) error {
	// split up the old node by (1) making two brand new nodes from the old
	// (2) ensure that both new children nodes have old's parent
	// (3) insert key & add Node to pointers
	// (4) make sure that each new node now has the appropriate new Pointers
	var r bool

	newLeftNode := &Node{}
	newLeftNode.Parent = old.Parent

	newRightNode := &Node{}
	newRightNode.Parent = old.Parent

	halfway := getMid(tree.Order)

	leftHalfway := halfway
	rightHalfway := halfway

	if idxToInsert >= halfway {
		newRightNode.Pointers = append(newRightNode.Pointers, nil)
		rightHalfway++
		r = true
	} else {
		newLeftNode.Pointers = append(newLeftNode.Pointers, nil)
		leftHalfway++
	}

	newLeftNode.Pointers = make([]*Node, leftHalfway)
	newLeftNode.Keys = make([]int, leftHalfway)

	// take first half of keys & pointesrs from old
	copy(newLeftNode.Pointers, old.Pointers)
	copy(newLeftNode.Keys, old.Keys)

	startIdx := halfway
	stopIdx := tree.Order

	if r == true {
		stopIdx++
		halfway++
		newRightNode.Pointers[idxToInsert+1] = right
		newRightNode.Keys[idxToInsert] = key
	} else {
		newLeftNode.Pointers[idxToInsert+1] = right
		newLeftNode.Keys[idxToInsert] = key
	}

	// fill in rest of right
	for i := startIdx; i < stopIdx; i++ {
		newRightNode.Pointers[i] = old.Pointers[i]
		newRightNode.Keys[i] = old.Keys[i]
	}

	tree.promoteKeyToParent(key, newLeftNode, newRightNode)

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

	// until we hit a leaf
	for !currentNode.IsLeaf {
		i = 0
		for i < currentNode.KeyCount {
			// check if each key is greater than the one we're looking for
			if key >= currentNode.Keys[i] {
				// if so, keep increasing in keys and looking for the next one greater than our key
				i++
			} else {
				// otherwise break, and we'll re-set the currentNode with this index
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
	copy(leaf.Keys[idx+1:], leaf.Keys[idx:]) // destination, source
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
