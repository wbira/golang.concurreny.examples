package weather_test

import (
	"fmt"
	"testing"

	"golang.concurrency/weather"
)

func Test(t *testing.T) {
	locations := []string{"London", "Liverpool"}
	ch := weather.FetchWeatherForLocactions(locations)

	for i := 0; i < len(locations); i++ {
		no := i
		fmt.Println(no)
	}

	for i := 0; i < len(locations); i++ {
		fmt.Println(<-ch)
	}
}
