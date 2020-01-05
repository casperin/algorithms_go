package sort

import "testing"

func TestHeapSort(t *testing.T) {
	for _, test := range many_lists {
		test = clone(test)
		HeapSort(test.list)
		listMustBeSorted(t, test)
	}
}
