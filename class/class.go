package class

import "main/date"

type Class struct {
	Name      string    `json:"name"`
	StartDate date.Date `json:"start_date"`
	EndDate   date.Date `json:"end_date"`
	Capacity  int       `json:"capacity"`
}

func NewClass(name string, startDate date.Date, endDate date.Date, capacity int) (Class, error) {
	return Class{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  capacity,
	}, nil
}
