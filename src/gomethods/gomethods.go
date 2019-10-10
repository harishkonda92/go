package gomethods

import (
	"fmt"
)

type Employee struct {
	firstname , lastname string
	age int
}

func (e Employee ) employeeReceiverMethod(name string)  {
	fmt.Println("employee details are : ", e)	
	fmt.Println("employee name is : ", name)	
}

func EmployeeCaller()  {
	emp := Employee {
		firstname: "harish",
		lastname: "konda",
		age: 26,
	}
	emp.employeeReceiverMethod(emp.firstname)
}	

// pointer receiver vs value receiver in methods

func (e Employee) changeFirstname(name string)  {
	// fmt.Println("inside change first name", e)
	e.firstname = name
	// fmt.Println("inside change first name after change: ", e)
}

func (e *Employee) changeLastname(name string)  {
	fmt.Printf("%T \n", e)
	e.lastname = name
}

func ExMethods2()  {
	emp := Employee{
		firstname: "harish",
		lastname: "konda",
		age: 26,
	}
	emp.changeFirstname("Ajay")
	fmt.Println("after changing fisrtname :",emp)
	emp.changeLastname("Nampally")
	fmt.Println("after changing lastname :",emp)

}

// anonymous structs

type address struct {
	city , state string
}

type person struct {
	name string
	age int
	address
}

func (a address) printFullAddress()  {
	fmt.Printf("full address is %s, %s \n" , a.city, a.state)
	fmt.Printf("full address is %v,\n " ,a)
}

func GetValueOfAnonymousStruct()  {
	p:= person {
		name: "harish",
		age: 26,
		address: address{
			city: "hyderabad",
			state: "Telangana",
		},
	}
	// p.address.printFullAddress() // for calling reciever adderess we dont need to use by inner struct 
	p.printFullAddress()
}

