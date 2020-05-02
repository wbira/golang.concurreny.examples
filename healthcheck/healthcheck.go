package healthcheck

import (
	"net/http"
	"sync"
)

type Result struct {
	Status int
	Error  error
}

func CheckServices(urls []string) <-chan *Result {
	results := make(chan *Result)
	go func() {
		for _, url := range urls {
			response, err := http.Get(url)
			results <- &Result{response.StatusCode, err}
		}
		close(results)
	}()

	return results
}

func CheckServicesAsync(urls []string) <-chan *Result {
	var wg sync.WaitGroup
	results := make(chan *Result)
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			response, err := http.Get(url)
			results <- &Result{response.StatusCode, err}
			wg.Done()
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}
