package booking

import (
	"errors"
	"main/date"
)

type Booking struct {
	Name string    `json:"name"`
	Date date.Date `json:"date"`
}

func NewBooking(name string, date date.Date) (Booking, error) {
	booking := Booking{
		Name: name,
		Date: date,
	}

	if !booking.IsValid() {
		return Booking{}, errors.New("invalid class")
	}

	return booking, nil
}

func (b Booking) IsValid() bool {
	return len(b.Name) > 0 && b.Date.IsValid()
}
