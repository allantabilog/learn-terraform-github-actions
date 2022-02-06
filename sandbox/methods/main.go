package main

import (
	"fmt"
	"math"
	"sandbox/methods/sandbox"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64){
	v.X = v.X * f 
	v.Y = v.Y * f
}

func main0() {

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())

	v2 := Vertex{1, 2}
	Scale(&v2, 100)
	fmt.Println(v2.Abs())





}

func main(){
	sandbox.Main()
	sandbox.StringerMain()
}