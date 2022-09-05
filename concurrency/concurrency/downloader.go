package concurrency

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	// FileURL is a csv of 1000 URLs.
	FileURL = "https://gist.githubusercontent.com" +
		"/bejaneps/ba8d8eed85b0c289a05c750b3d825f61/raw/" +
		"6827168570520ded27c102730e442f35fb4b6a6d/websites.csv"
	fileName = "urls.csv"
	filePath = "/src/personal/go/lgwt2/concurrency/concurrency/"
)

/*
create file
download csv into file
*/

// DownloadCSV to file on os.
func DownloadCSV(fs FileSystem) (*os.File, error) {
	file, err := fs.Create(fullFilePath(fs))
	if err != nil {
		err = fmt.Errorf("failed to create blank file. %w", err)
	}

	return file, err
}

func fullFilePath(fs FileSystem) string {
	dirname, err := fs.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return dirname + filePath + fileName
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

/*
// DownloadCSV downloads a zipped csv of 1 million URLs, unzips and saves the
// csv whilst deleting the zip file. Returning the File object of the csv.
// File downloaded from concurrency.FileURL.
func DownloadCSV() (*os.File, error) {
	var file *os.File
	var err error

	if !csvFileExists() {
		file, err = downloadZip()
	}

	return file, err
}

func downloadZip() (*os.File, error) {
	var file *os.File
	var err error

	return file, err
}

func csvFileExists() bool {
	_, err := os.Stat(fullFilePath())

	return !errors.Is(err, os.ErrNotExist)
}
*/
