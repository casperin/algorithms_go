package sort

// For small or nearly sorted lists. This is often used in conjunction with other sorts, when
// dealing with small lists.
//
// Works by taking each element (starting with second element) and moving it towards the front of
// the list until it reaches an element that is lower than itself.
//
// Best case (an already sorted list) is O(n).
func InsertionSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i - 1; j >= 0; j-- {
			if list[j] <= list[j+1] {
				break
			}
			swap(list, j, j+1)
		}
	}
}
