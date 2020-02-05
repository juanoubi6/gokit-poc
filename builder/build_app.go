package builder

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gokit-poc/authentications"
	"gokit-poc/commons"
	"gokit-poc/users"
	"net/http"
)

func BuildAppRouter(db *gorm.DB) http.Handler {
	router := mux.NewRouter()

	//Create user service, it's endpoints and add these to the router.
	userService := users.UserServiceFactory(db)
	userServiceEndpoints := users.MakeEndpoints(userService)
	users.AddHTTPHandlersToRouter(router, userServiceEndpoints)

	//Create authentication service, it's endpoints and add these to the router.
	authenticationService := authentications.AuthenticationServiceFactory(db)
	authenticationServiceEndpoints := authentications.MakeEndpoints(authenticationService)
	authentications.AddHTTPHandlersToRouter(router, authenticationServiceEndpoints)

	//Add metrics
	router.Methods(http.MethodGet).Path("/metrics").Handler(promhttp.Handler())

	return router

}

func CreateDatabase(dbUri string) *gorm.DB {
	return commons.CreateDatabase("sqlite3", dbUri)
}
