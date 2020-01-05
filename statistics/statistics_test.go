package statistics

import "testing"

func TestMean(t *testing.T) {
	data := []float64{1.0, 3.0, 3.0, 2.0, 1.0}
	assert(t, Mean(data), 2.0)
}

func TestMedian(t *testing.T) {
	data := []float64{2.0, 5.0, 1.0}
	assert(t, Median(data), 2.0)

	input2 := []float64{2.0, 5.0, 3.0, 1.0}
	assert(t, Median(input2), 2.5)
}

func TestMode(t *testing.T) {
	data := []float64{2.0, 5.0, 1.0, 3.0, 1.0}
	assert(t, Mode(data), 1.0)
}

func TestRange(t *testing.T) {
	data := []float64{89.0, 73.0, 84.0, 91.0, 87.0, 77.0, 94.0}
	rang, coef, smallest, largest := Range(data)
	assert(t, rang, 21.0)
	assertIsh(t, coef, 0.1257, 0.0001)
	assert(t, smallest, 73.0)
	assert(t, largest, 94.0)
}

func TestCorrelation(t *testing.T) {
	data := []Coord{
		Coord{43.0, 99.0},
		Coord{21.0, 65.0},
		Coord{25.0, 79.0},
		Coord{42.0, 75.0},
		Coord{57.0, 87.0},
		Coord{59.0, 81.0},
	}
	assertIsh(t, Correlation(data), 0.529809, 0.0000001)
}

func TestChiSquared(t *testing.T) {
	// Test 1
	data1 := []ChiSquaredItem{
		ChiSquaredItem{21.33333334, 29.0},
		ChiSquaredItem{21.33333334, 24.0},
		ChiSquaredItem{21.33333334, 22.0},
		ChiSquaredItem{21.33333334, 19.0},
		ChiSquaredItem{21.33333334, 21.0},
		ChiSquaredItem{21.33333334, 18.0},
		ChiSquaredItem{21.33333334, 19.0},
		ChiSquaredItem{21.33333334, 20.0},
		ChiSquaredItem{21.33333334, 23.0},
		ChiSquaredItem{21.33333334, 18.0},
		ChiSquaredItem{21.33333334, 20.0},
		ChiSquaredItem{21.33333334, 23.0},
	}
	assertIsh(t, ChiSquared(data1), 5.09375, 0.0000001)

	// Test 2
	data2 := []ChiSquaredItem{
		ChiSquaredItem{25.0, 23.0},
		ChiSquaredItem{16.0, 20.0},
		ChiSquaredItem{4.0, 3.0},
		ChiSquaredItem{24.0, 24.0},
		ChiSquaredItem{8.0, 10.0},
	}
	assert(t, ChiSquared(data2), 1.91)
}

func assert(t *testing.T, output, expected float64) {
	t.Helper()
	if output != expected {
		t.Fatalf("%v != %v (output|expected)", output, expected)
	}
}

func assertIsh(t *testing.T, output, expected, tolerance float64) {
	t.Helper()
	diff := output - expected
	if diff < 0 {
		diff = -diff
	}
	if diff > tolerance {
		t.Fatalf(
			"%v != %v (output|expected). Diff %v > %v (tolerance)",
			output, expected, diff, tolerance)
	}
}
