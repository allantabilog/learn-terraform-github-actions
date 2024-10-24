package main

import (
	"bytes"
	"testing"
)

func TestDisplayGreetings(t *testing.T) {
	buf := new(bytes.Buffer)
	displayGreetings(buf)

	expected := "Hello\nWorld\n"
	actual := buf.String()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}