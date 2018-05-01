package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App defines structure for the application
type App struct {
	Router *mux.Router
}

// CreateApp creates an instance off the app
func CreateApp() App {
	// create an instance of Gorilla router
	router := mux.NewRouter()
	return App{
		Router: router,
	}
}

// AddServices creates the subrouters for the passed services
func (a *App) AddServices(services ...Service) {
	for _, s := range services {
		s.AddSubrouter(a.Router)
	}
}

// Run will run the application
func (a *App) Run(addr string) {
	http.Handle("/", a.Router)
	log.Println("Starting HTTP service at " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
