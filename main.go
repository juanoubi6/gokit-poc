package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gokit-poc/commons"
	"gokit-poc/users"
	"net/http"
)

func main() {
	userService := users.UserServiceFactory()

	http.Handle("/user", users.CreateUserHandler(userService))
	http.Handle("/metrics", promhttp.Handler())

	println("Starting server on port: " + commons.Port)
	http.ListenAndServe(":8080", nil)
}
