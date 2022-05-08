package main

import (
	"testing"
)

func Test_WhenListIsEmptyShouldReturnZero(t *testing.T) {
	var xs = []int{}
	want := 0.0
	mean := getMean(xs)
	if want != mean {
		t.Fatalf("Should return: %f , returned: %f ", want, mean)
	}
}

func Test_WhenOneAndTwoShouldReturnOneAndHalf(t *testing.T) {
	var xs = []int{1, 2}
	want := 1.5
	mean := getMean(xs)
	if want != mean {
		t.Fatalf("Should return: %f , returned: %f ", want, mean)
	}
}

func Test_WhenDoubleZeroShouldReturnZero(t *testing.T) {
	var xs = []int{0, 0}
	want := 0.0
	mean := getMean(xs)
	if want != mean {
		t.Fatalf("Should return: %f , returned: %f ", want, mean)
	}
}
