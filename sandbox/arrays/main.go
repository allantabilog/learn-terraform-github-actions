package main

import (
	"fmt"
)

func main() {
	//sandbox.Main()
	Foo()
}

func Foo() {
	x := 0
	fmt.Println(x)
}

func main0() {
	fmt.Println(0)

	// declare array
	var fruitArr [2]string

	// assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Kiwifruit"

	anotherFruitArr := []string{"Apples", "Oranges", "Grapes"}

	fmt.Println(fruitArr, anotherFruitArr[0:1])

}
