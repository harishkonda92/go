package rectangle

import (
	"math"
)

func Area(length, width float64) float64 {
	area := length * width
	return area
}

func Diagonal(length, width float64) float64 {
	diagonal := math.Sqrt((length * length) + (width * width))
	return diagonal
}
