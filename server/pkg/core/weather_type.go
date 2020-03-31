package core

type Weather struct {
	City        string
	Temperature string
}

type CoreWeather interface {
	GetWeatherByCity(city string) *Weather
}
