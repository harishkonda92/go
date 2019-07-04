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
	fmt.Printf("\n type of a is %T, size of a is %d", a, unsafe.Sizeof(a))

	type myBool bool
	const someBool myBool = true
	fmt.Println(someBool)
	println("--------------------")

	const ab = 4
	var intVar = ab
	var int32Var int32 = ab
	var float64Var float64 = ab
	var complex64Var complex64 = ab
	fmt.Printf("\n type of intVar %T, type of int32Var %T, type of float64Var %T, type of complex64 %T", intVar, int32Var, float64Var, complex64Var)
	println("\n --------------------")
	// var complexVarr = intVar + 6i
	// fmt.Printf("type of complexxVar %T", complexVarr)

	// functions
	fmt.Println(calculateOf(10, 5))
	println("\n --------------------")
	// area, perimeter := rectProperties(8, 9)

	// fmt.Printf("\n area %f and perimeter %f", area, perimeter)
	println("\n --------------------")
	fmt.Println(rectProperties(8, 9))

	println("\n --------------------")
	fmt.Println(namedRectProperties(8, 9))

	area, _ := namedRectProperties(8, 9) // second returned value will be discarded
	fmt.Printf("area %f", area)
}

// function with single return value
func calculateOf(price, num int) int {
	var totalPrice = price * num
	return totalPrice
}

// function with multiple return values
func rectProperties(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// function with named return values

func namedRectProperties(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}
