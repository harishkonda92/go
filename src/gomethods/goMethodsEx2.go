package gomethods

 import (
	 "fmt"
 )

//  receivers in methods vs value arguments in function 

type SRectangle struct {
	length int
	width int
}

func (va SRectangle) area()  {
	fmt.Printf("area of a rectangle of length %d and width %d is  %d \n",va.length, va.width, va.length * va.width)
}

func area(r SRectangle)  {
	fmt.Printf("area1 of a rectangle of length %d and width %d is  %d \n",r.length, r.width, r.length * r.width)
}
func GetsArea()  {
	 r := SRectangle {
		 length: 3,
		 width: 2,
	 }
	//  area(r)
	//  r.area()

	//  defining a new pointer of rectangle "r"
	p := r
	// p.area()  // this works for receiver method 
	area(p)
 }