package gomethods

import (
	"fmt"
)

type myInt int
func (a myInt) add(b myInt) myInt   {
	return a + b
}

func AddsTwoNum()  {
	// fmt.Prinln("hi")
	num1 := myInt(1)
	num2 := myInt(3)
	 
	sum := num1.add(num2)
	fmt.Println(sum)
}