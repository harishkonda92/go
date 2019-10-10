package structs

import (
	"fmt"
	"mymodule/DefiningStructs/structs"
)

// named structure
type empl struct {
	firstname, lastname string
	age, salary int
}

func MainStruct()  {
	// anonymous struct
	emp := struct {
		firstname , lastname string
		age int
	}{
		firstname:"harish" , 
		lastname:"konda",
		age: 26,
	}

	emp2 := empl{
		firstname: "harish ",
		lastname: "Konda",
		age: 26,
		salary: 33000,
	}

	// defaults stored as 0 when not defined a struct
	var emp3 empl
	// emp4 := empl{"konda", "harish"}
	fmt.Println("employee of the decade ", emp)
	fmt.Println("employee of the month ", emp2 )
	fmt.Println("emlpuee first name is", emp.firstname)
	emp3.firstname = "anuradha"
	// fmt.Println("employee of the day ", emp3)
	// fmt.Println("employee of the umm ", emp4)
	fmt.Println("emp3",emp3)
}

func PointersAndStruct()  {
	// emp1 := &empl {"harish", "konda", 26, 33300} // short hand declarion of a pointer 
	var emp1 *empl= &empl {"harish", "konda", 26, 33300} // general declaration of a pointer
	fmt.Printf("%v \n", *emp1) // or 
	fmt.Printf("%v \n", emp1.firstname) // same as above

	fmt.Println((*emp1).age)
}

// anonymous fields in structs
 type anony struct {
	 string
	 int
 }

 func MoreTypeOfStructs()  {
	 emp := anony {"harish", 33}
	 fmt.Println("emp", emp)
	 // to access the individual property, should call by its type i.e.
	 fmt.Println("string is ", emp.string)
 }

 // nested structs

 type SDetails struct {
	 address string
	 details empl
 }

 func NestedStructs (){
	 var person SDetails
	 person.address = "metpally, jagityal"
	 person.details.firstname = "Harish"
	 person.details.lastname = "Konda"
	 person.details.salary = 33300
	 person.details.age = 26

	 fmt.Println("employee name is: ", person.details.firstname)
 }

 // using exported struct

 func UsingExportedStruct()  {
	 var emp  structsType.Employee
	 emp.Name = "harish"
	 fmt.Println("name : ", emp.Name)
 }

//  comparing structs

func ComparingStruct()  {
	emp1 := empl{
		firstname: "harish",
		lastname: "Konda",
		age: 26,
		salary: 1,
	}
	emp2 := empl{
		firstname: "harish",
		lastname: "Konda",
		age: 26,
		salary: 1,
	}

	if emp2 == emp1 {
		fmt.Println("emp1 and emp2 are equal")
	} else {
		fmt.Println("emp1 and emp2 are inequal")
	}
}

// struct with maps

type image struct {
	name  map[int]int  // first int represents key and seccond is value inn Map
}

func MapsStruct()  {
	img := image{name: map[int]int{
		0: 1,
		2: 2,
	}}
	img2 := image{name: map[int]int{
		0: 1,
		2: 2,
	}}
	// if img == img2{  // maps in struct cannot be compared 
	// 	fmt.Println("equal")
	// } else {
	// 	fmt.Println("inequal")
	// }
	fmt.Println(img)
	fmt.Println(img2)
}