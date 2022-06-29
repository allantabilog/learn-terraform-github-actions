package main

import "fmt"
import "sandbox/goroutines/mergesort/merge"

func main() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}

	fmt.Printf("%v\n%v", data, merge.MergeSort(data))
}
