package greetings

import "fmt"

func Hello(name string) string {
	msg := fmt.Sprintf("Hi there %v! Welcome!", name)
	return msg
}
