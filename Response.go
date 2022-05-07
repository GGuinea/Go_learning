package main

type Response struct {
	StandardDeviation float64 `json:"stddev"`
	NumList           []int   `json:"data"`
}
