package sort

/// The classic slow sorting algorithm that you shouldn't use.
func BubbleSort(list []int) {
	size := len(list)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if list[j] > list[j+1] {
				swap(list, j, j+1)
			}
		}
	}
}
