package main

import "fmt"

func main() {

	var height int = 6
	fmt.Println("Height is: ", height)

	// simple if statement
	if height > 4 {
		fmt.Println("You are tall enough")
	}

	// else if and else
	if height > 6 {
		fmt.Println("You are super tall!")
	} else if height > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are not tall enough!")
	}

	// if statement and initial condition
	if number := 5; number > 0 {
		fmt.Println("The number is positive: ", number)
	}


}