package core

import (
	"errors"
)

var (
	// ErrCityNotFound is the error if the given city is invalid
	ErrCityNotFound = errors.New("desired city is not found")
	// ErrBadRequest gets thrown if the request does not come with the necessary arguments
	ErrBadRequest = errors.New("error validating the parameters")
	// ErrQueryingExternalWeathersAPI gets thrown if there is an error from the external
	// weathers API
	ErrQueryingExternalWeathersAPI = errors.New("error querying external weathers API")
)

// Weather contains a location and the current summary of the weather
// TODO: Improve Weather so it returns valuable weather info too (sunrise, humidity, etc)
type Weather struct {
	City        string `json:"city"`        // location of the desired weather check
	Temperature string `json:"temperature"` // temperature of a given location
	Condition   string `json:"condition"`   // whether given location is sunny, rainy, cloudy, etc.
}

// OpenWeatherResponse is the external Weathers API that is used to query weathers by city
type OpenWeatherResponse struct {
	Weather []struct {
		ID   int    `json:"id"`
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

// CoreWeather contains business logic of the Weather's API
//go:generate mockgen -destination=./core_weather.mock.go -self_package=github.com/nicoalimin/weathering/server/pkg/core -package core github.com/nicoalimin/weathering/server/pkg/core CoreWeather
type CoreWeather interface {
	// GetWeatherByCity obtains a particular city's current weather. Returns a weather
	// object or an error.
	GetWeatherByCity(city string) (*Weather, error)
}
