package messaging

import "fmt"

type MessageService interface {
	SendChargeNotification(int) error
}

// this will implement MessageService
type SMSService struct {
}

// this uses a MessageService to notify clients
type MyService struct {
	messageService MessageService
}

func (sms SMSService) SendChargeNotification(value int) error {
	fmt.Println("Sending production charge notification")
	return nil
}

func (myService MyService) ChargeCustomer(value int) error {
	myService.messageService.SendChargeNotification(value)
	fmt.Printf("Charging customer for the value of %d\n", value)
	return nil
}

func MessagingMain() {
	fmt.Println("Running the service")
	smsService := SMSService{}
	myService := MyService{smsService}
	myService.ChargeCustomer(100)
}
