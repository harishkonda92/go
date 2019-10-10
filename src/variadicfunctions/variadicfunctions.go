package variadicfunctions
import (
	"fmt"
)
func find(num int, nums ... int){
	fmt.Printf("type of nums is %T \n ", nums)
	found := false
	for i, numm := range nums {
		if numm == num {
			fmt.Println("found index", i ,"in ", nums)
			found = true
		}
	}
	if !found {
		fmt.Println("num not found")
	}
}

func takesString(a ...string){
	fmt.Printf("type of a is, %T\n", a)
	a = append(a, "konda")
	fmt.Println(a)
}

func MainFunc(){
	// find(1, 2, 3,1, 4, 5) // or
	num:= []int{2,3,4,5,1} // this will work with spread operator or elipsis
	find(1, num...)

	name := []string{"harish"}
	takesString(name...)
	fmt.Println(name)
}