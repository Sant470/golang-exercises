// ------------------
// Cancellation Pattern
// ------------------

package main

import (
	"fmt"
	"time"
)

/*
  This one illustrate a classical example of goroutine leak.
*/

func goroutineLeak() {

	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(500 * time.Millisecond))
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
	// The main goroutine will continue executing and we won't be listening on channel ch (line 30) and the other goroutine will be blocked on line 23
	// as this send can't get completed before receive.
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func main() {
	goroutineLeak()
}
