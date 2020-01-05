package tree

import "testing"

func TestAVLTreeInsert(t *testing.T) {
	// No rotation in this test
	a := AVLTree{}
	a.Insert(40)
	checkLists(t, a.PreOrder(), []int{40})

	// Inserting 30 should give us
	//   40
	// 30
	a.Insert(30)
	checkLists(t, a.PreOrder(), []int{40, 30})
}

func TestAVLRotateLeftRight(t *testing.T) {
	// Setup
	a := AVLTree{}
	a.Insert(40)
	a.Insert(30)

	// Inserting 35 should give us before balancing:
	//   __40__
	// 30      nil
	//   35
	//
	// After balancing:
	//   35
	// 30  40
	//
	// This requires two rotations: First rotating 30 left, then 40 right.
	a.Insert(35)
	checkLists(t, a.PreOrder(), []int{35, 30, 40})
}

func TestAVLRotateRightRight(t *testing.T) {
	a := AVLTree{}
	a.Insert(30)

	// 30
	//   40
	a.Insert(40)
	checkLists(t, a.PreOrder(), []int{30, 40})

	// 30
	//   40
	//     50
	//
	// After balancing:
	//   40
	// 30  50
	a.Insert(50)
	checkLists(t, a.PreOrder(), []int{40, 30, 50})
}

func TestAVLDelete(t *testing.T) {
	a := AVLTree{}
	a.Insert(40)
	a.Insert(30)                                       // goes left
	a.Insert(50)                                       // goes right
	a.Insert(20)                                       // goes left left
	checkLists(t, a.PreOrder(), []int{40, 30, 20, 50}) // just making sure

	a.Delete(50) // we delete right branch, so we have two left in the left branch, so it should balance itself
	checkLists(t, a.PreOrder(), []int{30, 20, 40})
}
