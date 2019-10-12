package main

import (
	"fmt"
	"sync"
)

var y = 0

func increment1(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	y = y + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment1(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of y is", y)
}
