package graph

import (
	"errors"
	"math"
)

var ErrNegativeCycle = errors.New("Negative Cycle")

func BellmanFord(edges []WeightedEdge, numVertices, src int) ([]int, error) {
	// Set all distances to "infinite"
	spt := make([]int, numVertices)
	for i := 0; i < numVertices; i++ {
		spt[i] = math.MaxInt64
	}
	spt[src] = 0

	// Relax them edges
	for i := 0; i < numVertices-1; i++ {
		for _, edge := range edges {
			u := edge.U
			v := edge.V
			w := edge.Weight

			if spt[u] != math.MaxInt64 && spt[u]+w < spt[v] {
				spt[v] = spt[u] + w
			}
		}
	}

	// Detect negative weight cycles.
	for _, edge := range edges {
		if spt[edge.U] != math.MaxInt64 && spt[edge.U]+edge.Weight < spt[edge.V] {
			return nil, ErrNegativeCycle
		}
	}

	return spt, nil
}
