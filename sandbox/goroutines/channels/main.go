package main

import (
	"sandbox/goroutines/channels/channel"
)
import "sandbox/goroutines/channels/buffer"

func main() {
	main2()
}
func main2() {
	buffer.BufferMain()
}
func main1() {
	channel.ChannelMain()
}
