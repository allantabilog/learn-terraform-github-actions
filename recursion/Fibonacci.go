package main

import "fmt"

var cache []int64

// Fibonacci calculation
// that memoizes results
func Fibonacci(n int64) int64 {
	if n >= int64(len(cache)) {
		cache = append(cache, 
			Fibonacci(n - 1) + Fibonacci(n - 2),
		)
	}
	return cache[n]
}

func main2() {
	cache = make([]int64, 2)
	cache[0] = 0
	cache[1] = 1

	fmt.Println(Fibonacci(10))
	fmt.Printf("rebuilt cache %v\n", cache)
	fmt.Println(Fibonacci(7))
	fmt.Println(Fibonacci(30))
	fmt.Printf("rebuilt cache %v\n", cache)

	fmt.Println(Fibonacci(50))
	fmt.Printf("rebuilt cache %v\n", cache)
}

func main() {
	repeatCountDown(10, 5)
}

func countDown(n int) {
	fmt.Println(n)
	if n <= 1 {
			return
	}
	countDown(n - 1)
}

func repeatCountDown(n int, step int) {
	fmt.Println(n)
	if step > 10 {
		return
	}
	if n <= 1 {
		return 
	}
	for i:=0; i< 10; i++ {

		repeatCountDown(n - 1, step - 1)
	}
}

