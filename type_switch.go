package main

import "fmt"

// type switch makes it easy to do serveral type assertions

func printNumericValue(num interface{}) {

	switch v := num.(type) {
	case int:
		// prints the type of the variable
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}

}

func main() {
	printNumericValue(1)
	printNumericValue("Hello")
	printNumericValue(struct{}{})
}
