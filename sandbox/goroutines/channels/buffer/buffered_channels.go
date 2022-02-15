package buffer

import (
	"fmt"
	"time"
)

func BufferMain() {
	c := make(chan int, 3)
	q := 10
	go worker(c, 1)
	go worker(c, 2)
	go worker(c, 3)
	go generateIntegers(c, q)

	time.Sleep(10 * time.Second)
}

func worker(c chan int, id int) {
	for i := range c {
		fmt.Printf("Worker ID[%d] received value [%d] from the channel\n", id, i)
		time.Sleep(1 * time.Second)
	}
}

func generateIntegers(c chan int, q int) {
	for i := 0; i < q; i++ {
		c <- i
		fmt.Printf("generateIntegers() goroutine sent value [%d] to the channel\n", i)
		time.Sleep(1 * time.Millisecond)
	}

}
