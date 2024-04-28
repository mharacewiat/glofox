package service

import (
	"main/booking"
	"main/class"
	"testing"
)

var service ServiceInterface

func init() {
	service, _ = NewService()
}

func TestNeSetClasss(t *testing.T) {
	class, _ := class.NewClass("Foo", "1970-01-01", "1970-01-01", 1)

	service.CreateClass(class)
}

func TestAddBooking(t *testing.T) {
	booking, _ := booking.NewBooking("Bar", "1970-01-01")

	service.RegisterBooking(booking)
}
