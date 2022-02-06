package struct_sandbox

import "fmt"

type Vertex struct {
	X, Y int 
}

func StructTest() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{X: 100, Y: 200})
	
	vertex := Vertex{0, 0}
	fmt.Println(vertex.X)

	vertexPointer := &vertex
	vertexPointer.X = 1e9
	fmt.Println(vertexPointer.X)
	fmt.Println(vertexPointer)
}
