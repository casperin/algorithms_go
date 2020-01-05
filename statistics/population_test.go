package statistics

import "testing"

func TestVariancePopulation(t *testing.T) {
	population := []f64{600.0, 470.0, 170.0, 430.0, 300.0}
	mu := Mean(population)
	assert(t, VariancePopulation(population, mu), 21704.0)
}
