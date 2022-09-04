// Package concurrency is a package.
package concurrency

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

type result struct {
	string
	bool
}

/*
type websiteCheck struct {
	url string
}

func (wc websiteCheck) Check(url string) bool {
	return false
}
*/
