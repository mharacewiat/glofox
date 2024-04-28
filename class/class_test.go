package class

import "testing"

func TestNewClass(t *testing.T) {
	name, startDate, endDate, capacity := "Foo", "1970-01-01", "1970-01-02", 1

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
