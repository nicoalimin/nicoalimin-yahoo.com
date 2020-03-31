package core

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	sampleOpenWeatherResponse = `{
    "coord": {
        "lon": -123.12,
        "lat": 49.25
    },
    "weather": [
        {
            "id": 521,
            "main": "Rain",
            "description": "shower rain",
            "icon": "09d"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 281.9,
        "feels_like": 275.83,
        "temp_min": 279.82,
        "temp_max": 283.71,
        "pressure": 1015,
        "humidity": 61
    },
    "visibility": 24140,
    "wind": {
        "speed": 6.2,
        "deg": 220
    },
    "rain": {
        "1h": 0.51
    },
    "clouds": {
        "all": 75
    },
    "dt": 1585692448,
    "sys": {
        "type": 1,
        "id": 954,
        "country": "CA",
        "sunrise": 1585662621,
        "sunset": 1585708932
    },
    "timezone": -25200,
    "id": 6173331,
    "name": "Vancouver",
    "cod": 200
}`
)

func Test_OpenWeather_API(t *testing.T) {
	city := "vancouver"

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("\nBlah\n")
		intendedURL := "/weather?q=" + city + "&appid="
		if strings.HasPrefix(req.URL.String(), intendedURL) {
			t.Errorf("expected url doesn't match actual. \ngot: %s, \nwant: %s", req.URL.String(), intendedURL)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(sampleOpenWeatherResponse))
	}))
	defer server.Close()

	serviceWeather := ServiceWeather{
		HTTPClient:     server.Client(),
		WeathersAPIURL: server.URL,
		AccessKey:      "",
	}
	_, err := serviceWeather.GetWeatherByCity(city)

	if err != nil {
		t.Errorf("expected no errors getting weather by city. \nGot: %v", err)
	}
}
