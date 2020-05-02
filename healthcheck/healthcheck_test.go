package healthcheck_test

import (
	"fmt"
	"testing"

	"golang.concurrency/healthcheck"
)

var urls = []string{"https://google.com", "https://gobyexample.com/", "https://amazon.com", "http://wbira.github.io/", "https://wykop.pl"}

func BenchmarkFetchingInSepperateGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(b, healthcheck.CheckServicesAsync(urls))
	}

}

func BenchmarkFetchingInOneGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runTest(b, healthcheck.CheckServices(urls))
	}
}

func runTest(b *testing.B, ch <-chan *healthcheck.Result) {
	b.StartTimer()
	for result := range ch {
		if result.Error != nil {
			fmt.Printf("Error %v\n", result.Error)
		} else {
			fmt.Printf("Response: %v\n", result.Status)
		}
	}
	b.StopTimer()
}
