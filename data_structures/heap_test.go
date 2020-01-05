package datastructures

import "testing"

func TestHeap(t *testing.T) {
	h := NewHeap()

	h.Push(2)
	h.Push(-1) // replaced
	h.Push(3)
	h.Push(0)
	i := h.Replace(1) // replacing

	if i != -1 {
		t.Fatalf("%v != %v", i, 42)
	}

	testHeap(t, &h, []int{0, 1, 2, 3})
}

func TestHeapMeld(t *testing.T) {
	h := NewHeap()
	h.PushAll([]int{1, 3, 5})
	k := NewHeap()
	k.PushAll([]int{0, 2, 4})
	h.Meld(&k)
	testHeap(t, &h, []int{0, 1, 2, 3, 4, 5})
}

func TestHeapLess(t *testing.T) {
	// Max heap
	h := NewHeapWithLess(func(i, j int) bool { return i > j })
	h.Push(0)
	h.Push(3)
	h.Push(2)
	h.Push(4)
	h.Push(1)
	testHeap(t, &h, []int{4, 3, 2, 1, 0})
}

func TestStructsWithLess(t *testing.T) {
	// First we build some structs in an array that we want to sort by a key.
	xs := []struct{ x rune }{
		{x: 'd'},
		{x: 'b'},
		{x: 'c'},
		{x: 'a'},
	}

	// We use the scope to make a less function.
	h := NewHeapWithLess(func(i, j int) bool {
		return xs[i].x < xs[j].x
	})

	for i := 0; i < 4; i++ {
		h.Push(i) // push indexes.
	}

	// We get out the indexes in the order of 'x' (a = 3, b = 1, etc).
	testHeap(t, &h, []int{3, 1, 2, 0})
}

func testHeap(t *testing.T, h *Heap, expected []int) {
	t.Helper()
	out := []int{}
	for {
		v, ok := h.Pop()
		if !ok {
			break
		}
		out = append(out, v)
	}
	if len(out) != len(expected) {
		t.Fatalf("Lengths aren't the same: input had %d items, expected had %d", len(out), len(expected))
	}
	for i := 0; i < len(out); i++ {
		x := out[i]
		y := expected[i]
		if x != y {
			t.Fatalf("At %d, I got %d (from heap) and %d (from expected).\n%v (actual)\n%v (expected)", i, x, y, out, expected)
		}
	}
}
