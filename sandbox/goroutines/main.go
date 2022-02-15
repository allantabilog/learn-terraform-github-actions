package main

import (
	"fmt"
	"math/rand"
	"sandbox/goroutines/fibonacci"
	"time"
)

func say(s string){
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func go_say(){
	go say("world")
	say("hello")
}

func sum(nums []int, channel chan int){
	sum := 0
	for _, num := range nums {
		sum += num
	}
	channel <- sum
}

func main_0() {
	ints := []int {7, 2, 8, -9, 4, 0}
	
	channel := make(chan int)

	go sum(ints[:len(ints)/2], channel)
	go sum(ints[len(ints)/2:], channel)

	sum1, sum2 := <- channel, <- channel
	fmt.Println(sum1, sum2, sum1 + sum2)
}

func main_1() {
	channel := make(chan int, 10)
	go fibonacci.Fibonacci(cap(channel), channel)
	for i:= range channel {
		fmt.Println(i)
	}
}

func main_4() {
	channel := make(chan int)
	quit := make(chan int)
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()
	fibonacci.FibWithSelect(channel, quit)


}

func display(text string){
	for i:=0; i<6;i++ {
		time.Sleep(10 * time.Second)
		fmt.Println(text)
	}
}
func main_3() {
	fmt.Println("Running...")
	go display("Welcome")
	//go display("Geeks for geeks")
}

func main_5(){
	messages := make(chan string, 3)

	messages <- "one"
	messages <- "two"
	messages <- "three"

	go func(messages *chan string){
		fmt.Println("Entering the goroutine")
		for {
			fmt.Println(<- *messages)
		}
	}(&messages)
	time.Sleep(10 * time.Second)
	fmt.Println("Done.")
}

func printHello(ch chan int){
	fmt.Println("Hello")
	ch <- 2
}

func main(){
	ch := make(chan int)

	go func() {
		fmt.Println("From main")
		ch <- 1
	}()

	go printHello(ch)
	i := <- ch
	fmt.Println("Received from channel: ", i)
}
func main_6() {
	go func() {
		fmt.Println(rand.Intn(1001))
		fmt.Println(rand.Intn(1001))
		fmt.Println(rand.Intn(1001))
		fmt.Println(rand.Intn(1001))

	}()
	time.Sleep(10 * time.Second)
}