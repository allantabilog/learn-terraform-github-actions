package main

import (
	"fmt"
	"sandbox/structs/struct_sandbox"
	"strconv"
)

// Define a person struct
// A struct is just a collection of fields
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// methods
// a value receiver:
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " age: " + strconv.Itoa(p.age)
}

// a pointer receiver
func (p *Person) incrementAge() {
	p.age++
}

func (p *Person) updateLastName(newLastName string) {
	p.lastName = newLastName
}

func main() {
	//struct_sandbox.StructTest()
	//struct_sandbox.WorkspaceTest()

	//fmt.Println(struct_sandbox.WorkspaceTest2())
	//struct_sandbox.TestStructUpdate()
	//testUpdateDocumentInWorkspace()
	testRemoveDocumentFromWorkspace()
	testAddDocumentToWorkspace()
}

func testUpdateDocumentInWorkspace() {
	workspace := struct_sandbox.WorkspaceTest3()
	fmt.Println(workspace)
	fmt.Println("workspace documents: ")

	for _, doc := range workspace.Documents {
		fmt.Printf("[id: %d name: %s status: %s]\n", doc.ID, doc.Name, doc.Status)
	}
	fmt.Println("Updating the status of document 2")
	//struct_sandbox.UpdateDocumentStatus(workspace, 2, "Signed")
	struct_sandbox.UpdateDocumentStatusV2(workspace, 2, "Signed")
	fmt.Println(workspace)
}

func testRemoveDocumentFromWorkspace() {
	workspace := struct_sandbox.WorkspaceTest3()
	fmt.Println(workspace)

	fmt.Println("Before removing Document: ", workspace)
	struct_sandbox.RemoveDocument(workspace, 1)
	fmt.Println("After removing document: ", workspace)
}

func testAddDocumentToWorkspace() {
	workspace := struct_sandbox.WorkspaceTest3()

	fmt.Println("Before adding a Document: ", workspace)
	struct_sandbox.AddDocument(workspace, 3)
	fmt.Println("After adding document: ", workspace)

}

func test_1() {
	// Init person struct
	person1 := Person{firstName: "Bob", lastName: "The.Builder", city: "Melbz", age: 30, gender: "M"}
	fmt.Println(person1)

	person2 := Person{"Bobb", "The.Bbuilder", "Melbourne", "M", 30}
	fmt.Println(person2)

	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	person1.incrementAge()
	person1.incrementAge()
	person1.incrementAge()
	person1.incrementAge()
	fmt.Println(person1.greet())

	person1.updateLastName("New.Builder")
	fmt.Println(person1.greet())

}
