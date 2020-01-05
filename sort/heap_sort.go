package sort

// * Time complexity: O(n*logn)
// * Sorting in place: Yes.
// * Stable: No.
//
// Works by "creating" a heap, or tree, out of the list and then move them around to make a tree
// where all children are lower than their parents. Once this is done, it takes the root element
// (which is the highest value) and swaps it with the last value. It then has the first element
// in its place (highest at the back), so it cuts off that part of the list and never touches it
// again. Then it lets the new root value (which is very unlikely to be the highest value) walk
// down the tree, until the heap again is in a state of "all parents are higher than their
// chilren". Then it repeats the process.
//
// The way the algorithm structures its heap is described below. a, b, c, ... refers to (unordered
// items), and it just decides that b and c are children of a, and so on.
//
// vector:  [a, b, c, d, e, f, g, h, i, j, k, l, m, n]
// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13]
// length:  14
// mid:     7 (h)
// tree:
//               a
//         b           c
//      d     e     f     g
//     h i   j k   l m   n
//
// To find children for item with index i:
//
//  left: i * 2 + 1    // eg. c: 2 * 2 + 1 = 5 (f)
// right: i * 2 + 2    // eg. e: 4 * 2 + 2 = 10 (k)
func HeapSort(list []int) {
	size := len(list)

	// Build max heap
	// After this we know that a > b and a > c (and so on downwards), but we don't know if d > m
	for i := int(size/2) - 1; i >= 0; i-- {
		heapify(list, size, i)
	}

	// We take the value at the top (we know this is the highest number) and position it last, and
	// then we "cut the last off" and let the new value at the top (which should be a pretty low
	// value) walk downwards, but not all the way to the last position (the cut off part).
	//
	// Then we repeat, but ignore the cut-off part.
	//
	// Looking at the example above, we swap the value of a with n and [a,...,m] becomes our tree,
	// and [n] is the sorted vector. Then we repeat, so the value of a swaps with m and [a,...,l]
	// is our tree, and [m,n] is the sorted vector.
	for i := size - 1; i >= 0; i-- {
		swap(list, 0, i)    // position highest value at the "end"
		heapify(list, i, 0) // let root element walk down
	}
}

func heapify(list []int, n int, root int) {
	largest := root // eg. 1 = b
	l := root*2 + 1 // 3 = d
	r := root*2 + 2 // 4 = e

	if l < n && list[l] > list[largest] {
		largest = l
	}

	// eg. e is bigger, so largest = 4
	if r < n && list[r] > list[largest] {
		largest = r
	}

	// If root is not largest, swap with largest and continue heapifying
	if largest != root {
		swap(list, root, largest) // we swap the value
		heapify(list, n, largest) // and call with 'e' as root
	}
}
