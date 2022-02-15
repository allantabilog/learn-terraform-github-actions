package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type smsServiceMock struct {
	mock.Mock
}

func (m *smsServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charged notification function")
	fmt.Printf("valued passed in: %d\n", value)

	args := m.Called(value)

	return args.Bool(0)
}
