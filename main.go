package main

import (
	"fmt"
	"getmega/future"
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
		return err
	}
	defer resp.Body.Close()
	pageStatus := "unknown"
	if resp.StatusCode == 200 {
		pageStatus = "exists"
	}
	if resp.StatusCode == 404 {
		pageStatus = "does not exist"
	}
	return (wikiPageURL + " - " + pageStatus)

}

func main() {
	wikiPageURLs := []string{"http://www.google.com", "http://www.facebook.com", "http://www.fakebook.com", "http://en.wikipedia.org/wiki/this_page_does_not_exist", "http://www.instagram.com", "http://www.getmega.com"}

	tasks := []*future.Task{}
	for _, url := range wikiPageURLs {
		task := future.Submit(getWikiPageExistence, url, 10)
		tasks = append(tasks, task)
	}

	for _, task := range tasks {
		if task.Result() != nil {
			fmt.Println(task.Result())
		}
	}
}
