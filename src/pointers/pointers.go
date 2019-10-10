package pointers
import (
	"fmt"
)

func PointersFunc()  {
	b:= 250 
	var a *int = &b // store address in pointers i.e. type *int
	fmt.Printf("a is a type %T\n ",a )
	fmt.Println("address of b is ", a)	

	var d *int
	if d == nil{
		fmt.Println("d is", d);
		d = &b
		fmt.Println("the address of b is ", d)
		fmt.Println("the value of a is %v ", a)
	}

}

func CreateNewPointer()  {
	z:= new(int) // new() is function to create a new pointer with size 0
	fmt.Printf("size of z %d, value of z is %v, type of z is %T \n", *z, z, z)
	*z = 85
	fmt.Printf("size of z is %d \n", *z)
}

func DereferencingPonter()  {
	a:= 100
	var b *int = &a
	fmt.Printf("value of a is %d size of b is %v \n", a, b )

	// dereferncing the pointer to its value of its address
	fmt.Printf(" the value of a is %d \n", *b) 

}

func ChangeValueUsingPointer()  {
	a:= 100
	var b *int= &a // b pointer stores address of a
	*b++
	fmt.Printf("a= %d", *b)
}

func PassingPointer()  {
	a := 100
	b := &a
	fmt.Println("a is =", *b)
	changeValInPointer(b)
	fmt.Println("a is =", *b)
}

func changeValInPointer(val *int)  {
	*val = 9
}

func CaputersAPointer()  {
	val := retrunsAPointer()
	fmt.Println("address is", val)
	fmt.Println("value is", *val)
	*val = "Harish Konda"
	fmt.Println("cahnged value is", *val)
}

func retrunsAPointer() *string  {
	i := "harish"
	return &i
}

func ArrayPointers(){
	a := [2]string{"harishh", "konda"}
	fmt.Println("a is ,", a)
	// modify(a[:])
	modify(&a) // this way of passing pointer to an arry as argument is not idiomatic 
			   // should use slices instead
	fmt.Println("after modification a is ,", a)
}

func modify(a *[2]string)  {
	a[0] = "Harish"
}