package handler

import (
	"errors"
	"github.com/nicoalimin/weathering/server/pkg/core"
	"net/http"
)

// Weather is the HTTP layer of the weathers API. It converts internal error objects into
// error codes and messages that makes sense for the consumer of the API
type Weather struct {
	Weather core.CoreWeather
}

// GetWeatherByCity returns a weather given a city
func (wh *Weather) GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city, ok := r.URL.Query()["city"]
	if !ok || len(city) < 1 {
		RespondWithError(w, http.StatusBadRequest, "city must be provided")
		return
	}
	weather, err := wh.Weather.GetWeatherByCity(city[0])

	switch {
	case err == nil:
		RespondWithSuccess(w, http.StatusOK, weather)
	case errors.Is(err, core.ErrBadRequest):
		RespondWithError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, core.ErrCityNotFound):
		RespondWithError(w, http.StatusNotFound, err.Error())
	default:
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
}