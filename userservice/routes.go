package userservice

import "github.com/dillonerhardt/rest-api-golang/app"

// routes defines the routes for this service
var routes = []app.Route{
	{
		Name:        "GetAuthorizedUserHandler",
		Method:      "GET",
		Pattern:     "/me",
		HandlerFunc: getAuthorizedUserHandler,
	},
	{
		Name:        "GetUserHandler",
		Method:      "GET",
		Pattern:     "/{id}",
		HandlerFunc: getUserHandler,
	},
	{
		Name:        "UpdateUserHandler",
		Method:      "PUT",
		Pattern:     "/{id}",
		HandlerFunc: updateUserHandler,
	},
	{
		Name:        "DeleteUserHandler",
		Method:      "DELETE",
		Pattern:     "/{id}",
		HandlerFunc: deleteUserHandler,
	},
}
