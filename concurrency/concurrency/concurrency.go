// Package concurrency is a package.
package concurrency

// CheckWebsites checks URLS using the WebsiteChecker, returning map of results.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc.Check(url)
	}

	return results
}

// WebsiteChecker checks a URL.
type WebsiteChecker interface {
	Check(string) bool
}

/*
type websiteCheck struct {
	url string
}

func (wc *websiteCheck) Check(url string) bool {
	return false
}
*/
