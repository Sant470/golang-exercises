// ----------------------
// Goroutine time slicing
// ----------------------

// How the Go's scheduler, even though it is a cooperating scheduler (not preemptive), it looks
// and feel preemptive because the runtime scheduler is making all the decisions for us. It is not
// coming for us.

// The program below will show us a context switch.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")

	// Create the first goroutine and manage its lifecycle here.
	go func() {
		printPrime("A")
		wg.Done()
	}()

	// Create the second goroutine and manage its lifecycle here.
	go func() {
		printPrime("B")
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.
// printPrime is not special. It just requires a little bit more time to complete.
// When we run the program, what we will see are context switches at some point for some particular
// prime number. We cannot predict when the context switch happen.
func printPrime(prefix string) {
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
