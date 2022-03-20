
// ------------------
// Drop Pattern
// ------------------

package main

import (
  "fmt"
  "time"
)

/*
  It allow us to drop tasks if we reach a certain threshold in terms of workload.
*/

func drop(){
  const cap = 5
  ch := make(chan string, cap)

  go func(){
    for p := range ch {
      fmt.Println("employee received signal: ",p)
    }
  }()

  const work = 20
  for w := 0; w < work; w++ {
    // Select allow single goroutine to handle multiple channel operations, whether they are send or receive at the same time.
    select {
      case ch <- "paper":
        fmt.Println("manager sent signal: ",w)
      default:
        fmt.Println("manager dropped data: ",w)
    }
  }
  close(ch)
  fmt.Println("manager sent shut down signal")
  time.Sleep(time.Second)
  fmt.Println("-------------------------------------------------")
}

func main(){
  drop()
}
