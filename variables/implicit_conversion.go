/*
In example below, we are able to assign u1 value of type user1 to value of anonymous struct.
The implicit conversion allow us to assign anonymous type to typed type provided that both type have same data profiling
*/


package main

import (
  "fmt"
)

type user1 struct {
  name string
  addr string
}


func main() {
  var u1 user1
  u1 = struct {
    name string
    addr string
  }{
    name: "Sant",
    addr: "E01, Swarna Homes!",
  }

  fmt.Printf("u1: %#v\n", u1)
}
