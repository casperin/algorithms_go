package graph

func Boruvka(edges []WeightedEdge, numVertices int) []WeightedEdge {
	// Just like in Kruskal, we need to keep track of which vertices are
	// already connected.
	parents := make([]int, numVertices)

	// index: parents index, value: edges index, pointing to the cheapest edge
	// of all the edges belonging to that parent.
	cheapest := make([]int, numVertices)

	// Only used to keep track of the size of the different sub trees. This is
	// so we can connect smaller trees to larger trees.
	rank := make([]int, numVertices)

	// We have len(vectices) trees to start with, and we need to find the
	// connection. When this is down to 1, we are done.
	numTrees := numVertices

	for i := 0; i < numVertices; i++ {
		// See Kruskal on how parents are used.
		parents[i] = i
		// This is a list of indexes, and we need to distinguish no value from
		// zeroth value.
		cheapest[i] = -1
	}

	mst := []WeightedEdge{}

	for numTrees > 1 {
		for idx, edge := range edges {
			parentU := findParent(parents, edge.U)
			parentV := findParent(parents, edge.V)

			// Only consider connected vertices
			if parentU == parentV {
				continue
			}

			// For each edge of a tree, we want the cheapest (lowest weight).
			// To find this, we keep track of the index (with similar parent
			// index) that is the cheapest.

			cheapestU := cheapest[parentU]
			cheapestV := cheapest[parentV]

			if cheapestU == -1 || edge.Weight < edges[cheapestU].Weight {
				cheapest[parentU] = idx
			}

			if cheapestV == -1 || edge.Weight < edges[cheapestV].Weight {
				cheapest[parentV] = idx
			}
		}

		// Once we have, for each tree, the cheapest edge, we go through each
		// and add it to our mst updating their parent as we do. Since we
		// update their parent within the loop, we have to make sure that we
		// ignore any of the edges that now (after parents changed) will create
		// a cycle.
		for i, idx := range cheapest {
			// We need to reset this at some point anyway. The idx won't change
			// just because we modify the array.
			cheapest[i] = -1

			if idx == -1 {
				continue
			}

			edge := edges[idx]

			parentU := findParent(parents, edge.U)
			parentV := findParent(parents, edge.V)

			// If they are now connected (because union call on previous edge)
			// then we ignore it.
			if parentU != parentV {
				boruvkaUnion(parents, rank, parentU, parentV) // Unite vertices
				mst = append(mst, edge)                       // Edge is good
				numTrees--                                    // One less tree
			}
		}
	}

	return mst
}

func boruvkaUnion(parents, rank []int, parentU, parentV int) {
	// We want to attach smaller trees to larger ones.
	if rank[parentU] < rank[parentV] {
		parents[parentU] = parentV
	} else if rank[parentU] > rank[parentV] {
		parents[parentV] = parentU
	} else {
		parents[parentV] = parentU
		rank[parentU] += 1
	}
}
