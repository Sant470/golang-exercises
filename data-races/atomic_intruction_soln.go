// ----------------
// Atomic Functions
// ----------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// counter is a variable incremented by all Goroutines.
// Notice that it's not just an int but int64. We are being very specific about the precision
// because the atomic function requires us to do so.
var counter int64

func main() {
	// Number of Goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				// Safely add one to counter.
				// We can run this program all day long and still get 4 every time.
				atomic.AddInt64(&counter, 1)
			}

			wg.Done()
		}()
	}

	// Wait for the Goroutines to finish.
	wg.Wait()

	// Display the final value.
	fmt.Println("Final Counter:", counter)
}
