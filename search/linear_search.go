package search

// For completeness. Only use if list is very very small.
func LinearSearch(list []int, x int) int {
	for i, item := range list {
		if item == x {
			return i
		}

		if item > x {
			return -1
		}
	}

	return -1
}
