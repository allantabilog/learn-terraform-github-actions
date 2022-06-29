package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"math"
	"sandbox/mockery/sandbox"
	"sandbox/mockery/sandbox/mocks"
	"testing"
)

func TestSandbox2(t *testing.T) {
	type AbserClient struct {
		abser sandbox.Abser
	}

	var mockAbser mocks.Abser
	mockAbser.On("Abs").Return(float64(-1)).Run(func(args mock.Arguments) {
		fmt.Println("hello world from mock")
	})
	abserClient := AbserClient{abser: &mockAbser}

	//abserClient.abser.Abs()
	fmt.Println(abserClient.abser.Abs())

}

func TestSandbox(t *testing.T) {
	var a sandbox.Abser

	f := sandbox.MyFloat(-math.Sqrt2)
	a = f

	fmt.Println(a.Abs())

	v := sandbox.Vertex{X: 10, Y: -20}
	a = &v
	fmt.Println(a.Abs())

}

func TestString(t *testing.T) {
	mock := mocks.Stringer{}

	mock.On("String").Return("hello from the mock")

}
