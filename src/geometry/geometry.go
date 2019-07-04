package main

import (
	"fmt"
	"geometry/rectangle" // this is how we import a custom package
)

func main() {
	var rectLen, rectWidth float64 = 5.2, 3.2
	fmt.Println("Geometrical shapes and properties")

	fmt.Printf("area of a rectangle %.2f", rectangle.Area(rectLen, rectWidth))
}
