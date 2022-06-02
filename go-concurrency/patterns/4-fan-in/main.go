package main

import (
	"fmt"
	"math/rand"
	"time"
)

//3-channel-as-service makes Joe and Ann count in lockstep. We can instead use a 4-fan-in function to let whosoever is ready to talk

// the boring function return a channel to communicate with it.
func boring(msg string) <-chan string { // <-chan string means receives-only channel of string.
	c := make(chan string)
	go func() { // we launch goroutine inside a function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c // return a channel to caller.
}

// <-chan string only get the receive value
// fanIn spawns 2 goroutines to reads the value from 2 channels
// then it sends to value to result channel( `c` channel)
func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for { // infinite loop to read value from channel.
			v1 := <-c1 // read value from c1. This line will wait when receiving value.
			c <- v1
		}
	}()
	go func() {
		for {
			c <- <-c2 // read value from c2 and send it to c
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ahn"))
	// merge 2 channels into 1 channel
	//c := fanInSimple(boring("Joe"), boring("Ahn"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // now we can read from 1 channel
	}
	fmt.Println("You're both boring. I'm leaving")
}
