
// ------------------
// Wait for finished pattern
// ------------------

package main

import (
  "fmt"
  "time"
  "math/rand"
)

/*
Think about being a manager and hiring a new employee. In this scenario you want your employee to perform a task immediately
and you need to wait for the result of their work, you need to wait b'cause you can't move until you know that they are done
and you also don't need anything from them.
*/

func waitForFinished(){
  ch := make(chan struct{})

  go func(){
    time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
    close(ch)
    fmt.Println("employee: sent siganl")
  }()

  _,wd := <-ch
  fmt.Println("manager: received signal: ", wd)
  time.Sleep(time.Second)
  fmt.Println("-------------------------------------------------")
}

func main(){
  waitForFinished()
}
