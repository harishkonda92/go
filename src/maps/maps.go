package maps

import (
	"fmt"
)

func MainFunc() {
	aMap := make(map[string]int)
	if aMap == nil {
		fmt.Println("the map is nill")
	}
	aMap["a"] = 1;
	aMap["b"] = 2;
	aMap["c"] = 3;
	fmt.Printf("type of aMap %T \n", aMap)
	fmt.Println( aMap)
	bmap := map[string]string {"name":"harish", "lastname": "konda" }
	fmt.Println(bmap)
	bmap["age"] = "26"
	fmt.Println(bmap)
}