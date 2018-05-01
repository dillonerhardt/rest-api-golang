package app

import (
	"github.com/dillonerhardt/rest-api-golang/middleware"
	"github.com/gorilla/mux"
)

// Service defines stucture for a service
type Service struct {
	Path   string  // Url prefix for service
	Routes []Route // Slice of routes for service
}

// AddSubrouter creates a subrouter with the services routes
func (s *Service) AddSubrouter(r *mux.Router) {
	sr := r.PathPrefix(s.Path).Subrouter()
	for _, r := range s.Routes {
		// set middle ware to check jwt auth
		handler := middleware.ValidateJWT(r.Public, r.HandlerFunc)
		// build the routes
		sr.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(middleware.Logger(handler, r.Name)) // apply logger middleware

	}
}
