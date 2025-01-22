package main

import "fmt"

func main() {

	// valid key types
	stringMap := map[string]int{"key": 42}
	fmt.Println("String Key: ", stringMap["key"])

	intMap := map[int]string{1: "one", 2: "two"}
	fmt.Println("Int Key: ", intMap[2])

	structMap := map[struct{ X, Y int }]string{
		{1, 2}: "point1",
		{3, 4}: "point2",
	}

	fmt.Println("Struct key: ", structMap[struct{ X, Y int }{1, 2}])

}
