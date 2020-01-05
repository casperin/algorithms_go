package sort

import "testing"

func TestMergeSort(t *testing.T) {
	for _, test := range many_lists {
		test = clone(test)
		MergeSort(test.list)
		listMustBeSorted(t, test)
	}
}
