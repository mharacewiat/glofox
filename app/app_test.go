package app

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app AppInterface

func init() {
	app, _ = NewApp("1234")
}

func TestGetStatus(t *testing.T) {
	request, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	response := httptest.NewRecorder()

	handler := http.HandlerFunc(app.HandleGetStatus)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler for GET /status returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPutClasses(t *testing.T) {
	body := []byte(`{"name": "Foo", "start_date": "1970-01-01", "end_date": "1970-01-02", "capacity": 1}`)
	request, err := http.NewRequest("PUT", "/classes", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler := http.HandlerFunc(app.HandlePutClasses)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler for PUT /classes returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPostBookings(t *testing.T) {
	body := []byte(`{"name": "Bar", "date": "1970-01-01"}`)
	request, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler := http.HandlerFunc(app.HandlePostBookings)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler for POST /bookings returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetDay(t *testing.T) {
	request, err := http.NewRequest("GET", "/day/1970-01-01", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	request.SetPathValue("day", "1970-01-01")
	response := httptest.NewRecorder()

	handler := http.HandlerFunc(app.HandleGetDay)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler for POST /day/1970-01-01 returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
