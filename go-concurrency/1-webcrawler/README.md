Web Crawler

In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not
safe for concurrent use!


Documentation:
 - [Can we have multiple channels having multiple go routines in golang](https://stackoverflow.com/questions/70542122/can-we-have-multiple-channels-having-multiple-go-routines-in-golang)
 - [Where to put wg.Add()](https://stackoverflow.com/questions/65213707/where-to-put-wg-add)
 - [Solution with channels](https://micknelson.wordpress.com/2012/11/14/a-tour-of-go-the-web-crawler-exercise/)
 - [Solution with channels - code](https://github.com/absoludity/go-tour-exercises/blob/e749a2c1a9602437e186de01a7a9e40927aecf76/70-web-crawler.go#L44)