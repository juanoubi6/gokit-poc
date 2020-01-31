package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gokit-poc/authentications"
	"gokit-poc/commons"
	"gokit-poc/users"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	//Create database
	var db *gorm.DB = commons.CreateDatabase(commons.DatabaseURI)

	//Create user service, it's endpoints and add these to the router.
	userService := users.UserServiceFactory(db)
	userServiceEndpoints := users.MakeEndpoints(userService)
	users.NewHTTPHandler(router, userServiceEndpoints)

	//Create authentication service, it's endpoints and add these to the router.
	authenticationService := authentications.AuthenticationServiceFactory(db)
	authenticationServiceEndpoints := authentications.MakeEndpoints(authenticationService)
	authentications.NewHTTPHandler(router, authenticationServiceEndpoints)

	//Add metrics
	router.Methods(http.MethodGet).Path("/metrics").Handler(promhttp.Handler())

	println("Starting server on port: " + commons.Port)
	if err := http.ListenAndServe(commons.Port, router); err != nil {
		println(err.Error())
	}

}
