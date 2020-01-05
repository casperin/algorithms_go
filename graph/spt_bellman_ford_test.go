package graph

import "testing"

func TestBellmanFord(t *testing.T) {
	edges := []WeightedEdge{
		{0, 1, -1},
		{0, 2, 4},
		{1, 2, 3},
		{1, 3, 2},
		{1, 4, 2},
		{3, 2, 5},
		{3, 1, 1},
		{4, 3, -3},
	}

	spt, err := BellmanFord(edges, 5, 0)

	if err != nil {
		t.Fatalf("Should not error")
	}

	if len(spt) != 5 {
		t.Fatalf("Should have length 5, but got %v", len(spt))
	}

	for i, n := range []int{0, -1, 2, -2, 1} {
		if spt[i] != n {
			t.Fatalf("%v != %v at %v", spt[i], n, i)
		}
	}
}

func TestBellmanFordNegativeCycle(t *testing.T) {
	// Negative cycle at: 1-2-3-1 (-3 + 2 + -1 = -2)
	edges := []WeightedEdge{
		{0, 2, -2},
		{1, 0, 4},
		{1, 2, -3},
		{2, 3, 2},
		{3, 1, -1},
	}

	_, err := BellmanFord(edges, 4, 0)

	if err == nil {
		t.Fatalf("Expected error")
	}
}
