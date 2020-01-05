package tree

import "testing"
import "fmt"

func TestRedBlack(t *testing.T) {
	rbt := NewRedBlack()

	rbt.Insert(3)
	// 3:b
	assertList(t, rbt.PreOrder(), []int{3})

	rbt.Insert(1)
	//    3:b
	// 1:r
	assertList(t, rbt.PreOrder(), []int{3, 1})

	rbt.Insert(2)
	//    2:b
	// 1:r   3:r
	assertList(t, rbt.PreOrder(), []int{2, 1, 3})

	rbt.Insert(4)
	//    2:b
	// 1:b   3:b
	//          4:r
	assertList(t, rbt.PreOrder(), []int{2, 1, 3, 4})

	if !rbt.Search(3) {
		t.Fatal("Expected to find 3")
	}
	if rbt.Search(5) {
		t.Fatal("Did not expect to find 5")
	}
}

func TestRBTree(t *testing.T) {
	rbt := NewRedBlack()
	rbt.Insert(3)
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(5)
	rbt.Insert(4)
	assertList(t, rbt.PreOrder(), []int{2, 1, 4, 3, 5})
	assertList(t, rbt.Ordered(), []int{1, 2, 3, 4, 5})
	if !rbt.Search(4) {
		t.Fatal("Expected to find 4")
	}
	if !rbt.Search(1) {
		t.Fatal("Expected to find 1")
	}
	if rbt.Search(42) {
		t.Fatal("Did not expect to find 42")
	}
}

func assertList(t *testing.T, a, b []int) {
	t.Helper()

	if len(a) != len(b) {
		t.Fatalf(
			"Length different: %v != %v.\n%v (output)\n%v (expected)",
			len(a), len(b), a, b)
	}

	for i, n := range b {
		if a[i] != n {
			t.Fatalf(
				"%v != %v at %v\n%v (output)\n%v (expected)",
				a[i], n, i, a, b)
		}
	}
}

func sprint(n *redBlack) string {
	if n == nil {
		return "."
	}
	c := "R"
	if n.color == black {
		c = "B"
	}
	return fmt.Sprintf("(%v:%v [%v %v])", n.key, c, sprint(n.left), sprint(n.right))
}
