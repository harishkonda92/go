package channels

import (
	"fmt"
)

func calcSquare(number int, squareOp chan int)  {
	sum := 0 
	for number!=0 {
		digit := number % 10
		fmt.Println(digit)
        sum += digit * digit
		number /= 10
		fmt.Println(number)		
	}
	squareOp <- sum
}

func calcCube(number int, cubeOp chan int)  {
	sum := 0
	for number != 0{
		digit := number % 10
        sum += digit * digit * digit
        number /= 10
	}
	cubeOp <- sum	
}

func CalcCubesnSquares()  {
	number := 5
	squareOp := make(chan int)
	cubeOp := make(chan int)
	go calcCube(number, cubeOp)
	go calcSquare(number, squareOp)

	squares, cubes := <- squareOp, <- cubeOp

	fmt.Printf("cubes are %d squares are %d", cubes, squares)


}