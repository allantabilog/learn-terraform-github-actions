package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println(even(2))
	fmt.Println(odd(3))
	fmt.Println(odd(6))
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func trace(s string) (n int, err error) {

	defer func() {
		log.Printf("trace(%q) = %d, %v", s, n, err)
	}()

	return calculate(s)

}

func calculate(s string) (n int, err error) {
	if strings.Contains(s, "ell") {
		return 100, nil
	} else {
		return 0, io.EOF
	}
}

func even(n int) (res bool) {

	defer func() {
		log.Printf("even(%v) =  %v", n, res)
	}()

	if n == 0 {
		res = true
		return res
	}
	return odd(abs(n) - 1)
}
func odd(n int) (res bool) {

	defer func() {
		log.Printf("odd(%v) =  %v", n, res)
	}()

	if n == 0 {
		res = false
		return res
	}
	return even(abs(n) - 1)
}

func ello() {
	arguments := os.Args
	for i := 0; i < len(arguments); i++ {
		fmt.Println("ehlo: ere is an argument for ya: ")
		fmt.Printf("[%d]: %v\n", i, arguments[i])
	}
}
