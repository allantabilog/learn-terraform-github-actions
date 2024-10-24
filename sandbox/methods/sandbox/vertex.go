package sandbox

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

type Abser interface {
	Abs() float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func VertexMain() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(100)
	fmt.Println(v.Abs())

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v = Vertex{4, 5}

	a = &f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())

}
