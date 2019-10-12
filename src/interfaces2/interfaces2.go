// pointer receivers vs value receivers
package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

// value reiver
func (p Person) Describe() {
	fmt.Printf("%s is %d years old \n", p.name, p.age)
}

// address struct
type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("state : %s , country : %s", a.state, a.country)
}

func main() {
	var d1 Describer
	person := Person{
		age:  26,
		name: "Harish",
	}
	person2 := Person{
		age:  30,
		name: "sai kiran",
	}
	d1 = person
	d1.Describe()
	d1 = &person2
	d1.Describe()

	var d2 Describer

	a := Address{state: "TG", country: "Ind"}
	// d2 = a
	d2 = &a
	d2.Describe()

}
