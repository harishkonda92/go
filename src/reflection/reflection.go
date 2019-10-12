package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId    int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)  // gives type of the argument passesd i.e, in order
	v := reflect.ValueOf(q) // so is the value of argument
	k := t.Kind()           // aswell the kind of the argument
	fmt.Println("Type ,", t)
	fmt.Println("Value, ", v)
	fmt.Println("Kind, ", k)

	/////////////////////////////////

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields,", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("field: %d type :%T value:%v \n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	o := order{
		orderId:    111,
		customerId: 3434,
	}
	createQuery(o)
}
