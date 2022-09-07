// Package racer is a package.
package racer

import (
	"errors"
	"log"
	"net/http"
	"time"
)

var errTimeout = errors.New("timed out waiting")

// Racer takes two urls, returns fastest response.
// Errors if both timeout at 10s.
func Racer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", errTimeout
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
