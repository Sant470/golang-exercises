package main

import "fmt"

func main() {
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for i := range five {
		five[1] = "Jack"

		if i == 1 {
			fmt.Printf("after[%s]\n", five[i])
		}
	}
	fmt.Println("after loop", five[1])
}
