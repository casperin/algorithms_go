package graph

import (
	"fmt"
	"testing"
)

func TestToposDFS(t *testing.T) {
	a := ToposDFS([][]int{{1}, {}})
	if s := assertToposList(a, []int{0, 1}); s != "" {
		t.Fatal(s)
	}

	a2 := ToposDFS([][]int{
		{1, 2},
		{3},
		{3},
		{},
	})
	s1 := assertToposList(a2, []int{0, 1, 2, 3})
	s2 := assertToposList(a2, []int{0, 2, 1, 3})
	if s1 != "" && s2 != "" {
		t.Fatalf("\n%s\n\n%s", s1, s2)
	}

	a3 := ToposDFS([][]int{
		{},
		{},
		{0},
		{1, 2},
	})
	s3 := assertToposList(a3, []int{3, 1, 2, 0})
	s4 := assertToposList(a3, []int{3, 2, 1, 0})
	if s3 != "" && s4 != "" {
		t.Fatalf("\n%s\n\n%s", s3, s4)
	}

	a4 := ToposDFS([][]int{
		{},
		{},
		{3},
		{1},
		{0, 1},
		{2, 0},
	})
	if s := assertToposList(a4, []int{5, 4, 2, 3, 1, 0}); s != "" {
		t.Fatal(s)
	}
}

func assertToposList(a, b []int) string {
	if len(a) != len(b) {
		return fmt.Sprintf("LENGTH: %v != %v\n%v\n%v", len(a), len(b), a, b)
	}

	for i, n := range a {
		if n != b[i] {
			return fmt.Sprintf("VALUE: %v != %v\n%v (output)\n%v", n, b[i], a, b)
		}
	}

	return ""
}
