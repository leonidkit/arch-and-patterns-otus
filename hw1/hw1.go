package hw1

import (
	"errors"
	"math"
)

const float64EqualityThreshold = 1e-9

var errWrongArgumentZeroA = errors.New("аргумент a не может быть равен нулю")

func Solve(a, b, c float64) []float64 {
	d := math.Pow(b, 2) - 4*a*c
	if floatEqual(a, 0, float64EqualityThreshold) {
		panic(errWrongArgumentZeroA)
	}
	if floatEqual(d, 0, float64EqualityThreshold) {
		x1_2 := -b / 2 * a
		return []float64{x1_2}
	}
	if d < 0 {
		return []float64{}
	}
	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / 2 * a
		x2 := (-b - math.Sqrt(d)) / 2 * a
		return []float64{x1, x2}
	}
	return []float64{}
}

func floatEqual(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
