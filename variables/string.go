/*
 Strings are two word data structure in golang

var b string
It is going to be a two word data structure, A pointer and a backing array of bytes, both of which are empty at the moments.


b := "hello"
In this example the pointer is not nil and points to a backing array of bytes []byte{'h', 'e', 'l', 'l', 'o'}.
*/

package main

import (
  "fmt"
  "unsafe"
)

func main(){
  i:= "hello"
  j:= "strings demonstration"
  k:= "Browse our collection of free string samples, string loops, strings sample packs, violin sounds, orchestral loops and symphony samples."
  fmt.Printf("i %T:  %d\n", i, unsafe.Sizeof(i))
  fmt.Printf("j %T:  %d\n", j, unsafe.Sizeof(j))
  fmt.Printf("k %T:  %d\n", k, unsafe.Sizeof(k))
}

/*
The size should be 16 bytes, which is confirmed from the output.
i string:  16
j string:  16
k string:  16
*/
