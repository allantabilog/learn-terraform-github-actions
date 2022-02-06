package main

import (
	"fmt"
	"sandbox/arrays/sandbox"
)

func main(){
	sandbox.Main()
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
