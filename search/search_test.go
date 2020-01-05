package search

import (
	"testing"
)

var searchTests = []struct {
	name string
	list []int
	x    int
	i    int
}{
	{"Simple", []int{1, 2, 3, 4, 5}, 4, 3},
	{"At the end", []int{1, 2, 3, 4}, 4, 3},
	{"At the beginning", []int{1, 2, 3, 4}, 1, 0},
	{"Missing", []int{1, 2, 4, 5}, 3, -1},
	{"Edgy one", []int{2, 3}, 3, 1},
	{"Edgy two", []int{2, 3}, 2, 0},
	{"Numbers many", []int{4, 6, 10, 23, 53, 100, 123, 132, 142, 231, 232, 259, 309}, 231, 9},
	{"Out of liste", []int{4, 6, 10, 23, 53, 100, 123, 132, 142, 231, 232, 259, 309}, 400, -1},
	{"In list, but missing", []int{4, 6, 10, 23, 53, 100, 123, 132, 142, 231, 232, 259, 309}, 200, -1},
	{"Empty list", []int{}, 42, -1},
	{"One element", []int{42}, 42, 0},
	{"One element, not it", []int{43}, 42, -1},
}

func TestSearch(t *testing.T) {
	for _, test := range searchTests {
		// Binary search
		if i := BinarySearch(test.list, test.x); i != test.i {
			t.Fatalf("Binary Search: %s failed. Expected %v, got %v", test.name, test.i, i)
		}

		// Exponential search
		if i := ExponentialSearch(test.list, test.x); i != test.i {
			t.Fatalf("Exponential Search: %s failed. Expected %v, got %v", test.name, test.i, i)
		}

		// Fibonacci search
		if i := FibonacciSearch(test.list, test.x); i != test.i {
			t.Fatalf("Fibonacci Search: %s failed. Expected %v, got %v", test.name, test.i, i)
		}

		// Interpolation search
		if i := InterpolationSearch(test.list, test.x); i != test.i {
			t.Fatalf("Interpolation Search: %s failed. Expected %v, got %v", test.name, test.i, i)
		}

		// Jump search
		if i := JumpSearch(test.list, test.x); i != test.i {
			t.Fatalf("Jump Search: %s failed. Expected %v, got %v", test.name, test.i, i)
		}

		// Linear search
		if i := LinearSearch(test.list, test.x); i != test.i {
			t.Fatalf("Linear Search: %s failed. Expected %v, got %v", test.name, test.i, i)
		}
	}
}
