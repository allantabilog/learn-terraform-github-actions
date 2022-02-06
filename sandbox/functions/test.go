package main

import (
	"fmt"
	"math"
)

// function that takes a function as a parameter
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func main() {
	hypotenuse := func(x, y float64) float64 {
		return math.Sqrt(x * x + y * y)
	}
	fmt.Println(hypotenuse(5, 12))
	fmt.Println(compute(hypotenuse))
	fmt.Println(compute(math.Pow))
}



func greeting(name string) string {
	return "Hello.there" + name
}

func add(a, b int) int {
	return a + b
}

func main_0() {
	fmt.Println(greeting("Brad.Pitt"))
	fmt.Println(add(3, 4))
}
