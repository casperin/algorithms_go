package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	for _, test := range many_lists {
		test = clone(test)
		BubbleSort(test.list)
		listMustBeSorted(t, test)
	}
}
