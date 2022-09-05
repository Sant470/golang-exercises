/*
 Variables are at the heart of of a langugage and provide the ability to read and write from memory.
 In Go, access to memory is type safe and it means that the compiler takes type seriously, and It
 Won't allow us to use it outside the scope of how they are defind.

A variable provide two pieces of information, 1) size 2) representation.

Word - A value that can change it's size depending on the architecture.

If your address and your word size is 64 bits or 8 bytes, lets just make the integer follow. In short, an integer is alaways
gonna be 1 word size.
*/


package main

import(
  "fmt"
  "unsafe"
)

func main(){
  i:= 9
  j:= 9999
  k:= 999999
  l:= 99999999999
  fmt.Printf("i %T:  %d\n", i, unsafe.Sizeof(i))
  fmt.Printf("j %T:  %d\n", j, unsafe.Sizeof(j))
  fmt.Printf("k %T:  %d\n", k, unsafe.Sizeof(k))
  fmt.Printf("l %T:  %d\n", l, unsafe.Sizeof(l))
}
