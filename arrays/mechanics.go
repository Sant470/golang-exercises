/*
The cleanest and easiest way to create the predictable access pattern is to use an array,
An array give us the ability to create a contigous block of memory.
We also define it based on an element size(int, strings, ...) that are our predictable stride.
Every index is a predictable stride from the other, making the prefetcher to bring in those cache lines way ahead of when you need them.
*/


package main
import "fmt"

func main() {
  fruits := [5]string {"Apple", "Banana", "Orange", "Grape", "Plum"}
  for i, _ := range fruits {
    fmt.Println("index: ", i, "address: ", &fruits[i])
  }
}


/*
index:  0 address:  0xc000100050
index:  1 address:  0xc000100060
index:  2 address:  0xc000100070
index:  3 address:  0xc000100080
index:  4 address:  0xc000100090

The two things are visible from the output

1. Each of the elements are next to each other, they are on contigous block of memory.
2. A string is two word data structure, which means we need 16 bytes of memory for each string (which is out predictable stride) on AMD64 bit architecture.
*/
