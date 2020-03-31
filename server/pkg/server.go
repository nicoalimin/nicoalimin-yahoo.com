package server

import (
	"fmt"
	"net/http"

	"github.com/nicoalimin/weathering/server/pkg/core"

	"github.com/nicoalimin/weathering/server/pkg/handler"

	"github.com/gorilla/mux"
)

const (
	v1Prefix = "/v1" // Version of the API
)

func loadRoutes(r *mux.Router) {
	v1Router := r.PathPrefix(v1Prefix).Subrouter()

	coreWeather := core.ServiceWeather{}
	handlerHealth := handler.Health{}

	handlerRouter := v1Router.PathPrefix("/health").Subrouter()
	handlerRouter.HandleFunc("", handlerHealth.HealthCheck).Methods(http.MethodGet)
}

// Execute runs a HTTP server that contains the backend for the ea game review app
func Execute() {

	// initialize router
	router := mux.NewRouter()

	// load custom routes
	loadRoutes(router)

	// initialize http server configs
	port := "3031"
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// start http server
	fmt.Printf("HTTP Server listening on port: %s\n", port)
	server.ListenAndServe()
}
