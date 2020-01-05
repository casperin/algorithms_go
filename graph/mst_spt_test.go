package graph

import (
	"math"
	"testing"
)

var graphTests = []struct {
	input [][]int
	mst   [][]int
	spt   []int
}{
	{
		input: [][]int{
			// Edges:
			// (0, 1, 10)
			// (0, 2,  6)
			// (0, 3,  5)
			// (1, 3, 15)
			// (2, 3,  4)
			[]int{0, 10, 6, 5},
			[]int{10, 0, 0, 15},
			[]int{6, 0, 0, 4},
			[]int{5, 15, 4, 0},
		},
		mst: [][]int{
			[]int{0, 10, 0, 5},
			[]int{10, 0, 0, 0},
			[]int{0, 0, 0, 4},
			[]int{5, 0, 4, 0},
		},
	},
	{
		input: [][]int{
			[]int{0, 3, 0, 3, 0},
			[]int{3, 0, 0, 0, 4},
			[]int{0, 0, 0, 2, 1},
			[]int{3, 3, 2, 0, 0},
			[]int{0, 4, 1, 0, 0},
		},
		mst: [][]int{
			[]int{0, 3, 0, 3, 0},
			[]int{3, 0, 0, 0, 0},
			[]int{0, 0, 0, 2, 1},
			[]int{3, 0, 2, 0, 0},
			[]int{0, 0, 1, 0, 0},
		},
	},
	{
		input: [][]int{
			{0, 28, 0, 0, 0, 10, 0},
			{28, 0, 16, 0, 0, 0, 14},
			{0, 16, 0, 12, 0, 0, 0},
			{0, 0, 12, 0, 22, 0, 18},
			{0, 0, 0, 22, 0, 25, 24},
			{10, 0, 0, 0, 25, 0, 0},
			{0, 14, 0, 18, 24, 0, 0},
		},
		mst: [][]int{
			{0, 0, 0, 0, 0, 10, 0},
			{0, 0, 16, 0, 0, 0, 14},
			{0, 16, 0, 12, 0, 0, 0},
			{0, 0, 12, 0, 22, 0, 0},
			{0, 0, 0, 22, 0, 25, 0},
			{10, 0, 0, 0, 25, 0, 0},
			{0, 14, 0, 0, 0, 0, 0},
		},
	},
	{
		input: [][]int{
			{0, 4, 0, 0, 0, 0, 0, 8, 0},
			{4, 0, 8, 0, 0, 0, 0, 11, 0},
			{0, 8, 0, 7, 0, 4, 0, 0, 2},
			{0, 0, 7, 0, 9, 14, 0, 0, 0},
			{0, 0, 0, 9, 0, 10, 0, 0, 0},
			{0, 0, 4, 14, 10, 0, 2, 0, 0},
			{0, 0, 0, 0, 0, 2, 0, 1, 6},
			{8, 11, 0, 0, 0, 0, 1, 0, 7},
			{0, 0, 2, 0, 0, 0, 6, 7, 0},
		},
		spt: []int{0, 4, 12, 19, 21, 11, 9, 8, 14},
	},
	{
		// 0 and 1 are connected
		// 2 and 2 are connected, but since we are staring with 0, these are
		// disconnected from the starting point, so we get INF distance
		input: [][]int{
			{0, 1, 0, 0},
			{1, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 0, 1, 0},
		},
		spt: []int{0, 1, math.MaxInt64, math.MaxInt64},
	},
}

func TestKruskal(t *testing.T) {
	for _, test := range graphTests {
		if len(test.mst) == 0 {
			continue
		}
		edges := MatrixToWeightedEdges(test.input)
		mst := Kruskal(edges, len(test.input))
		matrix := WeightedEgdesToMatrix(mst, len(test.mst))
		assertMatrix(t, matrix, test.mst)
	}
}

func TestPrims(t *testing.T) {
	for _, test := range graphTests {
		if len(test.mst) == 0 {
			continue
		}
		mst := Prims(test.input)
		matrix := WeightedEgdesToMatrix(mst, len(test.mst))
		assertMatrix(t, matrix, test.mst)
	}
}

func TestBoruvka(t *testing.T) {
	for _, test := range graphTests {
		if len(test.mst) == 0 {
			continue
		}
		edges := MatrixToWeightedEdges(test.input)
		mst := Boruvka(edges, len(test.input))
		matrix := WeightedEgdesToMatrix(mst, len(test.mst))
		assertMatrix(t, matrix, test.mst)
	}
}

func TestDijkstra(t *testing.T) {
	for _, test := range graphTests {
		if len(test.spt) == 0 {
			continue
		}
		spt := Dijkstra(test.input, 0)

		if len(spt) != len(test.spt) {
			t.Fatalf("Diff length of spt and expected spt: \n%v (spt)\n%v (expected)", spt, test.spt)
		}

		for i, dist := range test.spt {
			if dist != spt[i] {
				t.Fatalf(
					"Spt failed at %v (%v != %v).\n%v (spt)\n%v (expected)",
					i, spt[i], dist, spt, test.spt,
				)
			}
		}
	}
}

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
