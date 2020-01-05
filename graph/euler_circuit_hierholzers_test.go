package graph

import "testing"

func TestHierholzers(t *testing.T) {
	edges := [][]int{
		{1},    // 0 -> 1
		{2, 3}, // 1 -> 2 & 1 -> 3
		{0},    // 2 -> 0
		{4},    // 3 -> 4
		{1},    // 4 -> 1
	}

	expected := []int{0, 1, 3, 4, 1, 2, 0}

	circuit := Hierholzers(edges, 0)

	if len(circuit) != len(expected) {
		t.Fatalf("Expected circuit of length %v, got %v", len(expected), len(circuit))
	}

	for i, n := range expected {
		if circuit[i] != n {
			t.Fatalf("%v != %v at %v\n%v (circuit)\n%v (expected)",
				circuit[i], n, i, circuit, expected)
		}
	}
}

func TestHierholzers2(t *testing.T) {
	adjacencyList := [][]int{
		{1, 6},
		{2},
		{0, 3},
		{4},
		{2, 5},
		{0},
		{4},
	}

	expected := []int{0, 6, 4, 5, 0, 1, 2, 3, 4, 2, 0}

	circuit := Hierholzers(adjacencyList, 0)

	if len(circuit) != len(expected) {
		t.Fatalf("Expected circuit of length %v, got %v", len(expected), len(circuit))
	}

	for i, n := range expected {
		if circuit[i] != n {
			t.Fatalf("%v != %v at %v\n%v (circuit)\n%v (expected)",
				circuit[i], n, i, circuit, expected)
		}
	}
}
