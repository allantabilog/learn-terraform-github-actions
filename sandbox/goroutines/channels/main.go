package main

import (
	"sandbox/goroutines/channels/buffer"
	"sandbox/goroutines/channels/channel"
	"sandbox/goroutines/channels/communicate"
	"sandbox/goroutines/channels/wait"
)

func main() {
	main6()
}

func main6() {
	channel.ContextMain()
}
func main5() {
	wait.MultiplicationMain()
}
func main4() {
	wait.WaitMain()
}

func main3() {
	communicate.CommunicateMain()
}

func main2() {
	buffer.BufferMain()
}
func main1() {
	channel.ChannelMain()
}
