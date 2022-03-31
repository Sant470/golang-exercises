// ------------------
// Slice
// ------------------


package main

import "fmt"

func main() {
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for _, friend := range friends {
    friends = friends[:2]
    fmt.Println("friend: ", friend)
	}
  fmt.Println("after for loop: ",friends)
}

/*
friend:  Annie
friend:  Betty
friend:  Charley
friend:  Doug
friend:  Edward
after for loop:  [Annie Betty]

This is b'cause we are making a local copy of the friends and when we try to access any index greator than 2 after it get modified it refer to the local copy.
*/
