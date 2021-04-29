# Limit your crawler

Given is a crawler (modified from the Go tour) that requests pages
excessively. However, we don't want to burden the webserver too
much. Your task is to change the code to limit the crawler to at most
one page per second, while maintaining concurrency (in other words,
Crawl() must be called concurrently)

## Hint

This exercise can be solved in 3 lines only. If you can't do
it, have a look at this:
https://github.com/golang/go/wiki/RateLimiting
