package main

import (
	"log"

	"github.com/patrickquigley102/lgwt2/concurrency/concurrency"
)

const fileURL = "https://gist.githubusercontent.com" +
	"/bejaneps/ba8d8eed85b0c289a05c750b3d825f61/raw/" +
	"6827168570520ded27c102730e442f35fb4b6a6d/websites.csv"

func main() {
	fs := concurrency.OSFS{}
	_, err := concurrency.DownloadCSV(fs, fileURL)
	if err != nil {
		log.Fatal(err)
	}
}
