package main

// import "fmt"

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var age = 20
	fmt.Println("hy age is ", age)
	age = 29
	fmt.Println("hy age is ", age)
	age = 26
	fmt.Println("harish age is ", age)
	var fname, lname string = "harish", "konda"
	var (
		numr1 = 1
		numr2 = 2
	)

	// name1, agee := "harishh", 26

	name1, agee := "konda", "26"

	fmt.Println("firstname", fname, "lastname", lname)
	fmt.Println("number1", numr1, "number2", numr2)
	fmt.Println("name1", name1, "age", agee)

	a, b := 145.5, 533.5
	c := math.Max(a, b)
	fmt.Println("minimum value is ", c)

	// types
	// fmt.Printf("\n %T is the format specifier", a)
	println("-------------------")
	fmt.Printf("\n type of a is %T, size of a is %d", a, unsafe.Sizeof(a));



}
