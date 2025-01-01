package main

import(
	"fmt"
	"log"
)


type Copier interface {

	// methods with arguments and return data
	Copy(sourceFile string, destinationFile string) (bytesCopied int)

}

// shape interface
type newShape interface {
	area() float64
}

// ball struct
type ball struct {
	radius float64
}

// the area() method
func(c ball) area() float64 {
	return 3.14 * c.radius * c.radius
}


func main() {

	var shapeObject newShape = ball{radius: 5}

	// type assertion
	ballObject, ok := shapeObject.(ball)
	if !ok {
		log.Fatal("shapeObject is not a circle")
	}

	fmt.Println("Radius: ", ballObject.radius)
	fmt.Println("Area: ", ballObject.area())

}

