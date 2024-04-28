package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/booking"
	"main/class"
	"main/date"
	"main/service"
	"net/http"
)

type (
	App struct {
		Port    string
		Service service.ServiceInterface
	}
	AppInterface interface {
		Start()
		HandleStatus(w http.ResponseWriter, r *http.Request)
		HandlePutClasses(w http.ResponseWriter, r *http.Request)
		HandlePostBookings(w http.ResponseWriter, r *http.Request)
	}
)

func NewApp(port string) (AppInterface, error) {
	service, err := service.NewService()
	if err != nil {
		return &App{}, err
	}

	return &App{
		Port:    port,
		Service: service,
	}, nil
}

func (a *App) Start() {
	http.HandleFunc("GET /status", a.HandleStatus)
	http.Handle("PUT /classes", checkJsonContentType(http.HandlerFunc(a.HandlePutClasses)))
	http.Handle("POST /bookings", checkJsonContentType(http.HandlerFunc(a.HandlePostBookings)))
	http.HandleFunc("GET /day/{day}", a.HandleDay)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", a.Port), nil))
}

func (a *App) HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *App) HandlePutClasses(w http.ResponseWriter, r *http.Request) {
	var c class.Class

	err := decode(r.Body, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if !c.IsValid() {
		http.Error(w, "input data invalid", http.StatusBadRequest)

		return
	}

	_, err = a.Service.CreateClass(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) HandlePostBookings(w http.ResponseWriter, r *http.Request) {
	var b booking.Booking

	err := decode(r.Body, &b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if !b.IsValid() {
		http.Error(w, "input data invalid", http.StatusBadRequest)

		return
	}

	_, err = a.Service.RegisterBooking(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) HandleDay(w http.ResponseWriter, r *http.Request) {
	stringDay := r.PathValue("day")

	day, err := date.NewDate(stringDay)
	if err != nil {
		http.Error(w, "input data invalid", http.StatusBadRequest)

		return
	}

	classBookings, err := a.Service.GetClassBookings(day)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	response, err := json.Marshal(classBookings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func checkJsonContentType(next http.Handler) http.Handler {
	return checkContentType("application/json", next)
}

func checkContentType(contentType string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != contentType {
			http.Error(w, "invalid content-type", http.StatusBadRequest)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func decode(body io.Reader, output interface{}) error {
	return json.NewDecoder(body).Decode(&output)
}
