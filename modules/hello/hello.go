package main

import (
	"allantabilog/greetings"
	"fmt"
)

func main() {
	msg := greetings.Hello("Joe")
	fmt.Println(msg)
}
