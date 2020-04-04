package main

import (
	"fmt"

	"golang.concurrency/weather"
)

func main() {
	locations := []string{"London", "Liverpool"}
	ch := weather.FetchWeatherForLocactions(locations)

	for i := 0; i < len(locations); i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}
