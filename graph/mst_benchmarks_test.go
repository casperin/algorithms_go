package graph

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkKruskalDense(b *testing.B) {
	edges := MatrixToWeightedEdges(benchmarkMatrix100_3)

	for n := 0; n < b.N; n++ {
		Kruskal(edges, len(benchmarkMatrix100_3))
	}
}

func BenchmarkPrimsDense(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Prims(benchmarkMatrix100_3)
	}
}

func BenchmarkBoruvkaDense(b *testing.B) {
	edges := MatrixToWeightedEdges(benchmarkMatrix100_20)

	for n := 0; n < b.N; n++ {
		Boruvka(edges, len(benchmarkMatrix100_3))
	}
}

func BenchmarkKruskal(b *testing.B) {
	edges := MatrixToWeightedEdges(benchmarkMatrix100_20)

	for n := 0; n < b.N; n++ {
		Kruskal(edges, len(benchmarkMatrix100_20))
	}
}

func BenchmarkPrims(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Prims(benchmarkMatrix100_20)
	}
}

func BenchmarkBoruvka(b *testing.B) {
	edges := MatrixToWeightedEdges(benchmarkMatrix100_20)

	for n := 0; n < b.N; n++ {
		Boruvka(edges, len(benchmarkMatrix100_20))
	}
}

var printed = false

// "Matrix" is printed to m.txt
func createMatrix() {
	rand.Seed(time.Now().UnixNano() / 1000000)
	n := 100
	m := make([][]int, n)
	everyXConnected := 20
	file := "m.txt"

	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	for u := 0; u < n; u++ {
		for v := u + 1; v < n; v++ {
			create := rand.Intn(everyXConnected)
			if create == 1 {
				w := rand.Intn(n / 10)
				m[u][v] = w
				m[v][u] = w
			}
		}
	}

	s := ""
	for _, r := range m {
		s += fmt.Sprint("[]int{")
		for i, w := range r {
			if i == 0 {
				s += fmt.Sprint(w)
			} else {
				s += fmt.Sprintf(", %d", w)
			}
		}
		s += fmt.Sprint("},\n")
	}
	printed = true
	d1 := []byte(s)
	ioutil.WriteFile(file, d1, 0644)
}
