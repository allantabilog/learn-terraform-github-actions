package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(triple(1, 2))
	fmt.Println(triple(1, -2))
	fmt.Println(triple(-1, 2))
	fmt.Println(triplev2(1, 2))
	fmt.Println(triplev2(1, -2))
	fmt.Println(triplev2(-1, 2))

}

func triple(a, b int) (sum, prod, diff int) {
	sum = a + b
	prod = a * b
	diff = a - b
	return
}

func triplev2(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func errorTest() {
	fmt.Println(Hello("world"))

	result, err := Hello("")
	if err != nil {
		fmt.Printf("Error calling Hello(): %v", err)
	}
	fmt.Println(result)
}

func div(n int, m int) (bool, error) {
	if m == 0 {

	}
	return false, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(fmt.Sprintf("[%v] is an empty name", name))
	}

	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message, nil
}
