package sort

// * Time complexity: Best & Avg: O(n*logn), Worst: O(n*n).
// * Space complexity: O(n*logn) in average case.
// * Sorting in place: Yes.
// * Stable: No.<Paste>
func QuickSort(list []int) {
	QuickSortBound(list, 0, len(list)-1)
}

func QuickSortBound(list []int, lo, hi int) {
	if lo >= hi {
		return // we're done
	}

	pi := partition(list, lo, hi)

	QuickSortBound(list, lo, pi-1)
	QuickSortBound(list, pi+1, hi)
}

func partition(list []int, lo, hi int) int {
	pivot := list[hi]
	i := lo - 1

	for j := lo; j <= hi-1; j++ {
		if list[j] < pivot {
			i++
			swap(list, i, j)
		}
	}

	swap(list, i+1, hi)

	return i + 1
}
