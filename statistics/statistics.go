package statistics

import (
	"math"
	"sort"
)

type f64 = float64

type Coord struct {
	X f64
	Y f64
}

type ChiSquaredItem struct {
	Expected f64
	Observed f64
}

func panicEmpty(data []f64) {
	if len(data) == 0 {
		panic("Data set can not be empty")
	}
}

// Calculates the mean, or the average, of a list of floats.
// Panics if dataset is empty.
func Mean(data []f64) f64 {
	sum := 0.0
	for _, n := range data {
		sum += n
	}
	return sum / f64(len(data))
}

// Calculates the median.
//
// The function needs to clone and sort the dataset which is expensive, so if
// you know that your dataset is sorted, then use MedianFromSorted instead.
//
// Panics if dataset is empty.
func Median(data []f64) f64 {
	sort.Float64s(data)
	return MedianFromSorted(data)
}

// Calculates the median from a sorted list.
// Panics if dataset is empty.
func MedianFromSorted(data []f64) f64 {
	panicEmpty(data)

	size := len(data)
	halfIndex := (size - 1) / 2

	if size%2 == 1 {
		return data[halfIndex]
	}

	return (data[halfIndex] + data[halfIndex+1]) / 2.0
}

// Finds the mode of a dataset.
// Panics if dataset is empty.
func Mode(data []f64) f64 {
	nums := map[f64]int{}

	for _, n := range data {
		nums[n] = nums[n] + 1
	}

	maxNum := 0.0
	maxCount := 0

	for n, count := range nums {
		if count > maxCount {
			maxCount = count
			maxNum = n
		}
	}

	return maxNum
}

// Calculates the arithmetic range, and coefficient of range, of a dataset.
// Returns range, the coefficient of range, and the largest and smallest value.
// Panics if dataset is empty.
func Range(data []f64) (f64, f64, f64, f64) {
	panicEmpty(data)

	largest := data[0]
	smallest := data[0]

	for _, n := range data {
		if n > largest {
			largest = n
		}
		if n < smallest {
			smallest = n
		}
	}

	rang := largest - smallest
	coefOfRange := rang / (largest + smallest)

	return rang, coefOfRange, smallest, largest
}

// Pearson correlation coefficient. This is what is normally referred to when
// talking about finding the correlation. Returns a number beween -1 and 1.
func Correlation(data []Coord) f64 {
	n := f64(len(data))
	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumX2 := 0.0
	sumY2 := 0.0

	for _, c := range data {
		sumX += c.X
		sumY += c.Y
		sumXY += c.X * c.Y
		sumX2 += math.Pow(c.X, 2.0)
		sumY2 += math.Pow(c.Y, 2.0)
	}

	dividend := n*sumXY - sumX*sumY
	divisorLeft := n*sumX2 - math.Pow(sumX, 2.0)
	divisorRight := n*sumY2 - math.Pow(sumY, 2.0)
	divisor := math.Sqrt(divisorLeft * divisorRight)

	return dividend / divisor
}

// Returns the sum, over all observations, of the squared differences of each
// observation from the overall mean.
func SumOfSquares(data []f64, mean f64) f64 {
	sum := 0.0
	for _, n := range data {
		sum += math.Pow(n-mean, 2.0)
	}
	return sum
}

// Chi-squared
func ChiSquared(data []ChiSquaredItem) f64 {
	sum := 0.0
	for _, item := range data {
		sum += math.Pow(item.Expected-item.Observed, 2.0) / item.Expected
	}
	return sum
}
