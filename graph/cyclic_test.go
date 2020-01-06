package graph

import "testing"

func TestIsCyclic(t *testing.T) {
	adj := [][]int{
		{1, 2},
		{2},
		{0, 3},
		{3, 3},
	}
	if !IsCyclic(adj) {
		t.Fatalf("This should be a cyclic graph")
	}
}

func TestIsCyclic2(t *testing.T) {
	adj := [][]int{
		{1},
		{0},
	}
	if !IsCyclic(adj) {
		t.Fatalf("This should be a cyclic graph")
	}
}

func TestIsCyclic3(t *testing.T) {
	adj := [][]int{
		{1, 3},
		{3},
		{},
		{2},
	}
	if IsCyclic(adj) {
		t.Fatalf("This should not be a cyclic graph")
	}
}
