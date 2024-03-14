package main

import (
	"fmt"
	"time"
)

func main() {

	// Two channels of boolean
	var ch1 = make(chan bool)
	var ch2 = make(chan bool)

	// Launch two goroutines
	go a(ch1, ch2)
	go b(ch1, ch2)

	// Sending true on
	ch1 <- true

	time.Sleep(10 * time.Microsecond)
}

func a(ch1, ch2 chan bool) {
	for {
		// Reciece from ch1
		<-ch1
		fmt.Println("0")
		// Send true on ch2
		ch2 <- true
	}
}

func b(ch1, ch2 chan bool) {
	for {
		// Reciece from ch2
		<-ch2
		fmt.Println("1")
		// Send true on ch1
		ch1 <- true
	}
}
