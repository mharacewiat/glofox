package date

import "testing"

func TestNewDate(t *testing.T) {
	_, err := NewDate("1970-01-01")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}
}

func TestNewDateInvalid(t *testing.T) {
	_, err := NewDate("foo")
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}
}

func TestIsBefore(t *testing.T) {
	d1 := Date("1970-01-01")
	d2 := Date("1970-01-02")

	if !d1.IsBefore(d2) {
		t.Errorf("%v should be before %v", d1, d2)
	}

	if d2.IsBefore(d1) {
		t.Errorf("%v should not be before %v", d2, d1)
	}
}
