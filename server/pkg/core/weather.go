package core

var _ CoreWeather = (*ServiceWeather)(nil)

type ServiceWeather struct{}

func (sw *ServiceWeather) GetWeatherByCity(city string) *Weather {
	return &Weather{}
}
