package main

import "fmt"

func main() {
	fibonacci(100)
}

func fibonacci(n int) {
	// generate the first n Fibonacci numbers
	for i := 0; i < n; i++ {
		fmt.Printf("%v ... ", FibonacciMemoised(i))
	}

}

func repeatingWords() {
		fmt.Println(FindRepeatingWords(
			"The rain poured and poured, making the streets flooded flooded",
		))
		fmt.Println(FindRepeatingWords(
			"We scream, you scream, we all scream for ice cream",
		))
	}
	