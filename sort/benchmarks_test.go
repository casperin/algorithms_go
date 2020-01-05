package sort

import (
	"math/rand"
	"testing"
)

var randSeeded = false

// Helper
func cloneList(list []int) []int {
	var out = make([]int, len(list))
	for i, n := range list {
		out[i] = n
	}
	return out
}

// Helper
func benchmarkSort(sort func([]int), list []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sort(cloneList(list))
	}
}

////////////////////////////////////
// Already sorted list benchmarks //
////////////////////////////////////

// cached list. We do it this way, so all functions get the same list (more
// relevant when they contain random numbers)
var sortedArray = make([]int, 1000)
var sortedArrayDone = false

func getSortedArray() []int {
	if !sortedArrayDone {
		for i := 0; i < 1000; i++ {
			sortedArray[i] = i
		}
		sortedArrayDone = true
	}
	return sortedArray
}

func Benchmark_Sorted_BubbleSort(b *testing.B) {
	benchmarkSort(BubbleSort, getSortedArray(), b)
}

func Benchmark_Sorted_HeapSort(b *testing.B) {
	benchmarkSort(HeapSort, getSortedArray(), b)
}

func Benchmark_Sorted_InsertionSort(b *testing.B) {
	benchmarkSort(InsertionSort, getSortedArray(), b)
}

func Benchmark_Sorted_MergeSort(b *testing.B) {
	benchmarkSort(MergeSort, getSortedArray(), b)
}

func Benchmark_Sorted_QuickSort(b *testing.B) {
	benchmarkSort(QuickSort, getSortedArray(), b)
}

////////////////////////////////////
// Reverse sorted list benchmarks //
////////////////////////////////////

var reverseSortedArray = make([]int, 1000)
var reverseSortedArrayDone = false

func getReverseSortedArray() []int {
	if !reverseSortedArrayDone {
		for i := 0; i < 1000; i++ {
			reverseSortedArray[i] = 999 - i
		}
		reverseSortedArrayDone = true
	}
	return reverseSortedArray
}

func Benchmark_ReverseSorted_BubbleSort(b *testing.B) {
	benchmarkSort(BubbleSort, getReverseSortedArray(), b)
}

func Benchmark_ReverseSorted_HeapSort(b *testing.B) {
	benchmarkSort(HeapSort, getReverseSortedArray(), b)
}

func Benchmark_ReverseSorted_InsertionSort(b *testing.B) {
	benchmarkSort(InsertionSort, getReverseSortedArray(), b)
}

func Benchmark_ReverseSorted_MergeSort(b *testing.B) {
	benchmarkSort(MergeSort, getReverseSortedArray(), b)
}

func Benchmark_ReverseSorted_QuickSort(b *testing.B) {
	benchmarkSort(QuickSort, getReverseSortedArray(), b)
}

//////////////////////////////////////
// Random numbers between 0 and 100 //
//////////////////////////////////////

var random1000Array = make([]int, 1000)
var random1000ArrayDone = false

func getRandom1000Array() []int {
	if !randSeeded {
		rand.Seed(42)
		randSeeded = true
	}
	if !random1000ArrayDone {
		for i := 0; i < 1000; i++ {
			random1000Array[i] = rand.Intn(100)
		}
		random1000ArrayDone = true
	}
	return random1000Array
}

func Benchmark_Random1000_BubbleSort(b *testing.B) {
	benchmarkSort(BubbleSort, getRandom1000Array(), b)
}

func Benchmark_Random1000_HeapSort(b *testing.B) {
	benchmarkSort(HeapSort, getRandom1000Array(), b)
}

func Benchmark_Random1000_InsertionSort(b *testing.B) {
	benchmarkSort(InsertionSort, getRandom1000Array(), b)
}

func Benchmark_Random1000_MergeSort(b *testing.B) {
	benchmarkSort(MergeSort, getRandom1000Array(), b)
}

func Benchmark_Random1000_QuickSort(b *testing.B) {
	benchmarkSort(QuickSort, getRandom1000Array(), b)
}

////////////////
// Small list //
////////////////

func Benchmark_8Numbers_BubbleSort(b *testing.B) {
	list := []int{3, 6, 21, 1, 2, 5, 2, 1}
	benchmarkSort(BubbleSort, list, b)
}

func Benchmark_8Numbers_HeapSort(b *testing.B) {
	list := []int{3, 6, 21, 1, 2, 5, 2, 1}
	benchmarkSort(HeapSort, list, b)
}

func Benchmark_8Numbers_InsertionSort(b *testing.B) {
	list := []int{3, 6, 21, 1, 2, 5, 2, 1}
	benchmarkSort(InsertionSort, list, b)
}

func Benchmark_8Numbers_MergeSort(b *testing.B) {
	list := []int{3, 6, 21, 1, 2, 5, 2, 1}
	benchmarkSort(MergeSort, list, b)
}

func Benchmark_8Numbers_QuickSort(b *testing.B) {
	list := []int{3, 6, 21, 1, 2, 5, 2, 1}
	benchmarkSort(QuickSort, list, b)
}
