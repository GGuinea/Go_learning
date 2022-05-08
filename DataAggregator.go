package main

type DataAggregator struct {
	StandardDeviation float64 `json:"stddev"`
	NumList           []int   `json:"data"`
}
