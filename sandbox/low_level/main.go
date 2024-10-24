package main

import (
	"bytes"
	"fmt"
	"os"
)

func main(){
	tickets := []int {1, 2, 3}
	
	fmt.Println(enlarge2(tickets, 10))
}

func enlarge2(s []int, factor int) []int {
	// the educative.io solution 
	newSlice := make([]int, len(s) * factor)
	copy(newSlice, s)

	return newSlice
}

func enlarge(s []int, factor int) []int {
	initLength :=  len(s)

	newLength := initLength * factor 

	for len(s) < newLength {
		s = append(s, 0)
	}

	fmt.Println("Final slice: ", s, " length: ", len(s))

	return s

}


func reslice(){
	s := []int {2,3,5,7,11}
	fmt.Println(s)
	fmt.Println(cap(s))
	s = append(s, 13)
	fmt.Println(s)
	fmt.Println(cap(s))

	for i:=0;i<20;i++ {
		s = append(s, i+10)
		fmt.Println("slice: ", s)
		fmt.Println("capacity:", cap(s))
	}
	

}

func main2()  {
	fmt.Println("heya!")
	var b bytes.Buffer

	n, _ :=	b.WriteString("ABC")
	fmt.Println(n)
	n, _ = b.WriteString("DEFGHIJ")
	fmt.Println(n)

	fmt.Fprintf(&b, "")
	b.WriteString("[DONE]")

	fmt.Println(b.String())


}

func main3(){
	var b bytes.Buffer 
	b.Write([]byte ("Hello"))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
}
