package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
}

func (p person) fullName() {
	fmt.Print(p.firstname, " ", p.lastname)
}

func main() {
	// p := person{
	// 	firstname: "harish",
	// 	lastname:  "konda",
	// }
	// defer p.fullName()
	// fmt.Print("welcome ")

	// a := 5
	// defer printA(a)
	// a = 10
	// fmt.Println("value of a before deffered function call ", a)

	name := "Harish"
	fmt.Printf("Original string : %s\n ", string(name))
	fmt.Printf("reversed string: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

}

func printA(a int) {
	fmt.Println("value of a in deffered function ", a)
}
