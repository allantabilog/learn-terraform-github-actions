package sandbox

import "math"

type Abser interface {
	Abs() float64
}

// define some implementations of Abser
type MyFloat float64

// MyFloat implements Abser
func (myFloat MyFloat) Abs() float64 {
	if myFloat < 0 {
		return float64(-myFloat)
	}
	return float64(myFloat)
}

type Vertex struct {
	X, Y float64
}

// Vertex implements Abser
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}