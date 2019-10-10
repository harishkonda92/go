package interfacesI

import (
	"fmt"
)

type Describe interface {
	Describe()
}

type Person struct {
	name string
	age int
}
func (p Person)Describe () {
	fmt.Printf("name is %s age is %d \n", p.name , p.age)
}

func CompareType(i interface{})  {
	// V,T := i.(string)
	// fmt.Printf("type %T value %v", t,v)
	switch v:= i.(type) {
	case Describe:
		v.Describe()
	default:
		fmt.Println("undefined type", v)
	}
}

func Caller()  {
	CompareType("Harishh")

	a:= Person {
		name: "harish",
		age: 26,
	}
	CompareType(a)
}