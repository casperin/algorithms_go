package graph

import "math"

// Implementation of Prim's algorithm for finding minimum spanning tree (MST)
// of a connected undirected weighted graph.
// Panics (index out of range) if graph is not connected.
func Prims(matrix [][]int) []WeightedEdge {
	numVertices := len(matrix)
	mstSet := make([]bool, numVertices)
	parents := make([]int, numVertices)
	weights := make([]int, numVertices)

	// We set the weight of the path to every vertice to "infinite"
	for i, _ := range weights {
		weights[i] = math.MaxInt32
	}
	// First one has no weight because that's where we start
	weights[0] = 0

	for i := 0; i < numVertices-1; i++ {
		// Get the path to the next vertice with the lowest weight
		u := primsMinKey(weights, mstSet)

		mstSet[u] = true

		for v, weight := range matrix[u] {
			// We ignore vertices already in the mst and vertices that aren't
			// connected to u
			if mstSet[v] || weight == 0 {
				continue
			}

			// If this weight is lower than any other we have, then we update
			// its parent and its known weight.
			if weight < weights[v] {
				parents[v] = u
				weights[v] = weight
			}
		}
	}

	// Build mst from parents
	mst := []WeightedEdge{}

	for v := 1; v < numVertices; v++ {
		u := parents[v]
		weight := matrix[u][v]
		mst = append(mst, WeightedEdge{u, v, weight})
	}

	return mst
}

func primsMinKey(weights []int, mstSet []bool) int {
	min := math.MaxInt32
	index := -1

	for i, w := range weights {
		if !mstSet[i] && w < min {
			min = w
			index = i
		}
	}

	return index
}
