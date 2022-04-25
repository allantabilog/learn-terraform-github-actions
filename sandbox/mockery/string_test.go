package main

import (
	"fmt"
	"sandbox/mockery/sandbox/mocks"
	"testing"
)

func TestString(t *testing.T) {
	f := mocks.Stringer{}

	f.On("String").Return("hello from the mock")

	f.On("Add", 1, 2).Return(3)
	fmt.Println(f.String())
	fmt.Println(f.Add(1, 2))
	fmt.Println(f.Add(1, 21))
}
