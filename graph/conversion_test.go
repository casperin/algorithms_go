package graph

import "testing"

func TestWeightedEgdesToMatrix(t *testing.T) {
	edges := []WeightedEdge{
		WeightedEdge{0, 1, 10},
		WeightedEdge{0, 2, 6},
		WeightedEdge{0, 3, 5},
		WeightedEdge{1, 3, 15},
		WeightedEdge{2, 3, 4},
	}

	matrix := [][]int{
		[]int{0, 10, 6, 5},
		[]int{10, 0, 0, 15},
		[]int{6, 0, 0, 4},
		[]int{5, 15, 4, 0},
	}

	// Test 1: Convert edges to matrix
	matrixOutput := WeightedEgdesToMatrix(edges, 4)
	assertMatrix(t, matrixOutput, matrix)

	// Test 2: Convert matrix to edges
	edgesOutput := MatrixToWeightedEdges(matrix)

	if len(edgesOutput) != len(edges) {
		t.Fatalf("%v != %v", len(edgesOutput), len(edges))
	}

next:
	for _, e1 := range edgesOutput {
		for _, e2 := range edges {
			if e1.U == e2.U && e1.V == e2.V && e1.Weight == e2.Weight {
				continue next // found the edge
			}
		}
		t.Fatalf("Could not find (%v, %v, %v)", e1.U, e1.V, e1.Weight)
	}
}

func assertMatrix(t *testing.T, m1, m2 [][]int) {
	t.Helper()

	if len(m1) != len(m1[0]) || len(m1) != len(m2) {
		t.Fatalf("%v != %v != %v", len(m1), len(m1[0]), len(m2))
	}

	for u, _ := range m1 {
		for v, _ := range m1 {
			if m1[u][v] != m2[u][v] {
				t.Errorf(
					"Expected %v, got %v, at index (%v,%v)\n%v\n%v",
					m1[u][v], m2[u][v], u, v, m1, m2)
				return
			}
		}
	}
}
