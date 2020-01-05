package sort

import "testing"

func TestQuickSort(t *testing.T) {
	for _, test := range many_lists {
		test = clone(test)
		QuickSort(test.list)
		listMustBeSorted(t, test)
	}
}
