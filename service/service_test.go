package service_test

import (
	"fmt"
	"testing"

	"github.com/Dmitrylolo/golang-tests/service"
	"github.com/stretchr/testify/mock"
)

type smsServiceMock struct {
	mock.Mock
}

func (m *smsServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charge notification")
	args := m.Called(value)
	return args.Bool(0)
}

func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)
	smsService.On("SendChargeNotification", 100).Return(true)

	myService := service.MyService{smsService}

	myService.CargeCustomer(100)

	smsService.AssertExpectations(t)
}
