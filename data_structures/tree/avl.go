package tree

import "fmt"

type AVLTree struct {
	Root *AVL
}

type AVL struct {
	Value  int
	Left   *AVL
	Right  *AVL
	Height int
}

func SprintAVL(avl *AVL) string {
	if avl == nil {
		return "_"
	}
	l := SprintAVL(avl.Left)
	r := SprintAVL(avl.Right)
	return fmt.Sprintf("(%v %v %v)", avl.Value, l, r)
}

func (t *AVLTree) Insert(value int) {
	t.Root = t.Root.Insert(value)
}

func (t *AVLTree) Delete(value int) {
	t.Root = t.Root.Delete(value)
}

func (avl *AVL) Insert(value int) *AVL {
	if avl == nil {
		return &AVL{Value: value, Height: 1}
	}

	if value < avl.Value {
		avl.Left = avl.Left.Insert(value)
	} else if value > avl.Value {
		avl.Right = avl.Right.Insert(value)
	} else {
		return avl
	}

	// Update height of current node
	avl.Height = 1 + max(avl.Left.height(), avl.Right.height())

	balance := avl.balance()

	// Left Left case
	if balance > 1 && value < avl.Left.Value {
		return avl.rotateRight()
	}

	// Right Right case
	if balance < -1 && value > avl.Right.Value {
		return avl.rotateLeft()
	}

	// Left Right case
	if balance > 1 && value > avl.Left.Value {
		avl.Left = avl.Left.rotateLeft()
		return avl.rotateRight()
	}

	// Right Left case
	if balance < -1 && value < avl.Right.Value {
		avl.Right = avl.Right.rotateRight()
		return avl.rotateLeft()
	}

	return avl
}

func (avl *AVL) Delete(value int) *AVL {
	if avl == nil {
		return nil
	}

	// First we perform a standard recursive deletion. Once that's done, we
	// climb back up the tree (that's why it needs to be recursive) and balance
	// the tree out.

	if value < avl.Value {
		avl.Left = avl.Left.Delete(value) // go down left branch
	} else if value > avl.Value {
		avl.Right = avl.Right.Delete(value) // go down right branch
	} else {
		// Node found.

		if avl.Left == nil && avl.Right == nil {
			avl = nil // No children, we just delete it
		} else if avl.Left == nil {
			*avl = *avl.Right // No left child means we can just move the right branch up to this one
		} else if avl.Right == nil {
			*avl = *avl.Left // ... and vice versa
		} else {
			// Two children. Copy the lowest value we find down the right
			// branch to this one, then continue deleting down that branch.
			avl.Value = avl.Right.getMinValue()
			avl.Right = avl.Right.Delete(avl.Value)
		}
	}

	// Value has been deleted, so now we must rebalance the tree.

	// If this was a leaf node, then we are done. Nothing to balance.
	if avl == nil {
		return avl
	}

	// Update height of current node
	avl.Height = 1 + max(avl.Left.height(), avl.Right.height())

	balance := avl.balance()

	// Left Left case
	if balance > 1 && avl.Left.balance() >= 0 {
		return avl.rotateRight()
	}

	// Left Right case
	if balance > 1 && avl.Left.balance() < 0 {
		avl.Left = avl.Left.rotateLeft()
		return avl.rotateRight()
	}

	// Right Right case
	if balance < -1 && avl.Right.balance() <= 0 {
		return avl.rotateLeft()
	}

	// Right Left case
	if balance < -1 && avl.Right.balance() > 0 {
		avl.Right = avl.Right.rotateRight()
		return avl.rotateLeft()
	}

	return avl
}

// Helper for bst.Delete(x)
func (avl *AVL) getMinValue() int {
	if avl.Left == nil {
		return avl.Value
	}
	return avl.Left.getMinValue()
}

//    __avl__           __x__
//   x       T3  -->  T1     avl
// T1 T2                   T2   T3
func (avl *AVL) rotateRight() *AVL {
	x := avl.Left
	T2 := x.Right

	// Perform rotation
	x.Right = avl
	avl.Left = T2

	// Update height
	avl.Height = 1 + max(avl.Left.height(), avl.Right.height())
	x.Height = 1 + max(x.Left.height(), x.Right.height())

	return x
}

//    __avl__               __x__
//  T1       x    -->    avl     T3
//         T2 T3       T1   T2
func (avl *AVL) rotateLeft() *AVL {
	x := avl.Right
	T2 := x.Left

	// Perform rotation
	x.Left = avl
	avl.Right = T2

	// Update heights
	avl.Height = 1 + max(avl.Left.height(), avl.Right.height())
	x.Height = 1 + max(x.Left.height(), x.Right.height())

	return x
}

func (avl *AVL) balance() int {
	if avl == nil {
		return 0
	}
	return avl.Left.height() - avl.Right.height()
}

func (avl *AVL) height() int {
	if avl == nil {
		return 0
	}
	return avl.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *AVLTree) PreOrder() []int {
	return t.Root.preOrder([]int{})
}

func (node *AVL) preOrder(list []int) []int {
	if node == nil {
		return list
	}
	list = append(list, node.Value)
	list = node.Left.preOrder(list)
	list = node.Right.preOrder(list)
	return list
}
