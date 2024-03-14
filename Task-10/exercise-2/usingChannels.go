// Using Channels
// Original Program: How can we reduce the execution time of this program using concurrency?
package main

import (
	"fmt"
	"time"
)

func task(i int, ch chan<- bool) {
	// Task function simulates some work by printing a message and then sleeping for 100 milliseconds
	fmt.Println("Task", i)
	// send true on channel
	ch <- true
	time.Sleep(100 * time.Millisecond)
}

func main() {
	// start is the time where program start to execute
	start := time.Now()

	ch := make(chan bool)

	// Loop through 30 tasks sequentially
	for i := 1; i <= 30; i++ {
		// Launch goroutine
		go task(i, ch)
	}

	// Iterating 30 times to listen channel
	for i := 0; i < 30; i++ {
		<-ch
	}

	// Subtract the current time with started time and print it at end
	elapsed := time.Since(start)
	fmt.Printf("Time taken %s\n", elapsed)
}
