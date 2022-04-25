package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("idx", i, "z value", z)
	}
	return z
}

func main() {
	fmt.Println("Final Sqrt value", Sqrt(128))
	fmt.Println("gmath Sqrt value", math.Sqrt(128))
}
