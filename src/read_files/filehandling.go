package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	data, err := ioutil.ReadFile("text.txt")
	if err!= nil {
		fmt.Println()
	}
}
