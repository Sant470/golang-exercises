package main

import (
	"fmt"
	"sync"
)

var scores = make(map[string]int)

func main() {
	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()

	// Wait for the write Goroutine to finish.
	wg.Wait()
	fmt.Println("Program Complete")
}
