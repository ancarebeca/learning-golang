package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, startCrawl func(string) bool, crawlComplete chan string) {

	if depth <= 0 {
		crawlComplete <- url
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		crawlComplete <- url
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if startCrawl(u) {
			go Crawl(u, depth-1, fetcher, startCrawl, crawlComplete)
		}
	}
	crawlComplete <- url
}

type StartCrawlData struct {
	Url       string
	ReplyChan chan bool
}

func processCrawls(startCrawlReceiver chan StartCrawlData, crawlCompleteReceiver chan string, quitNow chan bool) {
	// The urlsCrawled map tells us:
	//   a) Whether a URL still needs to be crawled (ie. it is not present in the map).
	//   b) Whether a URL is being crawled currently (ie. it has a key, but value is false), or
	//   c) Whether the crawling of a URL is complete (ie. it has a key and the value is true).
	urlsCrawled := make(map[string]bool)
	for {
		select {

		case startCrawlData := <-startCrawlReceiver:
			_, alreadyStarted := urlsCrawled[startCrawlData.Url]
			if !alreadyStarted {
				urlsCrawled[startCrawlData.Url] = false
			}
			startCrawlData.ReplyChan <- !alreadyStarted

		case crawledUrl := <-crawlCompleteReceiver:
			urlsCrawled[crawledUrl] = true

			// If all the registered crawls have completed then
			// we signal that we can quit now.
			// Is there something similar to python's all([]bool) built-in?
			allChecked := true
			for _, checkComplete := range urlsCrawled {
				allChecked = allChecked && checkComplete
			}
			if allChecked {
				quitNow <- true
				return
			}
		}
	}

}

func main() {
	startCrawl := make(chan StartCrawlData)
	crawlComplete := make(chan string)
	quitNow := make(chan bool)
	go processCrawls(startCrawl, crawlComplete, quitNow)

	// Returns whether a crawl should be started for a given
	// URL.
	startCrawlFn := func(url string) bool {
		resultChan := make(chan bool)
		startCrawl <- StartCrawlData{url, resultChan}
		return <-resultChan
	}

	Crawl("http://golang.org/", 4, fetcher, startCrawlFn,
		crawlComplete)

	<-quitNow
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
