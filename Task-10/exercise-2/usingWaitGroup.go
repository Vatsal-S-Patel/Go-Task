// Using WaitGroup

// Original Program: How can we reduce the execution time of this program using concurrency?
package main

import (
	"fmt"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup, i int) {
	// Task function simulates some work by printing a message and then sleeping for 100 milliseconds
	fmt.Println("Task", i)
	time.Sleep(100 * time.Millisecond)

	// Reduce the WaitGroup counter by 1
	wg.Done()
}

func main() {
	// start is the time where program start to execute
	start := time.Now()

	// Declare a WaitGroup
	var wg sync.WaitGroup

	// Loop through 30 tasks sequentially
	for i := 1; i <= 30; i++ {
		// Adding 1 to WaitGroup counter
		wg.Add(1)
		// Launch goroutine
		go task(&wg, i)
	}

	// Wait till the WaitGroup counter become 0, means till all goroutines are executed
	wg.Wait()

	// Subtract the current time with started time and print it at end
	elapsed := time.Since(start)
	fmt.Printf("Time taken %s\n", elapsed)
}
