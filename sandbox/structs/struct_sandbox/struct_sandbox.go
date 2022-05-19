package struct_sandbox

import "fmt"

type Vertex struct {
	X, Y int
}

type Workspace struct {
	WorkspaceId  int
	Address      Address
	Participants []Participant
	Documents    []Document
}

type Address struct {
	StreetNumber int
	StreetName   string
}

type Participant struct {
	ParticipantId int
	Reference     string
}

type Document struct {
	ID     int
	TypeId int
	Name   string
	Status string
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

func WorkspaceWithTwoMortgageDocs() *Workspace {
	workspace := Workspace{}
	workspace.WorkspaceId = 1
	workspace.Documents = []Document{
		{ID: 1, TypeId: 22, Name: "National Mortgage", Status: "Created"},
		{ID: 2, TypeId: 22, Name: "National Mortgage", Status: "Created"},
		{ID: 3, TypeId: 4, Name: "Lodgement Instructions", Status: "Created"},
	}

	return &workspace

}

func WorkspaceWithOneMortgageDocs() *Workspace {
	workspace := Workspace{}
	workspace.WorkspaceId = 1
	workspace.Documents = []Document{
		{ID: 1, TypeId: 22, Name: "National Mortgage", Status: "Created"},
		{ID: 2, TypeId: 4, Name: "Lodgement instructions", Status: "Created"},
	}

	return &workspace

}

func UpdateDocumentStatus(workspace *Workspace, documentID int, documentStatus string) {
	for i, doc := range workspace.Documents {
		if doc.ID == documentID {
			fmt.Printf("Updating document %d to status %s\n", doc.ID, documentStatus)
			workspace.Documents[i].Status = documentStatus
		}
	}
}

func UpdateDocumentStatusV2(workspace *Workspace, documentID int, documentStatus string) {
	for i := range workspace.Documents {
		document := &workspace.Documents[i]
		if document.ID == documentID {
			fmt.Printf("Updating document %d to status %s\n", document.ID, documentStatus)
			document.Status = documentStatus
		}
	}
}

func RemoveDocument(workspace *Workspace, documentID int) {
	fmt.Printf("Removing document %d from workspace %d\n", documentID, workspace.WorkspaceId)

	var documents []Document

	for i := range workspace.Documents {
		if workspace.Documents[i].ID != documentID {
			documents = append(documents, workspace.Documents[i])
		}
	}
	// if the last remaining document in the collection is a lodgement instruction, remove it
	if len(documents) == 1 && documents[0].Name == "Lodgement instructions" {
		fmt.Printf("Also removing lodgement instructions from workspace %d\n", workspace.WorkspaceId)
		var emptyDocuments []Document
		workspace.Documents = emptyDocuments
	} else {
		workspace.Documents = documents
	}

}

func AddDocument(workspace *Workspace, id int) {
	document := Document{
		ID:     id,
		TypeId: 1,
		Name:   "Mortgage document",
		Status: "Created",
	}

	workspace.Documents = append(workspace.Documents, document)
}

type B struct {
	c int
}
type A struct {
	b B
}

func TestStructUpdate() {

	a := A{b: B{c: 5}}
	fmt.Println(a)
	a.updateB(42)
	fmt.Println(a)

}

func (a *A) updateB(n int) {
	a.b.c = n
}

func WorkspaceTest4() *Workspace {
	var documents []Document
	documents = append(documents, Document{ID: 1, Name: "Mortgage Document", Status: "Created"})
	workspace := Workspace{
		WorkspaceId: 1, Address: Address{
			StreetNumber: 1, StreetName: "Sesame",
		}, Documents: documents,
	}

	return &workspace
}

func (document Document) String() string {
	return fmt.Sprintf("Document: ID [%v] Type [%s] Status [%s]\n", document.ID, document.Name, document.Status)
}

func (workspace Workspace) String() string {
	return fmt.Sprintf("Workspace: ID [%v] Documents: [%v]\n", workspace.WorkspaceId, workspace.Documents)
}
