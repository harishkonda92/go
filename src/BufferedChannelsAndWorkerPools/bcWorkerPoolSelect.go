package main

import (
	"fmt"
	"time"
)

func main() {
	output1 := make(chan string)
	ouptut2 := make(chan string)
	go server1(output1)
	go server2(ouptut2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-ouptut2:
		fmt.Println(s2)
	}
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}
