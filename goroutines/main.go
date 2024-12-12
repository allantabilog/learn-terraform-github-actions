package main

import (
	"fmt"
	"sync"
)

func sumIntegers(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
func sumGoRoutine(inputSlice []int, c chan int) {
 //Write your code here
 sum := sumIntegers(inputSlice)
 // write the sum to the channel
 c <- sum
}
func findParallelSum(items [][]int) int{
	c := make(chan int)
	totalSum := 0
 	for _, itemList := range items {
		go sumGoRoutine(itemList, c)
 	}
	for i:=0; i < len(items); i++ {
		totalSum += <-c
	}
	return totalSum
}

func main() {
	// declare a 2D slice of integers
	items := [][]int{
		[]int{1, 2, 3},
		[]int{10, 20, 30},
		[]int{100, 200, 300},
	}

	// call the findParallelSum function
	fmt.Println(findParallelSum(items))

}
func mainChannel3() {

	itemsA := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	var totalSum int
	// declare the wait group
	var wg sync.WaitGroup
	
	
	for _, items := range itemsA {
		wg.Add(1)
		go func(itemList []int) {
			defer wg.Done()
			fmt.Println("Processing items: ", itemList)	
			var sum int = add(itemList...)
			fmt.Println("Sum of items: ", sum)
			totalSum += sum
		}(items)
	}

	wg.Wait()
	fmt.Println("Total sum: ", totalSum)
}

// declare a function that takes a variable number of integers
// and adds them up
func add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func mainChannel() {
	ch := make(chan string, 1)

	go func() {
		for _, word := range []string{"hello", "world"} {
			ch <- word
		}
		close(ch)
	}()

	for word := range ch {
		fmt.Println(word)
	}

	// this will return zero value of the channel type
	var zero = <- ch 
	fmt.Printf("Read from closed channel: [%v]\n", zero)

	ch <- "Panic" // This will panic
					// cos sending to a closed channel
}

func mainChannel2() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("Processing: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("All done")
}