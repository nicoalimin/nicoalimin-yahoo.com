package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nicoalimin/weathering/server/pkg/core"

	"github.com/golang/mock/gomock"
)

func Test_GetWeatherByCity(t *testing.T) {
	cases := map[string]struct {
		errorReturned      error
		expectedStatusCode int
		city               string
		numDatabaseCall    int
	}{
		"getWeatherSuccess": {
			errorReturned:      nil,
			expectedStatusCode: http.StatusOK,
			city:               "vancouver",
			numDatabaseCall:    1,
		},
		"getWeatherMissingCityError": {
			errorReturned:      nil,
			expectedStatusCode: http.StatusBadRequest,
			city:               "",
			numDatabaseCall:    0,
		},
		"getWeatherBadRequestError": {
			errorReturned:      core.ErrBadRequest,
			expectedStatusCode: http.StatusBadRequest,
			city:               "vancouver",
			numDatabaseCall:    1,
		},
		"getWeatherCityNotFoundError": {
			errorReturned:      core.ErrCityNotFound,
			expectedStatusCode: http.StatusNotFound,
			city:               "vancouver",
			numDatabaseCall:    1,
		},
		"getWeatherGenericError": {
			errorReturned:      errors.New("unexpected error"),
			expectedStatusCode: http.StatusInternalServerError,
			city:               "vancouver",
			numDatabaseCall:    1,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCoreWeather := core.NewMockCoreWeather(ctrl)

			handlerWeather := Weather{
				Weather: mockCoreWeather,
			}

			if c.errorReturned == nil {
				mockCoreWeather.EXPECT().GetWeatherByCity(gomock.Any()).Times(c.numDatabaseCall).Return(&core.Weather{}, c.errorReturned)
			} else {
				mockCoreWeather.EXPECT().GetWeatherByCity(gomock.Any()).Times(c.numDatabaseCall).Return(nil, c.errorReturned)
			}

			urlPath := "/weather"
			if c.city != "" {
				urlPath = urlPath + "?city=" + c.city
			}
			r := httptest.NewRequest(http.MethodGet, urlPath, nil)
			w := httptest.NewRecorder()

			handlerWeather.GetWeatherByCity(w, r)

			if w.Code != c.expectedStatusCode {
				t.Errorf("expected status code doesn't match actual. \nGot: %d \nWant: %d", w.Code, c.expectedStatusCode)
			}
		})
	}
}
