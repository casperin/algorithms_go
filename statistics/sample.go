package statistics

import "math"

// Implementation of two pass algorithm for calculating the variance of a
// sample.
//
// Notice there is a difference between VarianceSample and VariancePopulation.
// This function uses N-1 as divisor because it deals with a sample.
func VarianceSample(data []f64, mu f64) f64 {
	tss := SumOfSquares(data, mu)
	n := len(data)
	return tss / f64(n-1)
}

// Difined as the square root of the deviation.
//
// Notice that that is a difference between StandardDeviationSample and
// StandardDeviationPopulation, as they make use of different variance
// functions.
func StandardDeviationSample(data []f64, mu f64) f64 {
	return math.Sqrt(VarianceSample(data, mu))
}

func ZScore(muSample, muPopulation, sigmaPopulation f64, nSample int64) f64 {
	dividend := muSample - muPopulation
	divisor := sigmaPopulation / math.Sqrt(f64(nSample))
	return dividend / divisor
}

func ZScoreSingleSample(sample, mu, sigma f64) f64 {
	return ZScore(sample, mu, sigma, 1)
}
