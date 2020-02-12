package main

import (
	"gokit-poc/builder"
	"gokit-poc/commons"
	"net/http"
	"os"
	"path"
)

func main() {
	// Create DB
	wd, _ := os.Getwd()
	dbUri := path.Join(wd, commons.DatabaseURI)
	db := builder.CreateDatabase(dbUri)

	// Create router
	router := builder.BuildAppRouter(db)

	// Start listening
	println("Starting server on port: " + commons.Port)
	http.ListenAndServe(commons.Port, router)

}
