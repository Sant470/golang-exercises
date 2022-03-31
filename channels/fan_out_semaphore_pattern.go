// ------------------
// Fan Out Semaphore Pattern
// ------------------

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
  Fan Out Semaphore is basically fan out operation with restricted number of parallelisation
*/

func fanOutSem() {
	empls := 20
	ch := make(chan string, empls)
	const cap = 5
	sem := make(chan bool, 5)
	for e := 0; e < empls; e++ {
		go func(emp int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("emplyee sent signal: ", emp)
			}
			<-sem
		}(e)
	}

	for empls > 0 {
		p := <-ch
		fmt.Println("message: ", p)
		fmt.Println("manager: received signal ", empls)
		empls--
	}
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func main() {
	fanOutSem()
}
