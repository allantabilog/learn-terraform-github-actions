package main

import (
	"fmt"
)

func main() {
	//sandbox.Main()
	Foo()
}

func Foo() {
	main1()
}

func main1() {
	var x = []int{1, 2, 3}
	x = append(x, 2)
	fmt.Println(x)
	y := x[1:2]
	fmt.Println(y)
	y[0] = 100
	fmt.Println(y)
	fmt.Println(x)
	x = append(x, 200, 300, 400, 500)
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
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
