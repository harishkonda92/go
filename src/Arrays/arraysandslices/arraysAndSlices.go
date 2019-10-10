package arraysandslices

import (
	"fmt"
)
func Dummy()  {
	fmt.Println("im here!")
	a  :=[]int{1,12,4,55,6}
	fmt.Println(a)
	aslice := a[2:5]
	fmt.Println("before", a)
	for i := range aslice {
		aslice[i]++
	}
	fmt.Println("after", a)
	//  b :=[3]string{"harish", "konda", "vinnu"}
	//  for ind, vari := range b {
	// 	 fmt.Println(ind, vari)
	//  }
	// fmt.Println(aslice)
	fmt.Println(aslice[:cap(aslice)], 1)
	fmt.Printf("alsice length %d , capacity %d \n ", len(aslice), cap(aslice));

	// make slice of an arry using make() function
	i:= make([]string, 4, 6)
	fmt.Println(i)
	
	i = append(i, "1", "2", "3", "4", "5", "6", "7", "8");
	fmt.Printf("capacitu %d, length of i %d", cap(i), len(i))
	aandslice := append(aslice, a...)
	fmt.Println(aandslice)	 
	fruits := []string{"apple", "banana", "mango"}
	vegitables := []string{"tomato", "brinjal", "ladiesfinger"}
	food := append(fruits, vegitables...)
	fmt.Println(food)
}

func MoreAbooutSlices(){
	// num := [3]int{1, 3, 5}
	// num1 := num[:]
	// fmt.Println("before slice call", num, num1)
	// SubtractOne(num1)
	// fmt.Println("after slice call", num, num1)

	countriesneeded := countries()
	fmt.Println(countriesneeded)
}
func SubtractOne(numbers[]int) {
	for i := range numbers {
		numbers[i] -=2;
	}
}

func countries () []string {
	// defining a slice
	asia := []string {"india", "china", "japan", "honkong", "pakistan"}
	// slicing original array slice to lenth 3
	neededCoountries := asia[:len(asia)-2]
	// defining a copy slice of original array slice
	countriesCopy := make([]string, len(neededCoountries) )
	fmt.Println(cap(countriesCopy), 58)
	// coping the sliced array to copy array slice
	copy(countriesCopy, neededCoountries)
	fmt.Println(asia, 60)
	return (countriesCopy)

}