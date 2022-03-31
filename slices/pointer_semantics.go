// ------------------
// Slice
// ------------------

/*
slice is a three word data structure
1. Pointer to a backing array
2. Length
3.Capacity

It's reference type along with 'map', 'channels', 'interface values' and 'functions' for the reason:
1. It has a pointer
2. When any of these are set to zero value, they are considererd to be nil.

*/

package main

import "fmt"

func main() {
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i := range friends {
		friends = friends[0:2]
		fmt.Println("friend: ", friends[i])
	}
}

/*
friend:  Annie
friend:  Betty
panic: runtime error: index out of range [2] with length 2

This is b'cause we are not making a local copy of the friends and when we try to access any index greator than 2 after it get modified,
It's panicing.
*/
