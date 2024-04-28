package booking

import "main/date"

type Booking struct {
	Name string    `json:"name"`
	Date date.Date `json:"date"`
}

func NewBooking(name string, date date.Date) (Booking, error) {
	return Booking{
		Name: name,
		Date: date,
	}, nil
}
