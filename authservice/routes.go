package authservice

import "github.com/dillonerhardt/rest-api-golang/app"

// routes defines the routes for this service
var routes = []app.Route{
	{
		Name:        "LoginHandler",
		Method:      "POST",
		Pattern:     "/login",
		HandlerFunc: loginHandler,
		Public:      true,
	},
	{
		Name:        "SignupHandler",
		Method:      "POST",
		Pattern:     "/signup",
		HandlerFunc: signupHandler,
		Public:      true,
	},
}
