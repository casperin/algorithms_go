package graph

func WeightedEgdesToMatrix(edges []WeightedEdge, numOfVertices int) [][]int {
	matrix := make([][]int, numOfVertices)

	for i, _ := range matrix {
		matrix[i] = make([]int, numOfVertices)
	}

	for _, edge := range edges {
		matrix[edge.U][edge.V] = edge.Weight
		matrix[edge.V][edge.U] = edge.Weight
	}

	return matrix
}

func MatrixToWeightedEdges(matrix [][]int) []WeightedEdge {
	l := len(matrix)
	edges := []WeightedEdge{}

	for u := 0; u < l; u++ {
		for v := u + 1; v < l; v++ {
			weight := matrix[u][v]
			if weight > 0 {
				// Direction strictly does not matter, but we prefer having u
				// be the smaller
				if u < v {
					edges = append(edges, WeightedEdge{u, v, weight})
				} else {
					edges = append(edges, WeightedEdge{v, u, weight})
				}
			}
		}
	}

	return edges
}
