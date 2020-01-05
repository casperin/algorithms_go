package graph

// Returns the shortest Euler Circuit of a directed graph.
// Returns the wrong result if no path is possible.
func Hierholzers(adjacencyList [][]int, src int) []int {
	if len(adjacencyList) == 0 {
		return []int{}
	}

	// Contains the number of edges for a given vertice.
	edgeCount := make([]int, len(adjacencyList))
	for i := 0; i < len(adjacencyList); i++ {
		edgeCount[i] = len(adjacencyList[i])
	}

	currentPath := []int{src}
	circuit := []int{}
	currentVertice := src

	for len(currentPath) > 0 {
		if edgeCount[currentVertice] > 0 {
			// If the vertice has any connections then we keep going deeper and
			// deeper down edges of vertices
			currentPath = append(currentPath, currentVertice)                         // push it to currentPath
			edgeCount[currentVertice] -= 1                                            // deduct edges to check
			currentVertice = adjacencyList[currentVertice][edgeCount[currentVertice]] // set current vertice to be the last of its edges

		} else {
			// We've depleted this vertice's edges, so we add it to our circuit
			// (we are deep though, so it'll be in reverse order)
			circuit = append(circuit, currentVertice)        // append it to our curcuit
			currentVertice = currentPath[len(currentPath)-1] // last element of currentPath
			currentPath = currentPath[:len(currentPath)-1]   // [a,b,c] -> [a,b] (remove last element)
		}
	}

	// Reverse circuit
	for i, j := 0, len(circuit)-1; i < j; i, j = i+1, j-1 {
		circuit[i], circuit[j] = circuit[j], circuit[i]
	}

	return circuit
}
