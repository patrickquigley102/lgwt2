package main

import (
	"log"

	"github.com/patrickquigley102/lgwt2/concurrency/concurrency"
)

func main() {
	fs := concurrency.OSFS{}
	_, err := concurrency.DownloadCSV(fs)
	if err != nil {
		log.Fatal(err)
	}
}
