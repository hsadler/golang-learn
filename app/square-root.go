
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	// get closer to actual value with every iteration
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	// compare defined function with standard library function
	n := 2.0
	fmt.Println(Sqrt(n), math.Sqrt(n))
}
