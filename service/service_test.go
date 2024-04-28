package service

import (
	"main/booking"
	"main/class"
	"main/date"
	"testing"
)

var service ServiceInterface

func init() {
	service, _ = NewService()
}

func TestCreateClass(t *testing.T) {
	class, _ := class.NewClass("Foo", "1970-01-01", "1970-01-02", 1)

	ok, err := service.CreateClass(class)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !ok {
		t.Error("Expected success, got failure")
	}
}

func TestRegisterBooking(t *testing.T) {
	booking, _ := booking.NewBooking("Bar", "1970-01-01")

	ok, err := service.RegisterBooking(booking)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !ok {
		t.Error("Expected success, got failure")
	}
}

func TestGetClassBooking(t *testing.T) {
	day := date.Date("1970-01-01")
	classBookings, err := service.GetClassBookings(day)

	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if classBookings.Name != "Foo" {
		t.Errorf("Expected class name to be Foo, got %s", classBookings.Name)
	}

	if classBookings.Capacity != 0 {
		t.Errorf("Expected class capacity to be 0, got %d", classBookings.Capacity)
	}

	if classBookings.Date != day {
		t.Errorf("Expected class capacity to be %s, got %s", day, classBookings.Date)
	}

	if len(classBookings.Participants) != 1 {
		t.Errorf("Expected class capacity to be 1, got %d", len(classBookings.Participants))
	}

	if classBookings.Participants[0] != "Bar" {
		t.Errorf("Expected class capacity to be Bar, got %s", classBookings.Participants[0])
	}
}
