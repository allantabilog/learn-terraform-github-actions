package main

import "math"

// map to store memoised Fibonacci numbers
var fibs = map[int]int{}

// recursive function to calculate the nth Fibonacci number
// inefficient - repeats calculations
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// iterative function to calculate the nth Fibonacci number
func FibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	var a, b = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func FibonacciMemoised(n int) int {
	// check if the Fibonacci number is already memoised
	if val, ok := fibs[n]; ok {
		return val
	}
	// calculate the Fibonacci number
	if n <= 1 {
		fibs[n] = n
	} else {
		fibs[n] = FibonacciMemoised(n-1) + FibonacciMemoised(n-2)
	}
	if fibs[n] > math.MaxInt64 - 1 {
		panic("Fibonacci number is too large")
	}
	return fibs[n]
}