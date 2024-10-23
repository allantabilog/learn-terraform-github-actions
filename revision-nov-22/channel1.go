package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	rand.Seed(time.Now().UnixNano())

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)
		}
	}

}
