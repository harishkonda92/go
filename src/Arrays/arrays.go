package main

import (
	"fmt"
)

func main() {
	var a [5]int // array with length 3, initial values are given 0 by default
	a[0] = 10
	a[1] = 10
	a[2] = 12

	b := [3]int{12, 13} //  intializing array with values , its not necessary to initialize array
	fmt.Println(b)

	c := [...]int{13, 15}
	fmt.Println(c)

	countries := [...]string{"india", "usa", "china", "japan"}
	_ = countries
	// fmt.Println(duplicateCountries[3])

	// changeLocal(a)
	fmt.Println("afeter passing to a function", a, len(a))

	floatingArr := [...]float64{1, 2.3, 4.5, 4, 7, 7, 7.8}

	// for index := 0; index < len(floatingArr); index++ {
	// 	fmt.Println(floatingArr[index])
	// }

	for i, v := range floatingArr {
		fmt.Printf(" i = %d ,v = %.2f", i, v)
	}
}

func changeLocal(num [5]int) {
	num[4] = 66
	fmt.Println("inside changeLocal", num)
}
