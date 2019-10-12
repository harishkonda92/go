package main

import (
	"fmt"
)

func finished() {
	fmt.Println("finished finding the laargest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("started finding the largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("largest number in ", nums, " is ", max)
}

func main() {
	nums := []int{778, 109, 34, 343, 34}
	largest(nums)
}
