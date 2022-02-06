package main

import "fmt"

func main(){
	i := 42 
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 22
	fmt.Println(*p)
}
func f1() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*b)

	// update value through pointer
	*b = 10
	fmt.Println(a)
}
