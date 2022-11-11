package misc

import (
	"fmt"
	"math"
)

type flt func(int) bool

type Customer struct {
	ID   int
	Name string
}

func SetName(c *Customer) {
	c.Name = "Default"
}

func IsOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func Filter(slice []int, f flt) []int {
	var res []int
	for _, val := range slice {
		if f(val) {
			res = append(res, val)
		}
	}
	return res

}

func IsEvenV2(n int) bool {
	return n%2 == 0
}

func JoelMain() {
	for range [2]int{} {
		fmt.Println("Good stuff.")
	}
}

// function that takes a function as a parameter
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func anonymousFunc() {
	hypotenuse := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
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
