// ------------------
// Wait for task pattern
// ------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Think about being a manager and hiring a new employee.
In this scenario you want your new employee to perform a task but they need to wait unit you are ready.
This is b'cause you need to hand them a piece of paper before they start.
*/

func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("employee: received signal : ", p)
	}()

	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	ch <- "paper"
	fmt.Println("manager: sent signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func main() {
	waitForTask()
}
