package main

import "fmt"

func main() {

	// setting up a for loop
	// INITIAL; CONDITION, AFTER
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// a while loop is a for loop that only has a condition
	plantHeight := 1
	for plantHeight < 5 {
		fmt.Println("Still Growing! Current Height: ", plantHeight)
		plantHeight++
	}
	fmt.Println("Plant has grown to:", plantHeight, "inches")

	// continuing through a loop
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// stops current iteration of a loop and continues to the next one
			continue
		}
		fmt.Println(i)
	}

	// break out of a loop
	for i := 0; i < 10; i++ {
		if i == 5 {
			// stops current iteration of a loop and exits the loop
			break
		}
		fmt.Println(i)
	}

}
