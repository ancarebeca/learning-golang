package main

import (
	"testing"
)

func TestHandleRequest(t *testing.T) {
	u0 := User{ID: 0, IsPremium: false}
	//u1 := User{ID: 1, IsPremium: true}

	go HandleRequest(shortProcess, &u0)
	go HandleRequest(shortProcess, &u0)
	go HandleRequest(shortProcess, &u0)
	go HandleRequest(shortProcess, &u0)
}