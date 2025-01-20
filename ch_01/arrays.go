package main

import "fmt"

// variadic functions
func concat(strs ...string) string {
	final := ""
	for _, str := range strs {
		final += str
	}
	return final
}

// function for printing strings
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {

	fmt.Println("Start of the Program!")

	// array of 10 integers
	var myInts [10]int

	// delcare initiliazed literal
	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println("Array of 10 Integers: ", myInts)
	fmt.Println("Initialized Literal: ", primes)

	// slicing of lists
	mySlice := primes[1:4]

	fmt.Println("Sliced: ", mySlice)

	// low index is inclusive and high index is exclusive

	// creating new slices
	// func make([]T, len, cap) []T
	myNewSlice := make([]int, 5, 10)
	fmt.Println("Making a New Slice: ", myNewSlice)

	// capacity is usually omitted
	myNewerSlice := make([]int, 5)
	fmt.Println("Making a Newer Slice: ", myNewerSlice)

	// making a slice with specific values
	mySpecificSlice := []string{"I", "love", "go"}
	fmt.Println("Love Go:", mySpecificSlice)
	fmt.Println("Length: ", len(mySlice))

	// using the variadic function
	final := concat("Hello", " there", " friend")
	fmt.Println(final)
	// output: hello there friend

	// println vs Sprintf
	fmt.Println("Salvin", "Is", "Here")

	// sprintf formats strings based on a specifier
	name := "Salvin"
	age := 21
	pi := 3.14159

	formatted := fmt.Sprintf("Hello, %s! You are %d years old. Pi is %.2f.", name, age, pi)
	fmt.Println(formatted)

	// the spread operator
	// allows to pass a slice into the variadic function
	names := []string{"bob", "sue", "alice"}
	printStrings(names...)

	// appending a slice
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("After appending a single element:", slice)

	// ranging over a slice
	fruits := []string{"apple", "banana", "grape"}

	// for looping through it
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

}
