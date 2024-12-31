package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("Hi there!")

	// passing in values by variable
	z := 5
	increment(z)

	fmt.Println("Using Increment: ", z)

	// ignoring a return value
	x, _ := getPoint()
	fmt.Println("Extracting x: ", x)

	valueOne, valueTwo := getCoords()
	fmt.Println("Value One & Two: ", valueOne, valueTwo)

	valueThree, valueFour := getNewCoords()
	fmt.Println("Value Three & Four: ", valueThree, valueFour)

	result, error := divide(10, 9)
	fmt.Println("Result & Error: ", result, error)

}

// functions in go can take one or more arguments
func sub(x int, y int) int {
	return x - y
}

// multiple parameters of the same type
func add(x, y int) int {
	return x + y
}

func increment(x int) {
	x++
}

func getPoint() (x int, y int) {
	return 3, 4
}

// named return values
func getCoords() (x, y int) {
	return
}

// explicit returns
func getNewCoords() (x, y int) {
	return x, y
}

// early returns
// first two are parameters
// last two are returns
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		// the early return
		return 0, errors.New("Can't divide by zero")
	}
	return dividend/divisor, nil
}



