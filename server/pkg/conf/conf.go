package conf

import (
	"github.com/kelseyhightower/envconfig"
)

// Conf contains the environment variables required to properly run the application
type Conf struct {
	OpenWeatherKey string `envconfig:"OPEN_WEATHER_KEY" default:"df04f1248445346d798760340bd65a74"`
	BackendPort    string `envconfig:"BACKEND_PORT" default:"80"`
}

// NewConf is the initializer that returns a Conf object, that is populated with the
// desired environment variables
func NewConf() Conf {
	c := Conf{}
	envconfig.MustProcess("", &c)
	return c
}
