package graph

import "fmt"

// For this type of graph, all vertices in the graph are expected to be
// numbered from 0 to V, where V (the size of the grapth) is expected to be
// known ahead of time.
type DirectedGraph struct {
	// List with length V, has one slot for each vertice in the graph with a
	// list of ints that it is connected to. So for instance a graph with two
	// vertices, 0 and 1, with an edge from 0 to 1, will be represented as:
	// List = [[1], []]
	// The vertice at position 0 has an edge to 1. The vertice at position 1
	// has no edges.
	List [][]int
}

func NewDirectedGraph(v int) DirectedGraph {
	return DirectedGraph{make([][]int, v)}
}

// Add edge to graph.
// Panics if either vertice (int) does not fit in the list of the graph.
func (g *DirectedGraph) addEdge(u, v int) {
	if l := len(g.List); u >= l || v >= l {
		panic(fmt.Sprintf("Edge (%v, %v) does not fit in a graph with list of length %v", u, v, l))
	}
	g.List[u] = append(g.List[u], v)
}

// An undirected weighted graph often consists of just a list of weighted edges.
type WeightedEdge struct {
	U      int
	V      int
	Weight int
}
