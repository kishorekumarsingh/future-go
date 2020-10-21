package main

import (
	"getmega/future"
	"testing"
)

func TestRunning(t *testing.T) {
	wikiPageURLs := []string{"http://www.google.com", "http://www.facebook.com", "http://www.fakebook.com", "http://en.wikipedia.org/wiki/this_page_does_not_exist", "http://www.instagram.com", "http://www.getmega.com"}
	tasks := []*future.Task{}
	for _, url := range wikiPageURLs {
		task := future.Submit(getWikiPageExistence, url, 10)
		status := task.Running()
		if status == false && task.Result() == nil {
			t.Errorf("got %v, want %v", status, true)
		}
		tasks = append(tasks, task)
	}
}

func TestResult(t *testing.T) {
	wikiPageURLs := []string{"http://www.google.com", "http://www.facebook.com", "http://www.fakebook.com", "http://en.wikipedia.org/wiki/this_page_does_not_exist", "http://www.instagram.com", "http://www.getmega.com"}
	tasks := []*future.Task{}
	for _, url := range wikiPageURLs {
		task := future.Submit(getWikiPageExistence, url, 5)
		tasks = append(tasks, task)
	}

	for _, task := range tasks {
		if task.Result() == nil {
			t.Errorf("Task did not return a result.")
		}
	}
}

func TestException(t *testing.T) {
	wikiPageURLs := []string{"http://www.google.com", "http://gasjhidqwdhui.quxhqi", "http://www.facebook.com", "http://www.fakebook.com", "http://en.wikipedia.org/wiki/this_page_does_not_exist", "http://www.instagram.com", "http://www.getmega.com"}
	tasks := []*future.Task{}
	for _, url := range wikiPageURLs {
		task := future.Submit(getWikiPageExistence, url, 1)
		tasks = append(tasks, task)
	}

	flag := 0
	for _, task := range tasks {
		if task.Exception() != nil {
			flag = 1
		}
	}

	if flag == 0 {
		t.Errorf("Expected an exception, got none")
	}
}
