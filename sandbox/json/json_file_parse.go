package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Customers struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	RoleID int    `json:"role_id"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func (cust Customer) String() string {
	return fmt.Sprintf("Name: %s, type: %s, age: %d, Facebook: %s Twitter %s", cust.Name, cust.Type, cust.Age, cust.Social.Facebook, cust.Social.Twitter)

}

type DocumentStatusEventDetails struct {
	BusinessEventID     int       `json:"business_event_id"`
	CreatedDate         time.Time `json:"created_date"`
	UserID              int       `json:"user_id"`
	LodgementCaseID     string    `json:"lodgement_case_id"`
	SubscriberID        int       `json:"subscriber_id"`
	WorkspaceID         int       `json:"workspace_id"`
	WorkspaceRole       string    `json:"workspace_role"`
	Username            string    `json:"username"`
	BusinessEventTypeID int       `json:"business_event_type_id"`
	DocumentID          int       `json:"document_id"`
	DocumentTypeID      int       `json:"document_type_id"`
	DocumentType        string    `json:"document_type"`
	DocumentStatus      string    `json:"status"`
	CreatedBySubscriber int       `json:"created_by_subscriber_id"`
	CreatedByRole       int       `json:"created_by_role_id"`
}

func (event DocumentStatusEventDetails) String() string {
	return fmt.Sprintf("event id: [%d] subscriber: [%d] document type: [%s] document status: [%s] created by subscriber: [%d] created by role: [%d]",
		event.BusinessEventID, event.SubscriberID, event.DocumentType, event.DocumentStatus, event.CreatedBySubscriber, event.CreatedByRole,
	)
}

func main() {
	//testOpen()
	//testReadAll()
	testUnmarshall2()
}
func testUnmarshall() {
	jsonFile, err := os.Open("json/customers.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("successfully opened file")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println("JSON content: ", string(byteValue))

	var customers Customers
	err = json.Unmarshal(byteValue, &customers)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
	}

	fmt.Println("--- Customer Data unmarshalled ---")
	for i := 0; i < len(customers.Customers); i++ {
		fmt.Println(customers.Customers[i])
	}

}
func testUnmarshall2() {
	jsonFile, err := os.Open("json/document_created_event.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("successfully opened file")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Println("JSON content: ", string(byteValue))

	var statusEvent DocumentStatusEventDetails
	json.Unmarshal(byteValue, &statusEvent)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
	}

	fmt.Println("--- Customer Data unmarshalled ---")
	fmt.Println(statusEvent)

}

func testOpen() {
	fmt.Println("testing")
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Current working directory:", currentDir)
	jsonFile, err := os.Open("json/testFile.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("successfully opened file")
	contents, err := ioutil.ReadFile("json/testFile.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file contents: %s", contents)
	defer jsonFile.Close()
}

func testReadAll() {
	reader := strings.NewReader("Go is a general-purpose language designed for sytems programming")

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", bytes)
}
