package geometry

import "math"

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base   float64
	height float64
}

// make rect and circle implement the geometry interface by implementing all of the methods in the interface
func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.length)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	// perimeter == circumference
	return 2 * math.Pi * c.radius
}

// triangle won't implement all the methods of geometry interface
// lets see what happens if we try and treat it as a geometry

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}
