package interfacesI

import (
	"fmt"
)

func test(i interface{})  { // passing an empty interface
	fmt.Printf("value %v, type %T \n", i, i) // prints type of argument and accepts everything
	
	//  assertion of interfaces
	v, t:= i.(int) // could get value and type from  v, t := interface.(T)
	fmt.Println(v, t ) // gives value and a boolean value 
}

func TestInterface()  {
	s:= "harish"
	test(s)
	ha:= struct{
		name string
	} {
		name : "harish",
	}
	test(ha)

	var in interface {} // empty interface
	in = 20
	test(in)

}

func InterfaceSwitchCase()  {

	var a interface{} = 33
	var b interface{} = 33.33
	var c interface{} = "harish"
	
	PrintsType(a)
	PrintsType(b)
	PrintsType(c)
}

func PrintsType(i interface{})  {
	switch i.(type) {
	case string:
		fmt.Printf("Its a %s \n", i.(string))
	case int:
		fmt.Printf("Its a %s \n", i.(int))
	default:
		fmt.Println("unknown type")
	} 
}
