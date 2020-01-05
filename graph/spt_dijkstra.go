package graph

import (
	"math"
)

func Dijkstra(matrix [][]int, src int) []int {
	numVertices := len(matrix)
	dists := make([]int, numVertices)
	sptSet := make([]bool, numVertices)

	// Set distance to all vertices to infinite.
	for i := 0; i < numVertices; i++ {
		dists[i] = math.MaxInt64
	}
	dists[src] = 0 // except starting node.

	// Loop through all vertices (except starting node).
	for i := 1; i < numVertices-1; i++ {

		// Find the index of the vertice with the lowest dist.
		u := -1
		uDist := math.MaxInt64

		for i := 0; i < numVertices; i++ {
			if !sptSet[i] && dists[i] <= uDist {
				u = i
				uDist = dists[i]
			}

		}

		// Mark the vertice as added
		sptSet[u] = true

		// Ignore we only have disconnected vertices left, then we're done
		if uDist == math.MaxInt64 {
			break
		}

		// Loop through all vertices again
		for v := 0; v < numVertices; v++ {
			// Ignore if...
			switch {
			case sptSet[v]: // we already added this
				continue
			case matrix[u][v] == 0: // they are not connected (or u == v)
				continue
			case uDist+matrix[u][v] >= dists[v]: // new path is not shorter
				continue
			}

			dists[v] = uDist + matrix[u][v]
		}
	}

	return dists
}
