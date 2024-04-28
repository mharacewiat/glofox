package booking

import "testing"

func TestNewBooking(t *testing.T) {
	name, date := "Bar", "1970-01-01"

	newBooking, err := NewBooking(name, date)
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if newBooking.Name != name {
		t.Errorf("Expected booking name to be %s, got %s", name, newBooking.Name)
	}

	if newBooking.Date != date {
		t.Errorf("Expected booking date to be %s, got %s", date, newBooking.Date)
	}
}
