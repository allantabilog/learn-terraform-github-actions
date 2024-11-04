package main

import (
	"fmt"

	"github.com/allantabilog/increment/impl1"
	"github.com/allantabilog/increment/impl2"
	"github.com/allantabilog/increment/incrementer"
)

func main() {
	var inc1 incrementer.Incrementer = impl1.NewIncrementerImpl1()
	var inc2 incrementer.Incrementer = impl2.NewIncrementerImpl2()

	fmt.Println(inc1.Increment())
	fmt.Println(inc2.Increment())
}