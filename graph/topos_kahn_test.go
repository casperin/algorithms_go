package graph

import (
	"testing"
)

func TestToposKahn(t *testing.T) {
	a, err := ToposKahn([][]int{{1}, {}})
	if err != nil {
		t.Fatal(err)
	}
	if s := assertToposList(a, []int{0, 1}); s != "" {
		t.Fatal(s)
	}

	a2, err2 := ToposKahn([][]int{
		{1, 2},
		{3},
		{3},
		{},
	})
	if err2 != nil {
		t.Fatal(err2)
	}
	s1 := assertToposList(a2, []int{0, 1, 2, 3})
	s2 := assertToposList(a2, []int{0, 2, 1, 3})
	if s1 != "" && s2 != "" {
		t.Fatalf("\n%s\n\n%s", s1, s2)
	}

	a3, err3 := ToposKahn([][]int{
		{},
		{},
		{0},
		{1, 2},
	})
	if err3 != nil {
		t.Fatal(err3)
	}
	s3 := assertToposList(a3, []int{3, 1, 2, 0})
	s4 := assertToposList(a3, []int{3, 2, 1, 0})
	s5 := assertToposList(a3, []int{3, 2, 0, 1})
	if s3 != "" && s4 != "" && s5 != "" {
		t.Fatalf("\n%s\n\n%s\n%s", s3, s4, s5)
	}

	a4, err4 := ToposKahn([][]int{
		{},
		{},
		{3},
		{1},
		{0, 1},
		{2, 0},
	})
	if err4 != nil {
		t.Fatal(err4)
	}
	s6 := assertToposList(a4, []int{5, 4, 2, 3, 1, 0})
	s7 := assertToposList(a4, []int{5, 2, 3, 4, 1, 0})
	if s6 != "" && s7 != "" {
		t.Fatalf("\n%s\n\n%s", s6, s7)
	}
}
