package date

import (
	"errors"
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

	return err == nil
}

func (d Date) IsBefore(other Date) bool {
	parsedDate, _ := d.Parse()
	parsedOther, _ := other.Parse()

	return parsedDate.Before(parsedOther)
}
