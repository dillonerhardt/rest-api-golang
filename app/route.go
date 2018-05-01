package app

import "net/http"

// Route struct defines a single route. A human readable name, HTTP method, the
// pattern and the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Public      bool
}
