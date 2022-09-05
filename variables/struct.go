
/*
This examples demonstrate the padding left while allocating memory for a value of struct,
In the below examples it seems that a value of type example will just needs 7 bytes of memory, but actually it needs 8 bytes.

The Alignment (it comes from the hardware, it basically state that memory should be accessed based on word size, for amd64 bit processor it's 8 bytes) does not allow a variable to sit across two different blocks of memory, as it would require two read/writes to access them.
So, the initialization begin from the multiple of the size of the feild.

1. bool requires one byte, so it's get the 0th location of an 8 byte memory block.
2. int16 requires 2 bytes, so it's intialised from 2nd position, leaving 1st position empty (or leaving a padding of 1 bytes).
3. float32 requires 4 bytes of memory, so it's initialised from 4th to 7th position.

Values of example required 8 bytes of memory out of which 1 byte of padding is there.
*/

package main

import (
  "fmt"
  "unsafe"
)

type example struct {
  flag bool
  age int16
  pi float32
}

func main(){
  var e1 example
  e2 := example{true, 16, 3.14}
  fmt.Printf("e1 size  %d\n", unsafe.Sizeof(e1))
  fmt.Printf("e2 size %d\n", unsafe.Sizeof(e2))
}

/*
The size should be 8 bytes which is confirmed by the output
e1 size  8
e2 size 8
*/
