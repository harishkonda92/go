package gomethods

import (
	"fmt"
)


//  receivers in methods vs value arguments in function 

type Rectangle struct {
	length int
	width int
}

func (va *Rectangle) area1()  {
	fmt.Printf("2 area of a rectangle of length %d and width %d is  %d \n",va.length, va.width, va.length * va.width)
}

func area1(r *Rectangle)  {
	fmt.Printf("2 area1 of a rectangle of length %d and width %d is  %d \n",r.length, r.width, r.length * r.width)
}
func GetsArea1()  {
	r := Rectangle {
		length: 3,
		width: 2,
	}

	p := &r
	p.area1()
	// area1(r) // this doesnt work as the function expects a pointer
}