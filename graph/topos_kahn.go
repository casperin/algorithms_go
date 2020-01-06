package graph

import (
	"errors"
)

var ErrGraphContainsCycle = errors.New("Graph contains cycle")

func ToposKahn(adj [][]int) ([]int, error) {
	inDegree := make([]int, len(adj))

	for _, children := range adj {
		for _, c := range children {
			inDegree[c] += 1
		}
	}

	queue := []int{}

	for i, _ := range adj {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	result := make([]int, 0, len(adj))

	for len(queue) > 0 {
		u := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		result = append(result, u)

		for _, c := range adj[u] {
			inDegree[c] -= 1
			if inDegree[c] == 0 {
				queue = append(queue, c)
			}
		}

		count += 1
	}

	if count < len(adj) {
		return nil, ErrGraphContainsCycle
	}
	return result, nil
}
