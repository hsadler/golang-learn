package main

import (
	"fmt"
	"math/cmplx"
)

// multiple variable declaration as "factored block"
var (
	b      bool       = false
	s      string     = "hi there"
	i      int        = -1234
	f      float32    = -1234.5678
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// print var types and values
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
