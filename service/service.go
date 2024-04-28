package service

import (
	"fmt"
	"main/booking"
	"main/class"
	"main/date"
	"main/storage"
)

type (
	Service struct {
		ClassesStorage  storage.StorageInterface[date.Date, *class.Class]
		BookingsStorage storage.StorageInterface[date.Date, *[]booking.Booking]
	}
	ServiceInterface interface {
		CreateClass(class class.Class) (bool, error)
		RegisterBooking(booking booking.Booking) (bool, error)
		GetClassBookings(day date.Date) (ClassBookings, error)
	}
	ClassBookings struct {
		Name         string    `json:"name"`
		Date         date.Date `json:"date"`
		Capacity     int       `json:"capacity"`
		Participants []string  `json:"participants"`
	}
)

func NewService() (ServiceInterface, error) {
	classesStorage := storage.Storage[date.Date, *class.Class]{}
	bookingsStorage := storage.Storage[date.Date, *[]booking.Booking]{}

	return &Service{
		ClassesStorage:  &classesStorage,
		BookingsStorage: &bookingsStorage,
	}, nil
}

func (s *Service) CreateClass(c class.Class) (bool, error) {
	days := getDays(c.StartDate, c.EndDate)

	for _, day := range days {
		if s.ClassesStorage.Has(day) {
			return false, fmt.Errorf("a different class exists on day %s", day)
		}

		s.ClassesStorage.Set(day, &c)
		s.BookingsStorage.Set(day, &[]booking.Booking{})
	}

	return true, nil
}

func (s *Service) RegisterBooking(b booking.Booking) (bool, error) {
	day := b.Date

	if !s.ClassesStorage.Has(day) {
		return false, fmt.Errorf("there's no class on day %s", day)
	}

	class := s.ClassesStorage.Get(day)

	if class.Capacity == 0 {
		return false, fmt.Errorf("there's no capacity left on day %s", day)
	}

	class.Capacity -= 1

	bookings := s.BookingsStorage.Get(day)
	*bookings = append(*bookings, b)

	return true, nil
}

func (s *Service) GetClassBookings(day date.Date) (ClassBookings, error) {
	if !s.ClassesStorage.Has(day) {
		return ClassBookings{}, fmt.Errorf("there's no class on day %s", day)
	}

	participants := []string{}

	if s.BookingsStorage.Has(day) {
		for _, b := range *s.BookingsStorage.Get(day) {
			participants = append(participants, b.Name)
		}
	}

	c := s.ClassesStorage.Get(day)

	return ClassBookings{
		Name:         c.Name,
		Date:         day,
		Capacity:     c.Capacity,
		Participants: participants,
	}, nil
}

func getDays(from date.Date, to date.Date) []date.Date {
	days := []date.Date{}

	for current := from; current.IsBefore(to.AddDay()); current = current.AddDay() {
		days = append(days, current)
	}

	return days
}
