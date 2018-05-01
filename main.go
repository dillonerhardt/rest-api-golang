package main

import (
	"github.com/dillonerhardt/rest-api-golang/app"
	"github.com/dillonerhardt/rest-api-golang/authservice"
	"github.com/dillonerhardt/rest-api-golang/db"
	"github.com/dillonerhardt/rest-api-golang/userservice"
	"github.com/dillonerhardt/rest-api-golang/utils"
)

// server entry point
func main() {
	// load the app config
	utils.LoadConfig("./config.json")
	// create MongoDB connection
	db.ConnectMongoDB(utils.Config.Database.URL)
	// create app
	a := app.CreateApp()
	// create the services and add them
	// to the app
	a.AddServices(
		userservice.CreateService("/users"),
		authservice.CreateService("/auth"),
	)
	// run the app
	a.Run(":" + utils.Config.Port)
}
