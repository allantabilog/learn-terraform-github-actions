package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// the worker goroutine
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Recieved job ", j)
			} else {
				fmt.Println("Recieved all jobs.")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job ", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs, closed channel.")

	<-done

}
