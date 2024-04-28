package class

import (
	"main/date"
	"testing"
)

func TestNewClass(t *testing.T) {
	name, startDate, endDate, capacity := "Foo", date.Date("1970-01-01"), date.Date("1970-01-02"), 1

	newClass, err := NewClass(name, startDate, endDate, capacity)
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if newClass.Name != name {
		t.Errorf("Expected class name to be %s, got %s", name, newClass.Name)
	}

	if newClass.StartDate != startDate {
		t.Errorf("Expected class start date to be %s, got %s", startDate, newClass.StartDate)
	}

	if newClass.EndDate != endDate {
		t.Errorf("Expected class end date to be %s, got %s", endDate, newClass.EndDate)
	}

	if newClass.Capacity != capacity {
		t.Errorf("Expected class capacity to be %d, got %d", capacity, newClass.Capacity)
	}
}

func TestNewClassInvalidName(t *testing.T) {
	_, err := NewClass("", date.Date("1970-01-01"), date.Date("1970-01-02"), 1)
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}
}

func TestNewClassInvalidDates(t *testing.T) {
	_, err := NewClass("Foo", date.Date("1970-01-02"), date.Date("1970-01-01"), 1)
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}
}

func TestNewClassInvalidCapacity(t *testing.T) {
	_, err := NewClass("Foo", date.Date("1970-01-01"), date.Date("1970-01-02"), 0)
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}
}
