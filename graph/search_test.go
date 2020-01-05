package graph

import "testing"

func TestSearch(t *testing.T) {
	g := NewDirectedGraph(4)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	// BFS
	bfsPath := g.BFS(2)
	assertList(t, bfsPath, []int{2, 0, 3, 1})

	// DFS
	dfsPath := g.DFS(2)
	assertList(t, dfsPath, []int{2, 0, 1, 3})
}

func assertList(t *testing.T, a, b []int) {
	t.Helper()

	if len(a) != len(b) {
		t.Fatalf(
			"%v != %v (length of inputs)\n%v\n%v",
			len(a), len(b), a, b,
		)
	}

	for i, n := range b {
		m := a[i]
		if n != m {
			t.Fatalf(
				"%v != %v at index: %v of:\nOutput:   %v\nExpected: %v",
				m, n, i, a, b,
			)
		}
	}
}
