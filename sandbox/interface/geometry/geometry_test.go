package geometry

import (
	"fmt"
	"testing"
)

func TestGeometry(t *testing.T) {
	var r rect
	var c circle

	r = rect{
		length: 10,
		width:  20,
	}

	c = circle{radius: 1}

	measure(r)
	measure(c)

	var tr triangle
	tr = triangle{
		base:   10,
		height: 57,
	}

	//measure(tr) // this is a compile error. The compiler knows that t does not implement geometry

	fmt.Println(tr.area())

}

func measure(g geometry) {
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perimeter())
}
