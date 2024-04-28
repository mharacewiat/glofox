package booking

type Booking struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func NewBooking(name string, date string) (Booking, error) {
	return Booking{
		Name: name,
		Date: date,
	}, nil
}
