package main

import (
	"fmt"
)

func main() {

	// using the make function
	ages := make(map[string]int)
	ages["Braden"] = 20
	ages["Salvin"] = 21
	ages["Mason"] = 19

	// will overwrite
	ages["Mason"] = 21

	fmt.Println(ages)

	// using a literal to make a map
	ages = map[string]int{
		"John": 37,
		"Mary": 21,
	}

	fmt.Println(ages)

	// len() returns total number of key / value pairs
	fmt.Println(len(ages))

	// grade map
	grades := make(map[string]int)
	grades["Alice"] = 85
	grades["Bob"] = 90
	grades["Charlie"] = 78

	// printing out the map
	fmt.Println("Initial Map:", grades)

	// getting an element
	aliceGrades := grades["Alice"]
	fmt.Println("Alice's Grade:  ", aliceGrades)

	// checking if a key exists
	grade, ok := grades["David"]
	if ok {
		fmt.Println("David's grade: ", grade)
	} else {
		fmt.Println("David's grade not found.")
	}

	// deleting an element
	delete(grades, "Charlie")
	fmt.Println("Map after deleting Charlie:", grades)

	// retreiving a deleted key
	_, ok = grades["Charlie"]
	if ok {
		fmt.Println("Charlie's grade is still in the map")
	} else {
		fmt.Println("Charlie's grade has been removed")
	}

	// iterating over the map
	fmt.Println("Remaining students and their grades")
	for student, grade := range grades {
		fmt.Printf("%s: %d\n", student, grade)
	}

}
