package datastructures

import (
	"testing"
)

func TestLinkedListPushPopPeekDelete(t *testing.T) {
	l := LinkedList{}
	l.Push(2)
	l.Push(42)
	l.Push(1)
	if x, ok := l.Peek(); x != 1 || !ok {
		t.Fatalf("Expected 1 and ok, got %v", x)
	}
	l.Push(0)

	if !l.Delete(42) {
		t.Fatalf("Expected deleting a key in the list to return true")
	}

	if l.Delete(43) {
		t.Fatalf("Expected deleting a key not in the list to return false")
	}

	if size := l.Size(); size != 3 {
		t.Fatalf("Size is wrong. Expected 3, got %v", size)
	}

	linkedListFits(t, &l, []int{0, 1, 2})

	if y, ok := l.Pop(); ok {
		t.Fatalf("Expected no value, but got ok, and %d", y)
	}
}

func TestLinkedListStringer(t *testing.T) {
	l := LinkedList{}
	l.PushAll([]int{1, 2, 3})
	s := l.String()
	if s != "3->2->1->nil" {
		t.Fatalf("Did not get expected string representaion: %s", s)
	}
}

func TestLinkedListDelete(t *testing.T) {
	// There is an edge case with just one element
	l := LinkedList{}
	l.Push(1)
	l.Delete(1)
	if !l.IsEmpty() {
		t.Fatal("Expected empty list")
	}
}

func TestLinkedListReverse(t *testing.T) {
	l := LinkedList{}
	l.PushAll([]int{0, 1, 2, 3})
	l.Reverse()
	linkedListFits(t, &l, []int{0, 1, 2, 3})
}

func TestLinkedListRotate(t *testing.T) {
	l := LinkedList{}
	l.PushAll([]int{6, 5, 4, 3, 2, 1})
	l.Rotate(4)
	linkedListFits(t, &l, []int{5, 6, 1, 2, 3, 4})

	// Should be noop, if rotation distance is larger than list
	m := LinkedList{}
	m.PushAll([]int{3, 2, 1})
	m.Rotate(4)
	linkedListFits(t, &m, []int{1, 2, 3})

	// len < 2 should be a noop
	n := LinkedList{}
	n.PushAll([]int{1})
	n.Rotate(2)
	linkedListFits(t, &n, []int{1})
}

func linkedListFits(t *testing.T, list *LinkedList, values []int) {
	t.Helper()
	out := []int{}
	for {
		n, ok := list.Pop()
		if !ok {
			break
		}
		out = append(out, n)
	}
	if len(out) != len(values) {
		t.Fatalf("Lengths aren't the same: input had %d items, expected had %d", len(out), len(values))
	}
	for i := 0; i < len(out); i++ {
		x := out[i]
		y := values[i]
		if x != y {
			t.Fatalf("At %d, I got %d (from linkedlist) and %d (from expected).\n%v (actual)\n%v (expected)", i, x, y, out, values)
		}
	}
}
