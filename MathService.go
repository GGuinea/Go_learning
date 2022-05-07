package main

import "math"

func getStdDev(t2 []int) float64 {
	mean := getMean(t2)
	var standardDeviation float64 = 0
	for _, i := range t2 {
		standardDeviation += math.Pow(float64(i)-mean, 2)
	}
	return math.Sqrt(standardDeviation / float64(len(t2)))
}

func getMean(t2 []int) float64 {
	var mean float64 = 0
	for _, i := range t2 {
		mean += float64(i)
	}
	return mean / float64(len(t2))
}
