/*
Explicit conversion - The typed can be only assigned to typed one by converting it explicitily even though both have same memory profiling.
*/

package main

import (
  "fmt"
)

type user1 struct {
  name string
  addr string
}

type user2 struct {
  name string
  addr string
}


func main() {
  var u1 user1
  u2 := user2{name: "Sant", addr: "E1, Swarna Homes!!"}

  u1 = user1(u2)
  fmt.Printf("u1:  %#v", u1)
}
