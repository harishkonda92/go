package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("ok", ok)
		// fmt.Println("file at path", err.Path, "failed to open")
		fmt.Println(err.Path)
		return
	}
	fmt.Println(f.Name(), "opened successfully")

}
