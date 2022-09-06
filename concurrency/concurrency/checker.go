// Package concurrency is a package.
package concurrency

import (
	"net/http"
	"time"
)

// CheckWebsites checks URLS using the WebsiteChecker, returning map of results.
func CheckWebsites(webCheck WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, webCheck.Check(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// WebsiteChecker checks a URL.
type WebsiteChecker interface {
	Check(string) bool
}

// WebsiteCheck checks a URL.
type WebsiteCheck struct{}

// Check a URL. true if 200, false otherwise.
func (wc WebsiteCheck) Check(url string) bool {
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(url) //nolint
	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}

const timeout = 2 * time.Second

type result struct {
	string
	bool
}
