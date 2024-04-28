package app

import (
	"fmt"
	"log"
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
