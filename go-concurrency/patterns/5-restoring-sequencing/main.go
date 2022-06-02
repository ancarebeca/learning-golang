package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

// the boring function return a channel to communicate with it.
func boring(msg string) <-chan Message { // <-chan Message means receives-only channel of Message.
	c := make(chan Message)
	waitForIt := make(chan bool) // share between all messages
	go func() {                  // we launch goroutine inside a function.
		for i := 0; ; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			// every time the goroutine send message.
			// This code waits until the value to be received.
			<-waitForIt
		}

	}()
	return c // return a channel to caller.
}

type Message struct {
	str  string
	wait chan bool
}

func main() {
	c := fanIn(boring("Joe"), boring("Ahn"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're both boring. I'm leaving")
}
