package statistics

import "testing"

func TestVarianceSample(t *testing.T) {
	sample := []f64{600.0, 470.0, 170.0, 430.0, 300.0}
	mu := Mean(sample)
	assert(t, VarianceSample(sample, mu), 27130.0)
}

func TestZScoreSingleSample(t *testing.T) {
	sample1 := 105.0
	mu1 := 100.0
	sigma1 := 4.0
	result1 := ZScoreSingleSample(sample1, mu1, sigma1)
	assert(t, result1, 1.25)

	sample2 := 100.0
	mu2 := 100.0
	sigma2 := 4.0
	result2 := ZScoreSingleSample(sample2, mu2, sigma2)
	assert(t, result2, 0.0)
}
