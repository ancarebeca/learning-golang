package main

import (
	"sync"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	u0 := User{ID: 0, IsPremium: false}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func () {
			HandleRequest(shortProcess, &u0)
			wg.Done()
		}()
	}
	wg.Wait()
}