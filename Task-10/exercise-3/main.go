package main

import (
	"fmt"
	"time"
)

// fizz function checks that number passed in fizzChan is divisible of 3 or not,
// if yes, it will send true on outputChan otherwise false
func fizz(fizzChan chan int, outputChan chan bool) {
	for {
		if <-fizzChan%3 == 0 {
			outputChan <- true
		} else {
			outputChan <- false
		}
	}
}

// buzz function checks that number passed in buzzChan is divisible of 5 or not,
// if yes, it will send true on outputChan otherwise false
func buzz(buzzChan chan int, outputChan chan bool) {
	for {
		if <-buzzChan%5 == 0 {
			outputChan <- true
		} else {
			outputChan <- false
		}
	}
}

// fizzBuzz function checks that number passed in buzzChan is divisible of 3 and 5 or not,
// if yes, it will send true on outputChan otherwise false
func fizzBuzz(fizzBuzzChan chan int, outputChan chan bool) {
	for {
		if <-fizzBuzzChan%15 == 0 {
			outputChan <- true
		} else {
			outputChan <- false
		}
	}
}

func main() {
	// now is a time when execution is started
	now := time.Now()

	// Declaration of channels
	var (
		fizzChan     = make(chan int) // channel for printing fizz, when number is multiple of 3
		buzzChan     = make(chan int) // channel for printing fizz, when number is multiple of 5
		fizzBuzzChan = make(chan int) // channel for printing fizz, when number is multiple of 3 and 5 both

		outputChan = make(chan bool) // channel for tracking that is number multiple of 3 or 5 or 3 and 5 both
	)

	// Invoked all 3 goroutines
	go fizz(fizzChan, outputChan)
	go buzz(buzzChan, outputChan)
	go fizzBuzz(fizzBuzzChan, outputChan)

	// Iterate over 1 to 30 number
	for i := 1; i <= 30; i++ {
		// giving input to fizzBuzzChan first, so that it can check that number is multiple of 3 and 5 both or not
		// if yes, then outputChan will recieve true and FizzBuzz will print
		fizzBuzzChan <- i
		if <-outputChan {
			fmt.Println(i, "FizzBuzz")
			continue
		}
		// giving input to fizzChan first, so that it can check that number is multiple of 3 or not
		// if yes, then outputChan will recieve true and Fizz will print
		fizzChan <- i
		if <-outputChan {
			fmt.Println(i, "Fizz")
			continue
		}
		// giving input to buzzChan first, so that it can check that number is multiple of 5 or not
		// if yes, then outputChan will recieve true and Buzz will print
		buzzChan <- i
		if <-outputChan {
			fmt.Println(i, "Buzz")
			continue
		}
		// If number is neither multiple of 3 nor 5 nor of both, it will print
		// fmt.Println(i)
	}

	// Prints execution time
	fmt.Println("Execution Time:", time.Since(now))
}

// -----------------------------------------
// func main() {
// 	now := time.Now()
// 	for i := 1; i <= 30; i++ {
// 		if i%15 == 0 {
// 			fmt.Println("1")
// 		} else if i%3 == 0 {
// 			fmt.Println("2")
// 		} else if i%5 == 0 {
// 			fmt.Println("3")
// 		} else {
// 			log.Println(i)
// 		}
// 	}
// 	fmt.Println(time.Since(now))
// }

// func main() {
// 	now := time.Now()
// 	ch := make(chan int)

// 	for i := 1; i <= 30; i++ {
// 		if i%15 == 0 {
// 			go multipleOfThreeAndFive(ch)
// 			<-ch
// 		} else if i%3 == 0 {
// 			go multipleOfThree(ch)
// 			<-ch
// 		} else if i%5 == 0 {
// 			go multipleOfFive(ch)
// 			<-ch
// 		} else {
// 			log.Println(i)
// 		}
// 	}
// 	fmt.Println(time.Since(now))
// }

// func multipleOfThree(ch chan int) {
// 	log.Println("Fizz")
// 	ch <- 1
// }

// func multipleOfFive(ch chan int) {
// 	log.Println("Buzz")
// 	ch <- 1
// }

// func multipleOfThreeAndFive(ch chan int) {
// 	log.Println("FizzBuzz")
// 	ch <- 1
// }
