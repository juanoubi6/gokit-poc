package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gokit-poc/users"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	userService := users.CreateUserService()

	http.Handle("/user", users.CreateUserHandler(userService))
	http.Handle("/metrics", promhttp.Handler())

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
