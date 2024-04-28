package date

import (
	"errors"
	"log"
	"time"
)

type Date string

func NewDate(dateString string) (Date, error) {
	date := Date(dateString)

	if !date.IsValid() {
		return "", errors.New("invalid date")
	}

	return date, nil
}

func (d Date) Parse() (time.Time, error) {
	return time.Parse("2006-01-02", string(d))
}

func (d Date) IsValid() bool {
	_, err := d.Parse()

	if err != nil {
		log.Print(err.Error())

		return false
	}

	return true
}

func (d Date) IsBefore(other Date) bool {
	parsedDate, _ := d.Parse()
	parsedOther, _ := other.Parse()

	return parsedDate.Before(parsedOther)
}

func (d Date) AddDay() Date {
	parsedDate, _ := d.Parse()
	parsedDate = parsedDate.AddDate(0, 0, 1)

	return Date(parsedDate.Format("2006-01-02"))
}
