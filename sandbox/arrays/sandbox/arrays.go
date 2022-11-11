package sandbox

import (
	"fmt"
)

type Array interface {
	Index(i int) int
}

type array struct {
	elements []int
}

func (arr *array) Index(i int) int {
	return arr.elements[i]
}

func Main() {
	sliceTests3()

}

func sliceTests3() {
	arr := array{
		elements: []int{1, 2, 3},
	}
	fmt.Println(arr.Index(0))
}

func sliceTests2() {
	a := make([]int, 5)

	fmt.Println(len(a) == 5)
	fmt.Println(cap(a) == 5)
}

func sliceTests1() {
	var a [10]int

	for i := 0; i < 10; i++ {
		a[i] = i * 2
	}

	fmt.Println(a)
	fmt.Println(a[:])
	fmt.Println(a[:10])
	fmt.Println(a[0:])

}
func InsertSlice(slice, sliceToInsert []string, index int) []string {
	testSlice := []string{"0", "1", "2", "3"}
	//testSliceToInsert := []string{"A", "B", "C"}

	newSlice := make([]string, len(testSlice), cap(testSlice))

	copy(newSlice, testSlice[1:index])

	return newSlice
}

func Main2() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	// s = s[:6] 	// slice bounds out of range [:6] with capacity 4
	// printSlice(s)

	var ss []int
	printSlice(ss)

	a := make([]int, 5)
	printSlice(a)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func Main_Slices() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}
func Main_ArraysSlices() {
	fmt.Println("this should work")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// a slice
	var s []int = primes[1:4]
	fmt.Println(s)
}
