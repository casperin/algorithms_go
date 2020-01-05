package search

// A variation of the binary search, where it looks at how big the value is and then tries to
// search in the relevant area of the list. Say, the first item in the list is 0 and the last is
// 100. If searching for 99, then it'll start by searching towards the end first.
func InterpolationSearch(list []int, x int) int {
	lo := 0
	hi := len(list) - 1

	for lo <= hi && x >= list[lo] && x <= list[hi] {
		if lo == hi {
			if list[lo] == x {
				return lo
			}
			return -1
		}

		index := lo + (((hi - lo) / (list[hi] - list[lo])) * (x - list[lo]))

		if list[index] == x {
			return index
		}

		if list[index] < x {
			lo = index + 1
		} else {
			hi = index - 1
		}
	}

	return -1
}
