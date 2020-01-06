package graph

// Implementation of topological sorting via depth first search.
func ToposDFS(adj [][]int) []int {
	stack := []int{}
	visited := make([]bool, len(adj))

	for i, v := range visited {
		if !v {
			stack = toposDFSUtil(i, stack, visited, adj)
		}
	}

	// Reverse stack
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}

func toposDFSUtil(v int, stack []int, visited []bool, adj [][]int) []int {
	visited[v] = true

	for _, u := range adj[v] {
		if !visited[u] {
			stack = toposDFSUtil(u, stack, visited, adj)
		}
	}

	return append(stack, v)
}
