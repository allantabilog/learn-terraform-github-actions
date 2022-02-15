package channel

import "fmt"

func ChannelMain() {
	channel := make(chan int) // essentially a "rendezvous" channel
	limit := 10
	go generateInts(channel, limit)

	for i := range channel {
		fmt.Println(i)
	}

}

func generateInts(channel chan int, limit int) {
	for i := 0; i < limit; i++ {
		channel <- i
	}
	close(channel)
}
