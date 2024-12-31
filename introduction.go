
// lets the compiler to know that we want to compile
package main

// imports the formatting package
import "fmt"

// defines the main function
func main() {

	fmt.Println("hello world")

	// declare a variable of int data type
	var number int = 9 
	fmt.Println("Number: " , number)

	// declaring a pi as variable
	var pi float64 = 3.14159
	fmt.Println("Value of pi:" , pi)

	// the basic variable types

	// bool string int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte 
	// rune
	// float32 float64
	// complex64 complex128

	// short variable declaration
	var empty string // is same as empty := ""

	numCars := 10
	temperature := 0.0
	var isFunny = true

	fmt.Println("Printing out Short Variables: ", empty, numCars, temperature, isFunny)

	// type inference
	// declaring a variable without specifying a explicit type

	var i int 
	j := i

	fmt.Println("Value for j: ", j)

	// same line variable declaration
	mileage, company := 80276, "Tesla"
	fmt.Println("Mileage & Company: ", mileage, company)

	// converting the types
	temperatureInt := 88
	temperatureFloat := float64 (temperatureInt)
	fmt.Println("Floating Point: ", temperatureFloat)

	// constants variables
	const myInt = 15

	// constants can be computed
	const firstName = "Lane"
	const lastName = "Wagner"
	const fullName = firstName + " " + lastName
	fmt.Println("Full Name: " , fullName)

	// formatting strings

	// using %v to interpolate the default representation
	stringOne := fmt.Sprintf("I am %v years old", 10)
	stringTwo := fmt.Sprintf("I am %v years old" , "way too much")
	fmt.Println("String One: ", stringOne,  "\n" , "String Two: ", stringTwo)


	// interpolate a string using %s
	integerOne := fmt.Sprintf("I am %s years old", "way too much")

	// interpolate a integer in decimal form using %d
	decimalOne := fmt.Sprintf("I am %d years old", 10)

	// interpolate a decimal using %f
	decimalTwo := fmt.Sprintf("I am %f years old", 10.523)
	decimalThree := fmt.Sprintf("I am %.2f years old", 10.523)

	fmt.Println("Integer One: ", integerOne , "\n")
	fmt.Println("Decimal One: ", decimalOne , "\n")
	fmt.Println("Decimal Two: ", decimalTwo, "\n")
	fmt.Println("Decimal Three: " , decimalThree, "\n")







}
