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

func FetchWeatherForLocactions(locations []string) <-chan *Weather {
	ch := make(chan *Weather)
	ids := fetchLocationIds(locations)
	for i := 0; i < len(locations); i++ {
		go func() {
			t := <-ids
			ch <- fetchWeather(t.Id)
		}()
	}
	return ch
}

func fetchLocationIds(locationNames []string) <-chan *Location {
	locationsIds := make(chan *Location)
	for _, location := range locationNames {
		location := location
		go func() {
			locationsIds <- fetchLocation(location)
		}()
	}
	return locationsIds
}

func fetchWeather(id int) *Weather {
	uri := fmt.Sprintf(WEATHER_API_URL, id)
	response, err := http.Get(uri)
	if err != nil {
		//todo error handling
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//todo error handling
		return nil
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		//todo error handling
		return nil
	}
	return &weather
}

func fetchLocation(locationName string) *Location {
	uri := fmt.Sprintf(LOCATION_API_URL, locationName)
	response, err := http.Get(uri)
	if err != nil {
		//todo error handling
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//todo error handling
		return nil
	}

	var locations []Location
	err = json.Unmarshal(body, &locations)
	if err != nil {
		//todo error handling
		return nil
	}

	if len(locations) > 0 {
		return &locations[0]
	}
	return nil
}
