package main

import (
	"fmt"
	"log"

	"github.com/patrickquigley102/lgwt2/concurrency/concurrency"
)

const fileURL = "https://gist.githubusercontent.com" +
	"/bejaneps/ba8d8eed85b0c289a05c750b3d825f61/raw/" +
	"6827168570520ded27c102730e442f35fb4b6a6d/websites.csv"

func main() {
	fs := concurrency.OSFS{}
	filePathName, err := concurrency.DownloadCSV(fs, fileURL)
	checkErr(err)

	urls, err := concurrency.FileToArray(filePathName)
	checkErr(err)

	websiteChecker := concurrency.WebsiteCheck{}
	responses := concurrency.CheckWebsites(websiteChecker, urls)
	fmt.Printf("%v", responses) //nolint
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
