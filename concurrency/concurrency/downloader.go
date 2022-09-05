package concurrency

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	fileName = "urls.csv"
	filePath = "/src/personal/go/lgwt2/concurrency/concurrency/"
)

// DownloadCSV to file on os.
func DownloadCSV(fs FileSystem, url string) (*os.File, error) {
	file, err := fs.Create(fullFilePath(fs))
	if err != nil {
		err = fmt.Errorf("failed to create blank file. %w", err)

		return nil, err
	}

	client := httpClient()
	req, err := httpRequest(url)
	if err != nil {
		return nil, err
	}

	resp, err := getFile(client, req)
	if err != nil {
		return nil, err
	}

	file, err = putFile(file, resp)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func fullFilePath(fs FileSystem) string {
	dirname, _ := fs.UserHomeDir()

	return dirname + filePath + fileName
}

func getFile(c *http.Client, r *http.Request) (*http.Response, error) {
	resp, err := c.Do(r)
	if err != nil {
		err = fmt.Errorf("failed to GET csv file from URL. %w", err)
	}

	return resp, err
}

func putFile(file *os.File, resp *http.Response) (*os.File, error) {
	_, err := io.Copy(file, resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to copy csv from URL into file. %w", err)
	}

	return file, err
}

func httpRequest(url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		context.Background(), http.MethodGet, url, nil,
	)
	if err != nil {
		err = fmt.Errorf("failed to build http request. %w", err)
	}

	return req, err
}

func httpClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

// FileSystem is a file system.
type FileSystem interface {
	Create(name string) (*os.File, error)
	UserHomeDir() (string, error)
}

// OSFS is a file system on host OS.
type OSFS struct{}

// Create calls os.Create().
func (OSFS) Create(name string) (*os.File, error) {
	file, err := os.Create(filepath.Clean(name))
	if err != nil {
		err = fmt.Errorf("failed to create blank file. %w", err)
	}

	return file, err
}

// UserHomeDir calls os.UserHomeDir().
func (OSFS) UserHomeDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf("failed to get home dir. %w", err)
	}

	return dir, err
}
