package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// addr, err := net.LookupHost("golangbot.com")
	// if err, ok := err.(*net.DNSError); ok {
	// 	// fmt.Println(addr)
	// 	// fmt.Println(err)
	// 	if err.Timeout() {
	// 		fmt.Println("operation timed out")
	// 	} else if err.Temporary() {
	// 		fmt.Println("temporary error ")
	// 	} else {
	// 		fmt.Println("generic error: ", err)
	// 	}
	// 	return
	// }
	// fmt.Println(addr)

	files, _ := filepath.Glob("/")
	// if error != nil && error == filepath.ErrBadPattern {
	// 	fmt.Println(error)
	// 	return
	// }
	fmt.Println("matched files, ", files)
}
