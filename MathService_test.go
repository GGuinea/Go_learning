package main

import (
	"testing"
)

func Test_Mean_WhenListIsEmptyShouldReturnZero(t *testing.T) {
	var xs = []int{}
	want := 0.0
	mean := getMean(xs)
	if want != mean {
		t.Fatalf("Should return: %f , returned: %f ", want, mean)
	}
}

func Test_Mean_WhenOneAndTwoShouldReturnOneAndHalf(t *testing.T) {
	var xs = []int{1, 2}
	want := 1.5
	mean := getMean(xs)
	if want != mean {
		t.Fatalf("Should return: %f , returned: %f ", want, mean)
	}
}

func Test_Mean_WhenDoubleZeroShouldReturnZero(t *testing.T) {
	var xs = []int{0, 0}
	want := 0.0
	mean := getMean(xs)
	if want != mean {
		t.Fatalf("Should return: %f , returned: %f ", want, mean)
	}
}

func Test_Stddev_WhenZeroShouldReturnZero(t *testing.T) {
	var xs = []int{0}
	want := 0.0
	stddev := getStdDev(xs)
	if want != stddev {
		t.Fatalf("Should return: %f , returned: %f ", want, stddev)
	}
}

func Test_Stddev_WhenDoubleZeroShouldReturnZero(t *testing.T) {
	var xs = []int{0, 0}
	want := 0.0
	stddev := getStdDev(xs)
	if want != stddev {
		t.Fatalf("Should return: %f , returned: %f ", want, stddev)
	}
}

func Test_Stddev_WhenOneAndTwoShouldReturnZeroAndHalf(t *testing.T) {
	var xs = []int{1, 2}
	want := 0.5
	stddev := getStdDev(xs)
	if want != stddev {
		t.Fatalf("Should return: %f , returned: %f ", want, stddev)
	}
}
