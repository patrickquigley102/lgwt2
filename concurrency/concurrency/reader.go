package concurrency

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const minimumCSVColumns = 2

var errFormat = errors.New("wrong csv format")

// FileToArray takes a csv file of the following format and returns an array of
// the 2nd column string.
// object,"string in quotes",object.
func FileToArray(fileName string) ([]string, error) {
	file, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		err = fmt.Errorf("failed to open file. %w", err)

		return nil, err
	}

	var urls []string
	csvReader := csv.NewReader(file)
	for {
		rec, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			err = fmt.Errorf("failed to read csv line. %w", err)

			return nil, err
		}

		if len(rec) < minimumCSVColumns {
			err = fmt.Errorf("wrap it up. %w", errFormat)

			return nil, err
		}

		urls = append(urls, "http://"+rec[1])
	}

	_ = file.Close()

	return urls, nil
}
