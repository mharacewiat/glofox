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

	handler := http.HandlerFunc(app.HandleStatus)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler for GET /status returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPutClasses(t *testing.T) {
	body := []byte(`{}`)
	request, err := http.NewRequest("PUT", "/classes", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler := http.HandlerFunc(app.HandlePutClasses)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler for PUT /classses returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPostBookings(t *testing.T) {
	body := []byte(`{}`)
	request, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler := http.HandlerFunc(app.HandlePutClasses)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Handler for POST /bookings returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
