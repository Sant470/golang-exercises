// ------------------
// Wait for result pattern
// ------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Think about being a manager and hiring a new employee.
In this scenario you want your new employee to hand you results before you can continue.
*/

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee: sent siganl")
	}()

	p := <-ch
	fmt.Println("manager: received signal ", p)
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func main() {
	waitForResult()
}
