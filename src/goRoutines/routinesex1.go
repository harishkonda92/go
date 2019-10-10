package goRoutines

import (
	"fmt"
	"time"
)

func Routine()  {
	fmt.Println("im routine")
}

func HelloRoutine()  {
	go Routine()
	time.Sleep(1* time.Second)
	fmt.Println("im hello routine")
}