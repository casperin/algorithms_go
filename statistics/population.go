package statistics

import "math"

// Functions dealing with populations

// Implementation of two pass algorithm for calculating the variance of a full
// population.
//
// Notice there is a difference between VariancePopulation and VarianceSample.
// This function uses N as divisor since it deals with the whole population,
// and not just a sample.
func VariancePopulation(data []f64, mu f64) f64 {
	tss := SumOfSquares(data, mu)
	n := len(data)
	return tss / f64(n)
}

// Difined as the square root of the deviation.
//
// Notice that that is a difference between StandardDeviationPopulation and
// StandardDeviationSample, as they make use of different variance functions.
func StandardDeviationPopulation(data []f64, mu f64) f64 {
	return math.Sqrt(VariancePopulation(data, mu))
}
