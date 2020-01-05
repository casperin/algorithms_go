package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	for _, test := range many_lists {
		test = clone(test)
		InsertionSort(test.list)
		listMustBeSorted(t, test)
	}
}
