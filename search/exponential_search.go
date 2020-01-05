package search

// Starts by checking the first value, then cuts off bigger and bigger chunks (exponentially
// growing in size) of the list. Once it finds that the cut-off value is bigger than the one
// searched for, it performs a binary search on the chunk.
//
// Exponential searching is useful in two circumstances:
// 1. The list is unbounded / infinite (not true in Go, at least for this implementation)
// 2. The searched for element is near the beginning of the list
func ExponentialSearch(list []int, x int) int {
	size := len(list)

	if size == 0 {
		return -1
	}

	if list[0] == x {
		return 0
	}

	bound := 1

	for bound < size && list[bound] <= x {
		bound *= 2
	}

	lo := int(bound / 2)
	hi := min(bound, size)

	for lo < hi {
		mid := int((lo + hi) / 2)

		if list[mid] == x {
			return mid
		}

		if list[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return -1
}
