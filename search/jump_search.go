package search

import "math"

// Jump Search is a glorified linear search. It is only useful if you know that the item you are
// looking for is among the first ones in the list. If that isn't the case, then binary search
// will have much better performance.
//
// It checks the first number for the value. If not there, it jumps sqrt(list.length) forward and
// checks again. When the checked value is lower than what we are searching for, we know we passed
// it. So we do a linear search (from the previous value up to the current value) to find it.
func JumpSearch(list []int, x int) int {
	size := len(list)
	step := int(math.Sqrt(float64(size)))
	prev := 0

	if size == 0 {
		return -1
	}

	for list[min(step, size)-1] < x {
		prev = step
		step += step
		if prev >= size {
			return -1
		}
	}

	for end := min(step, size); prev < end; prev++ {
		if list[prev] == x {
			return prev
		}
	}

	return -1
}
