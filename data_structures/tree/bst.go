package tree

// Implementation of a Binary Search Tree (BST) where the stored data is an
// int.
//
// The BST will not store multiples of the same Value. Calling Insert(42) on a
// BST that already contains 42 is a noop.
type BinarySearchTree struct {
	Root    *BST
	compare func(int, int) int
}

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func NewBST() BinarySearchTree {
	return BinarySearchTree{
		compare: func(i, j int) int {
			return j - i
		},
	}
}

func NewBSTWithCompare(compare func(int, int) int) BinarySearchTree {
	return BinarySearchTree{
		compare: compare,
	}
}

// Inserts a Value into the tree.
func (bst *BinarySearchTree) Insert(x int) {
	leaf := BST{Value: x}

	// Empty tree
	if bst.Root == nil {
		bst.Root = &leaf
		return
	}

	current := bst.Root

walk:
	switch {
	// we don't want duplicated Values
	// case x == current.Value:
	// return

	// We found a spot for the Value
	case bst.compare(x, current.Value) > 0 && current.Left == nil:
		current.Left = &leaf
		return
	case bst.compare(x, current.Value) < 0 && current.Right == nil:
		current.Right = &leaf
		return

	// Walk down branch
	case bst.compare(x, current.Value) > 0:
		current = current.Left
	case bst.compare(x, current.Value) < 0:
		current = current.Right
	}

	goto walk
}

func (bst *BinarySearchTree) InsertAll(Values []int) {
	for _, Value := range Values {
		bst.Insert(Value)
	}
}

// Searches the BST for a given Value. Returns true if the value is present in
// the BST. Otherwise false.
func (bst *BinarySearchTree) Search(x int) bool {
	current := bst.Root

walk:
	if current == nil {
		return false
	}

	w := bst.compare(x, current.Value)

	if w == 0 {
		return true
	}

	if w > 0 {
		current = current.Left
	} else {
		current = current.Right
	}

	goto walk
}

// Deletes a Value from the BST.
// Returns true if the value was present in the BST. Otherwise false.
func (bst *BinarySearchTree) Delete(x int) bool {
	current := &bst.Root
	w := 0

walk:
	node := *current

	if node == nil {
		return false // not found
	}

	// Compare the Value. If it's less we go down the left branch, if
	// greater we go down the right. If it's equal, then we found the node
	// we need to delete.
	w = bst.compare(x, node.Value)

	if w > 0 {
		current = &node.Left
		goto walk
	}

	if w < 0 {
		current = &node.Right
		goto walk
	}

	// Found our node

	// If we have no Left node, we just set the right one as the new root
	// and vice versa
	if node.Left == nil {
		*current = node.Right
		return true
	}
	if node.Right == nil {
		*current = node.Left
		return true
	}

	// The node to be deleted has items both left and right, so we copy the
	// smallest Value that is greater than the current node's Value (that
	// would be the lowest value down the right branch), and then delete
	// the same value down that branch, which should land us in one of the
	// two conditionals above.
	node.Value = node.Right.getMinValue()

	x = node.Value        // we will now try to delete the Value we just copied..
	current = &node.Right // ..down the Right branch
	goto walk
}

// Helper for bst.Delete(x)
func (node *BST) getMinValue() int {
	if node.Left == nil {
		return node.Value
	}
	return node.Left.getMinValue()
}

// Returns a sorted list containing the Values of the BST. Does not consume
// BST.
func (bst *BinarySearchTree) ToSortedList() []int {
	return bst.Root.toSortedList([]int{})
}

func (node *BST) toSortedList(list []int) []int {
	if node == nil {
		return list
	}
	list = node.Left.toSortedList(list)
	list = append(list, node.Value)
	list = node.Right.toSortedList(list)
	return list
}

func (bst *BinarySearchTree) PreOrder() []int {
	return bst.Root.preOrder([]int{})
}

// Returns a list in the preordered order. Does not consume BST.
func (node *BST) preOrder(list []int) []int {
	if node == nil {
		return list
	}
	list = append(list, node.Value)
	list = node.Left.preOrder(list)
	list = node.Right.preOrder(list)
	return list
}

// Balances the tree in O(n) time.
func (bst *BinarySearchTree) Balance() {
	otherBST := SortedListToBST(bst.ToSortedList())
	bst.Root = otherBST.Root
}

/**
 * Functions dealing with BST's below
 */

// Takes a sorted array and returns a balanced BST.
func SortedListToBST(list []int) BinarySearchTree {
	return BinarySearchTree{
		Root: sortedListToBSTNode(list, 0, len(list)),
	}
}

// To create a balanced BST from a sorted array, we need to take the middle
// point as the root, and then the middle point of the left half as the left
// child and vice versa for Right side.
func sortedListToBSTNode(list []int, lo, hi int) *BST {
	if lo >= hi {
		return nil
	}

	mid := (lo + hi) / 2

	return &BST{
		Value: list[mid],
		Left:  sortedListToBSTNode(list, lo, mid),
		Right: sortedListToBSTNode(list, mid+1, hi),
	}
}
