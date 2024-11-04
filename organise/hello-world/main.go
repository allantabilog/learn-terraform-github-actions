package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/allantabilog/hello-world/hello"
	"github.com/allantabilog/hello-world/world"
)

func displayGreetings(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, world.Greet())
}

func addBy(n int) int {
	var sum int
	defer func() {
		// sum = -1
		fmt.Println("Exiting the function with sum: ", sum)
	}()
	sum = 100 + n
	fmt.Println("addBy exiting with sum: ", sum)
	return sum
}
func deferExample() {
	defer fmt.Println("Exiting the function")
	fmt.Println("This will be printed first")
}
func main() {
	// get input from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number: ")
	text, _ := reader.ReadString('\n')
	// strip the newline character
	text = strings.TrimSpace(text)

	fmt.Printf("You entered: [%v]", text)
	if(text == "0"){
			panic(
				"0 is not allowed",
			)
		}
	addBy((10))
}