package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	LOCATION_API_URL = "https://www.metaweather.com/api/location/search/?query=%s"
	WEATHER_API_URL  = "https://www.metaweather.com/api/location/%d/"
)

func FetchWeatherForLocactions(locations <-chan *FetchLocationResult) <-chan *FetchWeatherResult {
	weatherResults := make(chan *FetchWeatherResult)
	for result := range locations {
		go func(result *FetchLocationResult) {
			if result.Error != nil {
				fmt.Printf("if res %v\n", result)
				weatherResults <- &FetchWeatherResult{Error: result.Error}
			} else {
				weather, err := fetchWeather(result.Location.Id)
				fmt.Println(weather)
				weatherResults <- &FetchWeatherResult{Weather: weather, Error: err}
			}
		}(result)
	}
	return weatherResults

}

func FetchLocationIds(locationNames []string) <-chan *FetchLocationResult {
	results := make(chan *FetchLocationResult)
	for _, location := range locationNames {
		fmt.Printf("location name %v\n", location)
		go func(location string) {
			l, err := fetchLocation(location)
			results <- &FetchLocationResult{Location: l, Error: err}

		}(location)
	}

	return results
}

func fetchWeather(id int) (*Weather, error) {
	uri := fmt.Sprintf(WEATHER_API_URL, id)
	response, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("get weather: %w", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read weather body: %w", err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, fmt.Errorf("unmarshal weather: %w", err)
	}
	return &weather, nil
}

func fetchLocation(locationName string) (*Location, error) {
	uri := fmt.Sprintf(LOCATION_API_URL, locationName)
	response, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("get location: %w", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read location body: %w", err)
	}

	var locations []Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return nil, fmt.Errorf("unmarshal location %w", err)
	}

	if len(locations) > 0 {
		return &locations[0], nil
	}
	return nil, fmt.Errorf("empty location slice")
}
