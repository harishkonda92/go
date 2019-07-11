package main

import (
	"fmt"
	"geometry/rectangle" // this is how we import a custom package
	"log"
)

var rectLen, rectWidth float64 = 5.2, 3.2

func init() {
	fmt.Println("geometry inititialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than 0")
	}

}
func main() {
	/*
		if else blocks
	*/
	// if rectLen/2 == 0 {
	// 	fmt.Println("length is even")
	// } else {
	// 	fmt.Println("length is odd")
	// }

	/*
		for loop
	*/
	// for i := 0; i < 10; i++ {
	// 	if i != 5 {
	// 		// break
	// 		continue
	// 	}
	// 	fmt.Print(i)
	// }

	// n := 5
	// for i := 0; i < n; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 		if i == j {
	// 			break
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	/*
		labels in for loop
	*/
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d \n", i, j)
			if i == j {
				break outer
			}
		}
	}

	fmt.Printf("area of a rectangle %.2f", rectangle.Area(rectLen, rectWidth))
}
