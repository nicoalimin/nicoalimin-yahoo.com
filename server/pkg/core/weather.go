package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// compile time check whether the ServiceWeather business logic layer satisfies the
// weathers API
var _ CoreWeather = (*ServiceWeather)(nil)

// ServiceWeather contains the core business logic regarding the Weathers API
type ServiceWeather struct {
	HttpClient     http.Client
	WeathersAPIURL string
	AccessKey      string
}

// GetWeatherByCity returns a weathers summary for a given city
func (sw *ServiceWeather) GetWeatherByCity(city string) (*Weather, error) {
	req, err := http.NewRequest(http.MethodPost, sw.WeathersAPIURL+"?q=Vancouver&appid="+sw.AccessKey, nil)
	if err != nil {
		return nil, fmt.Errorf("error obtaining weathers")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := sw.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w, %v", ErrQueryingExternalWeathersAPI, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w, %s", ErrQueryingExternalWeathersAPI, "open weather server returns a non-200 status code")
	}
	defer resp.Body.Close()

	var openWeatherResponse OpenWeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&openWeatherResponse)
	if err != nil {
		return nil, fmt.Errorf("%w, %v", ErrQueryingExternalWeathersAPI, err)
	}

	return &Weather{
		City:        city,
		Temperature: fmt.Sprintf("%f", openWeatherResponse.Main.Temp),
		Condition:   openWeatherResponse.Weather[0].Main,
	}, nil
}
