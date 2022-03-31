// ------------------
// Cancellation Pattern using context
// ------------------

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string, 1)

	// this would be came back with the result after 50 ms
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	go func() {
		// it'll take anywhere 0 to 500 ms to complete
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("sent signal")
	}()

	select {
	case p := <-ch:
		fmt.Println("work completed: ", p)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}
}
