package main

import "fmt"

func main() {
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for i, v := range five {
		five[1] = "Jack"
		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}
	fmt.Println("after loop", five[1])
}
