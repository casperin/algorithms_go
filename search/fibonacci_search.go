package search

// Uses the lower of a fibonacci triple (eg. 3, 5, and 8) as the middle point in a binary search.
// In practice fibonacci Search search is rarely faster than a binary search.
//
// Reasons that I have seen for using fibonacci search:
//
// 1. Fibonacci Search divides given array in unequal parts
// 2. Binary Search uses division operator to divide range. Fibonacci Search doesn’t use /, but
//    uses + and -. The division operator may be costly on some CPUs.
// 3. Fibonacci Search examines relatively closer elements in subsequent steps. So when input
//    array is so big that it cannot fit in CPU cache or even in RAM, Fibonacci Search can be
//    useful.
//
// Also this answer on Stack Overflow:
// https://stackoverflow.com/questions/22877763#answer-22877947
func FibonacciSearch(list []int, x int) int {
	size := len(list)
	fib2 := 0
	fib1 := 1
	fib := fib2 + fib1

	if size == 0 {
		return -1
	}

	for fib < size {
		fib2 = fib1
		fib1 = fib
		fib = fib2 + fib1
	}

	offset := -1

	for fib > 1 {
		index := min(offset+fib2, size-1)

		if list[index] == x {
			return index
		}

		if list[index] < x {
			fib = fib1
			fib1 = fib2
			fib2 = fib - fib1
			offset = index
		} else {
			fib = fib2
			fib1 = fib1 - fib2
			fib2 = fib - fib1
		}
	}

	if fib1 == 1 && list[offset+1] == x {
		return offset + 1
	}

	return -1
}
