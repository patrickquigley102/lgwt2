// Package racer is a package.
package racer

import (
	"log"
	"net/http"
)

// Racer takes two urls, returns fastest response.
// Errors if both timeout at 10s.
func Racer(url1, url2 string) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})

	go func() {
		_, err := http.Get(url) //nolint
		if err != nil {
			log.Fatal(err)
		}
		close(channel)
	}()

	return channel
}
