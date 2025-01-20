package main

import (
	"fmt"
	"math"
)

// interfaces are collections of method signatures

// creating an object
type shape interface {
	area() float64
	perimeter() float64
	getName() string
}

// creating a struct for rectangle
type rectangle struct {
	width, height float64
	rectangleName string
}


// function for area of rectangle
func (r rectangle) area() float64 {
	return r.height * r.width
}

// function for perimeter of rectangle
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// function for name of rectangle
func (r rectangle) getName() string {
	return "Rectangle"
}


// creating a struct for circle
type circle struct {
	radius float64
}

// function for area of circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// function for perimeter of circle
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// function for name of circle
func (c circle) getName() string {
	return "Circle"
}


// creating a struct for hexagon
type hexagon struct {
	sideLength float64
}

// calculating the area of hexagon
func (hexagon hexagon) area() float64 {
	return 3/2 * math.Sqrt(3) * float64(hexagon.sideLength)
}

// calculating the perimenter of hexagon
func (hexagon hexagon) perimeter() float64 {
	return 6 * hexagon.sideLength
}

// function for name of hexagon
func (hexagon hexagon) getName() string {
	return "Hexagon"
}


// printing details of the shape
func printingShapeDetails(shape shape) {
	fmt.Println("Area: ",math.Round(shape.area()))
	fmt.Println("Perimeter: ",math.Round(shape.perimeter()))
	fmt.Println("Name: ", shape.getName())
}

func main() {

	// creating a new rectangle object
	rectangleObject := rectangle{width: 5, height: 10}
	printingShapeDetails(rectangleObject)

	// creating a new circle object
	circleObject := circle{radius: 10}
	printingShapeDetails(circleObject)

	// creating a new hexagon object
	hexagonObject := hexagon{sideLength: 13}
	printingShapeDetails(hexagonObject)

}




