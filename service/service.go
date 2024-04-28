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

func getDays(from date.Date, to date.Date) []date.Date {
	days := []date.Date{}

	for current := from; current.IsBefore(to); current = current.AddDay() {
		days = append(days, current)
	}

	return days
}
