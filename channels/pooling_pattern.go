// ------------------
// Pooling pattern
// ------------------

package main

import (
	"fmt"
	"time"
)

/*
The idea here is to illustrate pooling machenism, a pool of threads reading message from a common channel and acting on it.
The message at any point of time will be delivered to only one thread or each thread will be working on an unique message.
*/

func pooling() {
	ch := make(chan int)
	const empls = 2
	for e := 0; e < empls; e++ {
		go func(emp int) {
			// channel receive (two threads blocked on same channel)
			for p := range ch {
				fmt.Printf("employee %d: received signal: %d\n", emp, p)
			}
			// signal change ternminate the for loop(basically close will terminate the for loop)
			fmt.Printf("employee %d: received shut down signal\n", emp)
		}(e)
	}
	const work = 10
	for w := 0; w < work; w++ {
		ch <- w
		fmt.Println("manager sent signal: ", w)
	}
	close(ch)
	fmt.Println("manager shut down signal")
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

func main() {
	pooling()
}
