package main

import "fmt"

type DOCUMENT_STATUS_EVENT string

const DocumentEventCreated DOCUMENT_STATUS_EVENT = "DOCUMENT_CREATED"
const DocumentEventUpdated DOCUMENT_STATUS_EVENT = "DOCUMENT_UPDATED"

type EventDetails struct {
	event string
}

func main() {
	event1 := EventDetails{event: "DOCUMENT_CREATED"}
	event2 := EventDetails{event: "DOCUMENT_UPDATED"}

	fmt.Println(DOCUMENT_STATUS_EVENT(event1.event))
	fmt.Println(DOCUMENT_STATUS_EVENT(event2.event))

}
