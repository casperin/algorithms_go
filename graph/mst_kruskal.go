package graph

import "sort"

// Implementation of kruskal's algorithm for finding the minimum spanning tree
// (MST) of a connected undirected weighted graph. Sorts the list of edges
// provided.
// Returns the wrong result if the graph is not connected.
func Kruskal(edges []WeightedEdge, numVertices int) []WeightedEdge {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	mst := []WeightedEdge{} // output

	//  For Kruskal's algorithm we need to keep track of which vertices are
	//  connected. We do this by having a map where each vertice can point to
	//  another vertice (their "parent") as away to say to say they are
	//  connected to it.
	parents := make([]int, numVertices)
	for i, _ := range parents {
		parents[i] = i
	}

	for _, edge := range edges {
		parentU := findParent(parents, edge.U)
		parentV := findParent(parents, edge.V)

		// If the two vertices have the same parent, then they are connected
		// already, so we do not want to add them to our mst.
		if parentU == parentV {
			continue
		}

		mst = append(mst, edge)    // Edge is good
		parents[parentU] = parentV // Connect the vertices
	}

	return mst
}

// Recursively find the parent of a vertice.
func findParent(parents []int, u int) int {
	parent := parents[u]
	// When no parent has been set for a vertice, then this is the parent.
	if parent == u {
		return u
	}
	return findParent(parents, parent)
}
