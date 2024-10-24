package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum 
}

func chan1Main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <- c // expecting two receives on channel c, one from each goroutine

	fmt.Println("Sum:", x, y, x + y)

	minute, _ := time.ParseDuration("60ms")
	go announce("That used a channel and two goroutines", minute)
	
}

func announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}()
}