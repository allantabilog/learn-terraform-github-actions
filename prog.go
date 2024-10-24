package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	dynamicSlice()
}

// use append function to dynamically re-allocate arrays
// https://pkg.go.dev/builtin#append
func dynamicSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	s = append(s, 0)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

	for i, v := range s {
		fmt.Println(i, v)
	}

	var nums []int
	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(1000))
		fmt.Println(len(nums), cap(nums), nums)
	}

}

func nestedSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][1] = "O"

	showBoard(board)
}

func showBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func sliceLiterals() {
	// this creates an array
	// and a slice that references it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, false},
	}
	fmt.Println(s)
	fmt.Println(q[0:])
	fmt.Println(q[:3])
	q = q[:3]
	fmt.Println(q, len(q), cap(q))
	// extend the length of the slice
	q = q[:5]
	fmt.Println(q, len(q), cap(q))
	// below line will error with
	// slice bounds out of range
	//q = q[:7]

	a := make([]int, 5, 10)
	fmt.Println(a, len(a), cap(a))

}

func basics() {
	v0 := Vertex{}
	v1 := Vertex{X: 1}
	v2 := Vertex{Y: 1}

	fmt.Printf("%+v\n", v0)
	fmt.Printf("%+v\n", v1)
	fmt.Printf("%+v\n", v2)

	var a [10]int // array: fixed-size, length is part of the type definition
	a[0] = 100
	a[1], a[2] = 10, 20
	a[7], a[8], a[9] = 7, 8, 9
	fmt.Println(a)

	// a slice is a view into the elements of an array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
	// slices are just views
	// slices don't store data
	// slices can have aliases
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	n1 := names[0:2]
	n2 := names[1:3]

	fmt.Println(n1, n2)
	n1[1] = "Updated name"

	fmt.Println(names, n1, n2)

}
