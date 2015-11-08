package main

import (
	"fmt";
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	i := 0
	old := 0.0
	for math.Floor(z) != math.Floor(old) {
		old = z
		z = z - (z*z-x)/(2*z)
		fmt.Printf("%v\n", z)
		i++
	}
	return z, i
}

func main() {
	fmt.Println(Sqrt(1000))
}