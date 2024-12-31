package main

import "fmt"

func main() {

	// Define Wheel struct first
	type Wheel struct {
		Radius  int
		Material string
	}

	// car struct
	type car struct {
		Make       string
		Model      string
		Height     int
		Width      int
		FrontWheel Wheel
		BackWheel  Wheel
	}

	// Accessing fields of a struct
	myCar := car{
		FrontWheel: Wheel{
			Radius:  5,
			Material: "rubber",
		},
		BackWheel: Wheel{
			Radius:  5,
			Material: "rubber",
		},
	}

	fmt.Println("Front Wheel Radius: ", myCar.FrontWheel.Radius)

	// Making and using anonymous structs
	myBus := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}

	fmt.Println("Bus Make: ", myBus.Make)
	fmt.Println("Bus Model: ", myBus.Model)

	// Embedded structs
	type cycle struct {
		Make  string
		Model string
	}

	type truck struct {
		cycle
		BedSize int
	}

	lanesTruck := truck{
		BedSize: 10,
		cycle: cycle{
			Make:  "toyota",
			Model: "camry",
		},
	}

	fmt.Println("Bed Size: ", lanesTruck.BedSize)
	fmt.Println("Make: ", lanesTruck.Make)
	fmt.Println("Model: ", lanesTruck.Model)


	r := rect {
		width: 5,
		height: 10,
	}

	fmt.Println("Area is: " , r.area())

}

// struct methods
type rect struct {
	width int
	height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
	return r.width * r.height
}