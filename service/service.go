package service

import "fmt"

// MessageService - interface for sending notifications to the user
type MessageService interface {
	SendChargeNotification(int) bool
}

// SMS Service - sends SMS notifications to the user
type SMSService struct{}

// My Service - my service
type MyService struct {
	MessageService MessageService
}

// SendChargeNotification - sends a notification to the user
func (sms *SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending SMS notification for charge of " + fmt.Sprint(value))
	return true
}

// ChargeCustomer - charges the customer for our service
func (ms *MyService) CargeCustomer(value int) error {
	ms.MessageService.SendChargeNotification(value)
	fmt.Printf("Charge of %d was successful\n", value)
	return nil
}
