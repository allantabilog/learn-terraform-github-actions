package main

import (
	"fmt"
	"io"
	"os"

	"github.com/amitsaha/using-go-modules/greetings/hello"
	"github.com/amitsaha/using-go-modules/greetings/world"
)

func displayGreetings(w io.Writer){
	fmt.Println(w, hello.Greet())
	fmt.Println(w, world.Greet())
}

func main(){
	displayGreetings(os.Stdout)
}