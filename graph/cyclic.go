package graph

// Tags an adjacency list representing a directed graph and returns true if a
// cycle is found.
func IsCyclic(adj [][]int) bool {
	visited := make([]bool, len(adj))
	recursionStack := make([]bool, len(adj))

	for v := 0; v < len(adj); v++ {
		if isCycleUtil(v, visited, recursionStack, adj) {
			return true
		}
	}

	return false
}

func isCycleUtil(v int, visited, recursionStack []bool, adj [][]int) bool {
	if recursionStack[v] {
		return true
	}
	if visited[v] {
		return false
	}

	visited[v] = true
	recursionStack[v] = true

	for _, c := range adj[v] {
		if isCycleUtil(c, visited, recursionStack, adj) {
			return true
		}
	}

	recursionStack[v] = false

	return false
}
