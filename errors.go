package main

import(
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by 0")
	}
	return a/b, nil
}

func main() {
	resultOne, errOne := divide(10, 2)
	if errOne != nil {
		fmt.Println("Error: ", errOne)
	} else {
		fmt.Println("Result: ", resultOne)
	}

	resultTwo, errTwo := divide(10, 0)
	if errTwo != nil {
		fmt.Println("Error: ", errTwo)
	} else {
		fmt.Println("Result: ", resultTwo)
	}
}