package handler

import (
	"github.com/nicoalimin/weathering/server/pkg/core"
	"net/http"
)

type Weather struct {
	Weather core.CoreWeather
}

func (wh *Weather) GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	return
}