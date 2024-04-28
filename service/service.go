package service

import (
	"main/booking"
	"main/class"
)

type (
	Service          struct{}
	ServiceInterface interface {
		CreateClass(class class.Class) (bool, error)
		RegisterBooking(booking booking.Booking) (bool, error)
	}
)

func NewService() (ServiceInterface, error) {
	return &Service{}, nil
}

func (s *Service) CreateClass(class class.Class) (bool, error) {
	return true, nil
}

func (s *Service) RegisterBooking(booking booking.Booking) (bool, error) {
	return true, nil
}
