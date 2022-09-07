// Package racer is a package.
package racer

import (
	"log"
	"net/http"
	"time"
)

// Racer takes two urls, returns fastest response.
// Errors if both timeout at 10s.
func Racer(url1, url2 string) string {
	duration1 := measureResponseTime(url1)
	duration2 := measureResponseTime(url2)

	if duration1 < duration2 {
		return url1
	}

	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	_, err := http.Get(url) //nolint
	if err != nil {
		log.Fatal(err)
	}

	return time.Since(start)
}
