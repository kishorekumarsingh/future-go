package main

import (
	"fmt"
	"getmega/future"
	"log"
	"net/http"
	"time"
)

// get_wiki_page_existence returns whether a page exists
func getWikiPageExistence(wikiPageURL string, timeout time.Duration) interface{} {
	client := http.Client{
		Timeout: timeout * time.Second,
	}
	resp, err := client.Get(wikiPageURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	pageStatus := "unknown"
	if resp.StatusCode == 200 {
		pageStatus = "exists"
	}
	if resp.StatusCode == 404 {
		pageStatus = "does not exist"
	}
	fmt.Println(wikiPageURL + " - " + pageStatus)
	return "FINISHED"

}

func main() {
	wikiPageURLs := []string{"http://www.google.com", "http://www.facebook.com", "http://www.fakebook.com", "http://en.wikipedia.org/wiki/this_page_does_not_exist", "http://www.instagram.com", "http://www.getmega.com"}

	for _, url := range wikiPageURLs {
		// getWikiPageExistence(url, 20)
		task := future.Submit(getWikiPageExistence, url, 20)
		task.Exception()
	}
}
