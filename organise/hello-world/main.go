package main

import (
	"fmt"
	"io"
	"os"

	"github.com/allantabilog/hello-world/hello"
	"github.com/allantabilog/hello-world/world"
)

func displayGreetings(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, world.Greet())
}
func main() {
	displayGreetings(os.Stdout)
}