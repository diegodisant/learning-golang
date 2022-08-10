package main

import (
	"fmt"
	"math"
)

type Circle struct {
	X      float64
	Y      float64
	Radius float64
}

func circleArea(circle Circle) float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func main() {
	// declare the circle as a pointer using new
	var circleOne *Circle = new(Circle)

	circleOne.X = 1
	circleOne.Y = 1
	circleOne.Radius = 10

	fmt.Println(circleOne)
	fmt.Println(*circleOne)
	fmt.Println(circleArea(*circleOne))

	// declare the circle as normal
	var circleTwo Circle

	circleTwo.X = 9
	circleTwo.Y = 10
	circleTwo.Radius = 5

	fmt.Println(circleTwo)
	fmt.Println(circleArea(circleTwo))

	// declare the circle as normal with default values
	var circleThree Circle = Circle{
		X:      5,
		Y:      5,
		Radius: 12,
	}

	fmt.Println(circleThree)
	fmt.Println(circleArea(circleThree))
}
