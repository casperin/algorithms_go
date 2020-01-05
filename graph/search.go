package graph

// Implementation of BFS for a directed graph.
func (g *DirectedGraph) BFS(u int) []int {
	// A list of visited vertices, with the starting point already set to true.
	visited := make([]bool, len(g.List))
	visited[u] = true

	// A queue of vertices to visit, with the starting point added already.
	queue := []int{u}

	// We need to visit each of a vertice's connections before we visit any
	// connection's connections. We do this by traversing through the immediate
	// connections, adding them to a queue if they are not already there. Then
	// repeat that process for each connection added, for as long as we have
	// anything in the queue.
	for i := 0; i < len(queue); i++ {
		for _, v := range g.List[queue[i]] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	return queue
}

// Implementation of DFS for a directed grapth.
func (g *DirectedGraph) DFS(u int) []int {
	visited := make([]bool, len(g.List))
	path := []int{}
	return g.dfs(u, visited, path)
}

// For a DFS we need to look at the connections of a vertice before we look any
// other vertices. To do this, we recursively visit one connection of a vertice
// (again, only connections not already dealt with) before we move on to the
// next connection.
func (g *DirectedGraph) dfs(u int, visited []bool, path []int) []int {
	visited[u] = true
	path = append(path, u) // add this vertice

	for _, v := range g.List[u] {
		if !visited[v] {
			// add the entirety of this connection (and its connections),
			// before dealing with the next
			path = g.dfs(v, visited, path)
		}
	}

	return path
}
