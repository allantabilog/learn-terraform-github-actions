package fibonacci

import "fmt"

func Fibonacci(n int, c chan int){
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y  = y, x + y
	}
	close(c)
}

func FibWithSelect(channel, quit chan int){
	x, y := 0, 1

	for {
		select {
		case channel <- x:
			x, y = y, x + y
		case <- quit :
			fmt.Println("quit")
			return
		}
	}
}