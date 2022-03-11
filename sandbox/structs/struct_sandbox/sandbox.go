package struct_sandbox

import "fmt"

type Vertex struct {
	X, Y int
}

type Workspace struct {
	WorkspaceId  int
	Address      Address
	Participants []Participant
}

type Address struct {
	StreetNumber int
	StreetName   string
}

type Participant struct {
	ParticipantId int
	Reference     string
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
func WorkspaceTest() {
	workspace := Workspace{
		WorkspaceId: 1,
		Address: Address{
			StreetNumber: 1,
			StreetName:   "Dorcas",
		},
	}
	fmt.Println(workspace)
}

func WorkspaceTest2() *Workspace {
	workspace := Workspace{
		WorkspaceId: 1,
		Address: Address{
			StreetNumber: 1,
			StreetName:   "Dorcas",
		},
		Participants: []Participant{
			{
				ParticipantId: 1,
				Reference:     "First Participant",
			},
			{
				ParticipantId: 2,
				Reference:     "2nd Participant",
			},
		},
	}

	return &workspace
}
