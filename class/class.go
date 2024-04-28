package class

import (
	"errors"
	"main/date"
)

type Class struct {
	Name      string    `json:"name"`
	StartDate date.Date `json:"start_date"`
	EndDate   date.Date `json:"end_date"`
	Capacity  int       `json:"capacity"`
}

func NewClass(name string, startDate date.Date, endDate date.Date, capacity int) (Class, error) {
	class := Class{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  capacity,
	}

	if !class.IsValid() {
		return Class{}, errors.New("invalid class")
	}

	return class, nil
}

func (c Class) IsValid() bool {
	return len(c.Name) > 0 &&
		c.StartDate.IsValid() &&
		c.EndDate.IsValid() &&
		c.StartDate.IsBefore(c.EndDate) &&
		c.Capacity > 0
}
