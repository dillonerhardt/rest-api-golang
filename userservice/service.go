package userservice

import "github.com/dillonerhardt/rest-api-golang/app"

// CreateService initialises the service
func CreateService(path string) app.Service {
	return app.Service{
		Path:   path,
		Routes: routes,
	}
}
