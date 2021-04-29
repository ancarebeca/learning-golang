//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, c chan<- Tweet) {
	defer close(c)

	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}
		c <- *tweet
	}

}

func consumer(c <-chan Tweet) {
	time.Sleep(1 * time.Second)

	for t := range c {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	c := make(chan Tweet)

	// Producer
	go producer(stream, c)

	// Consumer
	consumer(c)

	fmt.Printf("Process took %s\n", time.Since(start))
}
