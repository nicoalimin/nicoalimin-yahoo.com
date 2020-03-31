package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func loadRoutes(r *mux.Router) {
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
