package server

import (
	"fmt"
	"net/http"

	"github.com/nicoalimin/weathering/server/pkg/conf"
	"github.com/nicoalimin/weathering/server/pkg/core"

	"github.com/nicoalimin/weathering/server/pkg/handler"

	"github.com/gorilla/mux"
)

var (
	config conf.Conf
)

func init() {
	config = conf.NewConf()
}

const (
	v1Prefix = "/v1" // Version of the API
)

func loadRoutes(r *mux.Router) {
	v1Router := r.PathPrefix(v1Prefix).Subrouter()

	hClient := &http.Client{}
	coreWeather := core.ServiceWeather{
		HTTPClient:     hClient,
		WeathersAPIURL: "http://api.openweathermap.org/data/2.5/weather",
		AccessKey:      config.OpenWeatherKey,
	}
	handlerHealth := handler.Health{}
	handlerWeather := handler.Weather{
		Weather: &coreWeather,
	}

	healthRouter := v1Router.PathPrefix("/health").Subrouter()
	healthRouter.HandleFunc("", handlerHealth.HealthCheck).Methods(http.MethodGet)

	weatherRouter := v1Router.PathPrefix("/weather").Subrouter()
	weatherRouter.HandleFunc("", handlerWeather.GetWeatherByCity).Methods(http.MethodGet)

}

// Execute runs a HTTP server that contains the backend for the ea game review app
func Execute() {

	// initialize router
	router := mux.NewRouter()

	// load custom routes
	loadRoutes(router)

	// initialize http server configs
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", config.BackendPort),
		Handler: router,
	}

	// start http server
	fmt.Printf("HTTP Server listening on port: %s\n", config.BackendPort)
	server.ListenAndServe()
}
