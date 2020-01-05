package sort

func MergeSort(list []int) {
	MergeSortSection(list, 0, len(list))
}

func MergeSortSection(list []int, start, end int) {
	if start+1 >= end {
		return // noop if less than two items
	}

	mid := int((start + end) / 2)

	// Split vector into left and right and sort those recursively
	MergeSortSection(list, start, mid) // sort left side
	MergeSortSection(list, mid, end)   // sort right side

	left_index := start
	right_index := mid
	buffer := []int{}

	// While we have items in both left and right vector we push the lowest of them onto the buffer
	for left_index < mid && right_index < end {
		left_item := list[left_index]
		right_item := list[right_index]

		if left_item < right_item {
			buffer = append(buffer, left_item)
			left_index += 1
		} else {
			buffer = append(buffer, right_item)
			right_index += 1
		}
	}

	// Push whatever is left of the left (or the right) vector
	for left_index < mid {
		buffer = append(buffer, list[left_index])
		left_index += 1
	}

	for right_index < end {
		buffer = append(buffer, list[right_index])
		right_index += 1
	}

	// Items in the buffer are sorted, so we update the section of the original array that we are
	// dealing with to reflect the buffer
	for i, item := range buffer {
		list[i+start] = item
	}
}
