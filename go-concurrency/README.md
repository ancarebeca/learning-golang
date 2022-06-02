# Go Concurrency Exercises [![Build Status](https://travis-ci.org/mindworker/go-concurrency-exercises.svg?branch=master)](https://travis-ci.org/mindworker/go-concurrency-exercises) [![Go Report Card](https://goreportcard.com/badge/github.com/mindworker/go-concurrency-exercises)](https://goreportcard.com/report/github.com/mindworker/go-concurrency-exercises)
Exercises for Golang's concurrency patterns. Original source [here](https://github.com/loong/go-concurrency-exercises/blob/master/README.md)

## Why
The Go community has plenty resources to read about go's concurrency model and how to use it effectively. But *who actually wants to read all this*!? This repo tries to teach concurrency patterns by following the 'learning by doing' approach.

![Image of excited gopher](https://golang.org/doc/gopher/pkg.png)

## How to take this challenge
1. *Only edit `main.go`* to solve the problem. Do not touch any of the other files.
2. If you find a `*_test.go` file, you can test the correctness of your solution with `go test`
3. If you get stuck, join us on [Slack](https://gophersinvite.herokuapp.com/)! I'm sure there will be people who are happy to give you some code review (if not, find me via @beertocode ;) )

## Overview
| # | Name of the Challenge + URL           | 
| - |:-------------|
| 0 | [Limit your Crawler](0-webcrawler/README.md) |
| 1 | [Producer-Consumer](2-consumer-producer/README.md)  |
| 2 | [Race Condition in Caching Cache](https://github.com/mindworker/go-concurrency-exercises/tree/master/2-race-in-cache#race-condition-in-caching-szenario)  |
| 3 | [Limit Service Time for Free-tier Users](3-limit-service-time/README.md)  |
| 4 | [Graceful SIGINT Killing](4-graceful-sigint/README.md)  |
| 5 | [Clean Inactive Sessions to Prevent Memory Overflow](5-session-cleaner)  |

## License

```
 Copyleft from 2017 Long Hoang

 Everyone is permitted to copy and distribute verbatim or modified 
 copies of this license document, and changing it is allowed as long 
 as the name is changed.

```
Documentation:
- [Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Solutions](https://github.com/loong/go-concurrency-solutions)
- https://github.com/derekahn/go-concurrency-solutions/blob/master/1-producer-consumer/main.go
- [Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs&ab_channel=GoogleDevelopers)

More Exercises:
- https://github.com/andcloudio/go-concurrency-exercises
- [Rate Limiting](https://github.com/golang/go/wiki/RateLimiting)


https://github.com/syafdia/go-exercise/tree/master/src/concurrency
https://medium.com/code-chasm/go-concurrency-pattern-worker-pool-a437117025b1
https://github.com/lotusirous/go-concurrency-patterns/blob/main/7-quit-signal/main.go