// ------------------
// Cancellation Pattern
// ------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
  This pattern take care of the scenario that a task can't take forever and it need to be stopped through another goroutine if it's doesn't get completed before the deadline.
*/

func cancellation() {
	//  Delayed Guarantee
	ch := make(chan string, 1)

	go func() {
		// it'll take anywhere 0 to 500 ms to complete
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("emplyee: sent signal")
	}()

	// this would be came back with the result after 100 ms
	tc := time.After(100 * time.Millisecond)
	select {
	case p := <-ch:
		fmt.Println("manager: recv'd signal: ", p)
	case t := <-tc:
		fmt.Println("manager: timeout: ", t)
	}
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func main() {
	cancellation()
}
