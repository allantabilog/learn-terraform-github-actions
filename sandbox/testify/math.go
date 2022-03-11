package math

import "math"

func Add(n1 int, n2 int) int {
	return n1 + n2
}

func Multiply(n1 int, n2 int) int {
	return n1 * n2
}

func Abs(n int) float64 {
	return math.Abs(float64(n))
}
