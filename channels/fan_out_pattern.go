
// ------------------
// Fan Out pattern
// ------------------

package main

import (
  "fmt"
  "time"
  "math/rand"
)

/*
  Fan out basically allow us to take a piece of work and distribute it across n number of Gorouines that can run in parallel.
*/

func fanOut(){
  empls := 20
  ch := make(chan string, empls)
  for e:=0; e < empls; e++ {
    go func(emp int) {
      time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
      ch <- "paper"
      fmt.Println("emplyee sent signal: ")
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

func main(){
  fanOut()
}
