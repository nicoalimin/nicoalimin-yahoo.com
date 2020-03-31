package core

import (
	"errors"
)

const (
	// ErrCityNotFound is the error if the given city is invalid
	ErrCityNotFound = errors.New("Desired city is not found")
)

// Weather contains a location and the current summary of the weather
type Weather struct {
	City        string // location of the desired weather check
	Temperature string // temperature of a given location
	Condition   string // whether given location is sunny, rainy, cloudy, etc.
}

// CoreWeather contains business logic of the Weather's API
type CoreWeather interface {
	// GetWeatherByCity obtains a particular city's current weather. Returns a weather
	// object or an error.
	GetWeatherByCity(city string) (*Weather, error)
}
