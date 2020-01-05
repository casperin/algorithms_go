package tree

import "testing"

func TestBinarySearchTree(t *testing.T) {
	bst := NewBST()

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(8)

	if bst.Search(1) {
		t.Fatal("Did not expect to find 1")
	}

	if !bst.Search(7) {
		t.Fatal("Did expect to find 7")
	}
	if !bst.Search(6) {
		t.Fatal("Did expect to find 6")
	}

	if bst.Delete(42) {
		t.Fatalf("Should return false for non existing value")
	}

	bst.Delete(7)
	if bst.Search(7) {
		t.Fatal("No longer expected to find 7")
	}
	bst.Delete(2)
	if bst.Search(2) {
		t.Fatal("No longer expected to find 2")
	}
	bst.Delete(3)
	if bst.Search(3) {
		t.Fatal("No longer expected to find 3")
	}
	bst.Delete(8)
	if bst.Search(8) {
		t.Fatal("No longer expected to find 8")
	}
	bst.Delete(4)
	if bst.Search(4) {
		t.Fatal("No longer expected to find 4")
	}
	bst.Delete(6)
	if bst.Search(6) {
		t.Fatal("No longer expected to find 6")
	}
}

func TestBSTToSortedList(t *testing.T) {
	list := []int{4, 2, 1, 3, 6, 5, 0}
	l := NewBST()
	l.InsertAll(list)
	sorted := l.ToSortedList()
	for i, n := range sorted {
		if i != n {
			t.Fatalf("%d != %d", i, n)
		}
	}
}

func TestSortedListToBST(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	l := SortedListToBST(list)

	preOrdered := l.PreOrder()
	expected := []int{4, 2, 1, 3, 6, 5, 7}

	checkLists(t, preOrdered, expected)
}

func TestBinarySearchTreeBalance(t *testing.T) {
	list := []int{7, 6, 5, 4, 3, 2, 1}
	l := NewBST()
	// create a tree with only left branches with 7 at the root and 1 at the bottom
	l.InsertAll(list)

	checkLists(t, l.PreOrder(), list)

	l.Balance()
	checkLists(t, l.PreOrder(), []int{4, 2, 1, 3, 6, 5, 7})
}

func TestBinarySearchTreeWithCompare(t *testing.T) {
	xs := []struct{ x rune }{
		{x: 'c'}, // 2 = 0
		{x: 'd'}, // 3 = 1
		{x: 'b'}, // 1 = 2
		{x: 'a'}, // 0 = 3
		{x: 'e'}, // x = 4 -- we don't add this
	}

	bst := NewBSTWithCompare(func(i, j int) int {
		switch {
		case xs[i].x < xs[j].x:
			return 1
		case xs[i].x > xs[j].x:
			return -1
		default:
			return 0
		}
	})

	// We push them in this order so it's easier to reason about its preordered
	// output.
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(0)

	if !bst.Search(3) {
		t.Fatal("Three should be there")
	}
	if bst.Search(4) {
		t.Fatal("Four should not be there")
	}

	checkLists(t, bst.PreOrder(), []int{2, 3, 1, 0})
}

func checkLists(t *testing.T, output, expected []int) {
	t.Helper()
	if len(output) != len(expected) {
		t.Fatalf(
			"Lists are different length.\n%v (output: %v)\n%v (expected: %v)",
			output, len(output), expected, len(expected),
		)
	}
	for i := 0; i < len(output); i++ {
		n := output[i]
		m := expected[i]
		if n != m {
			t.Fatalf(
				"At %d, got %d, expected %d\n%v (output)\n%v (expected)",
				i, n, m, output, expected,
			)
		}
	}
}
