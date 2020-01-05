package search

// Your classic binary search on sorted lists. Works by cutting the list in
// half, checking if it's a match. If not, repeats on either the upper or lower
// half.
func BinarySearch(list []int, x int) int {
	lo := 0
	hi := len(list)

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
