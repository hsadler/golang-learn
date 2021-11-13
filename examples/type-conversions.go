package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int = 3, 4

	// cannot coerce by assignment
	// var f float64 = math.Sqrt(x*x + y*y)
	// var z uint = f

	// must use explicit casting
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)

	// more explicit casting examples
	i := 64
	j := string(i)
	fmt.Println(j)

}
