package channels

import (
	"fmt"
	"time"
)

func hello(done chan bool)  {
	fmt.Println("im in go routine hello")
	time.Sleep(4 * time.Second)
	done <- true
}

func CreateChannel()  {
	done := make(chan bool)
	go hello(done)
	<- done
	// time.Sleep(1 * time.Second)
	fmt.Println("im in createchannkl")
}