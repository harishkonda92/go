package channels

import (
	"fmt"
)

func TestChannel()  {
	var ch chan int // this is how we declare a channel, an undefined channel has nil value
	fmt.Printf("type is:%T, value is %d \n", ch, ch)
	if ch == nil {
		ch = make(chan int) // channel is defined using make function just as slices
		fmt.Printf("%T, %d", ch, ch)
	}
}