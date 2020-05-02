package weather

import "time"

type Location struct {
	Id int `json:"woeid,omitempty"`
}

type FetchLocationResult struct {
	Location *Location
	Error    error
}

type FetchWeatherResult struct {
	Weather *Weather
	Error   error
}

type ConsolidatedWeather struct {
	ID                   string `json:"id"`
	WeatherStateName     string `json:"weather_state_name"`
	WeatherStateAbbr     string `json:"weather_state_abbr"`
	WindDirectionCompass string `json:"wind_direction_compass"`
	Created              string `json:"created"`
	ApplicableDate       string `json:"applicable_date"`
	MinTemp              string `json:"min_temp"`
	MaxTemp              string `json:"max_temp"`
	TheTemp              string `json:"the_temp"`
	WindSpeed            string `json:"wind_speed"`
	WindDirection        string `json:"wind_direction"`
	AirPressure          string `json:"air_pressure"`
	Humidity             string `json:"humidity"`
	Visibility           string `json:"visibility"`
	Predictability       string `json:"predictability"`
}

type Weather struct {
	ConsolidatedWeather []ConsolidatedWeather
	Time                time.Time `json:"time"`
	SunRise             time.Time `json:"sun_rise"`
	SunSet              time.Time `json:"sun_set"`
	TimezoneName        string    `json:"timezone_name"`
	Title               string    `json:"title"`
	LocationType        string    `json:"location_type"`
	Woeid               int       `json:"woeid"`
	LattLong            string    `json:"latt_long"`
	Timezone            string    `json:"timezone"`
}
