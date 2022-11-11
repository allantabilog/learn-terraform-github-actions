package main

import (
	"fmt"
)

func FibMain() {
	fmt.Println(fib(3))
}

// Calculates the first n terms of the fibonacci sequence
func fib(n int) []int {
	var series []int = make([]int, n, 100)

	if n == 1 {
		series[0] = 0
		return series
	}

	series[0], series[1] = 0, 1

	for i := 2; i < n; i++ {
		series[i] = series[i-1] + series[i-2]
	}

	return series

}
