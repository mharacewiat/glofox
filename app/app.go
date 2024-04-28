package app

import (
	"fmt"
	"log"
	"net/http"
)

type (
	App struct {
		Port string
	}
	AppInterface interface {
		Start()
		HandleStatus(w http.ResponseWriter, r *http.Request)
		HandlePutClasses(w http.ResponseWriter, r *http.Request)
		HandlePostBookings(w http.ResponseWriter, r *http.Request)
	}
)

func NewApp(port string) (AppInterface, error) {
	return &App{
		Port: port,
	}, nil
}

func (a *App) Start() {
	http.HandleFunc("GET /status", a.HandleStatus)
	http.HandleFunc("PUT /classes", a.HandlePutClasses)
	http.HandleFunc("POST /bookings", a.HandlePostBookings)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", a.Port), nil))
}

func (a *App) HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *App) HandlePutClasses(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *App) HandlePostBookings(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
